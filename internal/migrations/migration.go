/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package migrations implements the Mediator design pattern used to manage migrations over sveltin versions.
package migrations

import (
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/internal/fsm"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/yinlog"
)

// IMigration is the interface defining the methods to be implemented by single migration.
type IMigration interface {
	Execute() error
	getServices() *MigrationServices
	getData() *MigrationData
	up() error
	down() error
	allowUp() error
}

// MigrationServices contains references to services used by the migrations.
type MigrationServices struct {
	fs        afero.Fs
	fsManager *fsm.SveltinFSManager
	pathMaker *pathmaker.SveltinPathMaker
	logger    *yinlog.Logger
}

// NewMigrationServices creates an instance of MigrationService struct.
func NewMigrationServices(fs afero.Fs, fsm *fsm.SveltinFSManager, pathmaker *pathmaker.SveltinPathMaker, logger *yinlog.Logger) *MigrationServices {
	return &MigrationServices{
		fs:        fs,
		fsManager: fsm,
		pathMaker: pathmaker,
		logger:    logger,
	}
}

// MigrationData is the struct with data used by migrations.
type MigrationData struct {
	PathToFile        string
	CliVersion        string
	ProjectCliVersion string
}

// MigrationRule is the struct with settings to be matched for running the migration.
type migrationRule struct {
	value           string
	pattern         string
	replaceFullLine bool
	replacerFunc    func(string) string
}

type matcherFunc = func([]byte, string, string) ([]byte, bool)

//=============================================================================

func isMigrationRequired(m IMigration, patterns []string, matcher matcherFunc) ([]byte, bool) {
	content, err := afero.ReadFile(m.getServices().fs, m.getData().PathToFile)
	if err != nil {
		return nil, false
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		for _, pattern := range patterns {
			if r, ok := matcher(content, pattern, line); ok {
				return r, true
			}
			continue
		}
	}
	return content, false
}

func applyMigrationRules(rules []*migrationRule) (string, bool) {
	for _, r := range rules {
		expression := regexp.MustCompile(r.pattern)

		if expression.MatchString(r.value) {
			if r.replaceFullLine {
				return r.replacerFunc(r.value), true
			}
			return expression.ReplaceAllStringFunc(r.value, r.replacerFunc), true
		}
	}
	return "", false
}

func findStringMatcher(content []byte, pattern, line string) ([]byte, bool) {
	rule := regexp.MustCompile(pattern)
	matches := rule.FindString(line)
	if len(matches) > 0 {
		return content, true
	}
	return nil, false
}

//=============================================================================

func retrieveFileContent(m IMigration) ([]byte, error) {
	content, err := afero.ReadFile(m.getServices().fs, m.getData().PathToFile)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func overwriteFile(m IMigration, content []byte) error {
	err := m.getServices().fs.Remove(m.getData().PathToFile)
	if err != nil {
		return err
	}

	if err = afero.WriteFile(m.getServices().fs, m.getData().PathToFile, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

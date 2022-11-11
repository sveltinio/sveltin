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
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/yinlog"
)

// Migration is the interface defining the methods to be implemented by single migration.
type Migration interface {
	Execute() error
	getFs() afero.Fs
	getPathMaker() *pathmaker.SveltinPathMaker
	getLogger() *yinlog.Logger
	getData() *MigrationData
	allowUp() error
}

// MigrationData is the struct with data used by migrations.
type MigrationData struct {
	PathToFile        string
	CliVersion        string
	ProjectCliVersion string
}

// MigrationRule is the struct with settings to be matched for running the migration.
type MigrationRule struct {
	Value             string
	Pattern           string
	IsReplaceFullLine bool
	GetMatchReplacer  func(string) string
}

type matcherFunc = func([]byte, string, string) ([]byte, bool)

func isMigrationRequired(m Migration, pattern string, matcher matcherFunc) ([]byte, bool) {
	content, err := afero.ReadFile(m.getFs(), m.getData().PathToFile)
	if err != nil {
		return nil, false
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if r, ok := matcher(content, pattern, line); ok {
			return r, true
		}
	}
	return nil, false
}

func applyMigrationRules(rules []*MigrationRule) (string, bool) {
	for _, r := range rules {
		rule := regexp.MustCompile(r.Pattern)

		if rule.MatchString(r.Value) {
			if r.IsReplaceFullLine {
				return r.GetMatchReplacer(r.Value), true
			}
			return rule.ReplaceAllStringFunc(r.Value, r.GetMatchReplacer), true
		}
	}
	return "", false
}

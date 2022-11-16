/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/internal/fsm"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/yinlog"
)

// Patterns used by MigrationRule
const (
	SitemapPattern               = `^sitemap`
	SvelteKitBuildPattern        = `^SVELTEKIT_BUILD_FOLDER`
	SvelteKitBuildCommentPattern = `^*# The folder where adaper-static`
)

//=============================================================================

// UpdateDotEnvMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateDotEnvMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

func (m *UpdateDotEnvMigration) getFs() afero.Fs { return m.Fs }
func (m *UpdateDotEnvMigration) getPathMaker() *pathmaker.SveltinPathMaker {
	return m.PathMaker
}
func (m *UpdateDotEnvMigration) getLogger() *yinlog.Logger { return m.Logger }
func (m *UpdateDotEnvMigration) getData() *MigrationData   { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m UpdateDotEnvMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDotEnvMigration) up() error {

	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	if exists {
		if fileContent, ok := isMigrationRequired(m, SitemapPattern, findStringMatcher); ok {
			m.Logger.Info(fmt.Sprintf("Migrating %s file", filepath.Base(m.Data.PathToFile)))
			if err := updateDotEnvFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateDotEnvMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDotEnvMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updateDotEnvFile(m *UpdateDotEnvMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*MigrationRule{
			newDotEnvSitemapRule(line),
			newDotEnvSvelteKitBuildCommentRule(line),
			newDotEnvSveltekitRule(line),
		}
		if res, ok := applyMigrationRules(rules); ok {
			lines[i] = res
		} else {
			lines[i] = line
		}
	}
	output := strings.Join(lines, "\n")
	err := m.Fs.Remove(m.Data.PathToFile)
	if err != nil {
		return err
	}

	cleanedOutput := removeMultiEmptyLines(output)
	if err = afero.WriteFile(m.Fs, m.Data.PathToFile, cleanedOutput, 0644); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newDotEnvSitemapRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           SitemapPattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return ""
		},
	}
}

func newDotEnvSvelteKitBuildCommentRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           SvelteKitBuildCommentPattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return ""
		},
	}
}

func newDotEnvSveltekitRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           SvelteKitBuildPattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return ""
		},
	}
}

// =============================================================================

func removeMultiEmptyLines(content string) []byte {
	rule := regexp.MustCompile(`[\n]+`)
	output := rule.ReplaceAllString(strings.TrimSpace(content), "\n")
	return []byte(output)
}

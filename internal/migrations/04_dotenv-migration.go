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
)

// Patterns used by MigrationRule
const (
	sitemapPattern               = `^sitemap`
	svelteKitBuildPattern        = `^SVELTEKIT_BUILD_FOLDER`
	svelteKitBuildCommentPattern = `^*# The folder where adaper-static`
)

//=============================================================================

// UpdateDotEnvMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateDotEnvMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdateDotEnvMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateDotEnvMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateDotEnvMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateDotEnvMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
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

	exists, err := common.FileExists(m.getServices().fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	migrationTriggers := []string{sitemapPattern, svelteKitBuildPattern, svelteKitBuildCommentPattern}
	if exists {
		if fileContent, ok := isMigrationRequired(m, migrationTriggers, findStringMatcher); ok {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
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
		rules := []*migrationRule{
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
	err := m.getServices().fs.Remove(m.Data.PathToFile)
	if err != nil {
		return err
	}

	cleanedOutput := removeMultiEmptyLines(output)
	if err = afero.WriteFile(m.getServices().fs, m.Data.PathToFile, cleanedOutput, 0644); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newDotEnvSitemapRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         sitemapPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newDotEnvSvelteKitBuildCommentRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         svelteKitBuildCommentPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newDotEnvSveltekitRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         svelteKitBuildPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
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

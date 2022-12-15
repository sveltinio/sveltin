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

	exists, err := common.FileExists(m.getServices().fs, m.Data.FileToMigrate)
	if err != nil {
		return err
	}

	if exists {
		fileContent, err := retrieveFileContent(m.getServices().fs, m.getData().FileToMigrate)
		if err != nil {
			return err
		}

		migrationTriggers := []string{patterns[sitemap], patterns[svelteKitBuildFolder], patterns[svelteKitBuildComment]}
		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.FileToMigrate)))
			if _, err := m.migrate(fileContent); err != nil {
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

func (m *UpdateDotEnvMigration) migrate(content []byte) ([]byte, error) {
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
	err := m.getServices().fs.Remove(m.Data.FileToMigrate)
	if err != nil {
		return nil, err
	}

	cleanedOutput := removeMultiEmptyLines(output)
	if err = afero.WriteFile(m.getServices().fs, m.Data.FileToMigrate, cleanedOutput, 0644); err != nil {
		return nil, err
	}
	return nil, nil
}

//=============================================================================

func newDotEnvSitemapRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[sitemap],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newDotEnvSveltekitRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[svelteKitBuildFolder],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newDotEnvSvelteKitBuildCommentRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[svelteKitBuildComment],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

// =============================================================================

func removeMultiEmptyLines(content string) []byte {
	rule := regexp.MustCompile(`^\n{1,}$`)
	output := rule.ReplaceAllString(strings.TrimSpace(content), "\n")
	return []byte(output)
}

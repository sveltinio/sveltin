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
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
)

// UpdateStringsTSMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateStringsTSMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *UpdateStringsTSMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateStringsTSMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateStringsTSMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateStringsTSMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateStringsTSMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateStringsTSMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.getServices().fs, m.Data.TargetPath)
	if err != nil {
		return err
	}

	if exists {
		fileContent, err := retrieveFileContent(m.getServices().fs, m.getData().TargetPath)
		if err != nil {
			return err
		}

		migrationTriggers := []string{
			patterns[importIWebSiteSeoType],
			patterns[icontententryTypeUsage],
			patterns[iwebsiteSeoTypeUsage],
		}

		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if _, err := m.migrate(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateStringsTSMigration) migrate(content []byte, filepath string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")

	// It must be executed twice to replace multiple triggers on the same line
	for loopCounter := 0; loopCounter <= 1; loopCounter++ {
		fmt.Println(loopCounter)
		for i, line := range lines {
			rules := []*migrationRule{
				newStringsTSImportRule(line),
				newStringsTSContentEntryUsageRule(line),
				newStringsTSIWebSiteUsageRule(line),
			}

			if res, ok := applyMigrationRules(rules); ok {
				lines[i] = res
			} else {
				lines[i] = line
			}

		}
	}
	output := strings.Join(lines, "\n")
	err := m.getServices().fs.Remove(m.Data.TargetPath)
	if err != nil {
		return nil, err
	}

	if err = afero.WriteFile(m.getServices().fs, m.Data.TargetPath, []byte(output), 0644); err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *UpdateStringsTSMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateStringsTSMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newStringsTSImportRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[importIWebSiteSeoType],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newStringsTSContentEntryUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[icontententryTypeUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "ResourceContent"
		},
	}
}

func newStringsTSIWebSiteUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[iwebsiteSeoTypeUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "Sveltin.WebSite"
		},
	}
}

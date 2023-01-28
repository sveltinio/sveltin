/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"fmt"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
)

// UpdateMenuTSMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateMenuTSMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *UpdateMenuTSMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateMenuTSMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateMenuTSMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateMenuTSMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateMenuTSMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMenuTSMigration) up() error {
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

		migrationTriggers := []string{patterns[importIMenuItemSeoType], patterns[imenuitemSeoTypeUsage]}
		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			localFilePath :=
				strings.Replace(m.Data.TargetPath, m.getServices().pathMaker.GetRootFolder(), "", 1)
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", localFilePath))
			if _, err := m.migrate(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateMenuTSMigration) migrate(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newMenuTSImportRule(line),
			newMenuTSUsageRule(line),
		}
		if res, ok := applyMigrationRules(rules); ok {
			lines[i] = res
		} else {
			lines[i] = line
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

func (m *UpdateMenuTSMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMenuTSMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newMenuTSImportRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[importIMenuItemSeoType],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "import type { Sveltin } from '$sveltin';"
		},
	}
}

func newMenuTSUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[imenuitemSeoTypeUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "Sveltin.MenuItem"
		},
	}
}

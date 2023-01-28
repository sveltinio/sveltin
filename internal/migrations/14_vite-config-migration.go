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

// UpdateViteConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateViteConfigMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdateViteConfigMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateViteConfigMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateViteConfigMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateViteConfigMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateViteConfigMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateViteConfigMigration) up() error {
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

		migrationTriggers := []string{patterns[viteAlias]}
		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if _, err := m.migrate(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateViteConfigMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateViteConfigMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateViteConfigMigration) migrate(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newViteConfigRule(line),
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

	cleanedOutput := removeMultiEmptyLines(output)
	if err = afero.WriteFile(m.getServices().fs, m.Data.TargetPath, cleanedOutput, 0644); err != nil {
		return nil, err
	}
	return nil, nil
}

//=============================================================================

func newViteConfigRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[viteAlias],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "\t\talias: {\n\t\t\t$sveltin: path.resolve(path.join(__dirname, './src/sveltin')),"
		},
	}
}

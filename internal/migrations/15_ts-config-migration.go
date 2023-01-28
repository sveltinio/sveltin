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

// UpdateTSConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateTSConfigMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdateTSConfigMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateTSConfigMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateTSConfigMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateTSConfigMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateTSConfigMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateTSConfigMigration) up() error {
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

		migrationTriggers := []string{patterns[tsPath]}
		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if _, err := m.migrate(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateTSConfigMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateTSConfigMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateTSConfigMigration) migrate(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newTSConfigRule(line),
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

func newTSConfigRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[tsPath],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "\t\t\"paths\": {\n\t\t\t\"$sveltin\": [\"./src/sveltin\"],"
		},
	}
}

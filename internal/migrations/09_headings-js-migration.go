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

// UpdateHeadingJSMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateHeadingJSMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *UpdateHeadingJSMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateHeadingJSMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateHeadingJSMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateHeadingJSMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateHeadingJSMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateHeadingJSMigration) up() error {
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

		migrationTriggers := []string{patterns[headingsTitleProp]}
		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.FileToMigrate)))
			if _, err := m.migrate(fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateHeadingJSMigration) migrate(content []byte) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newHeadingsJSRule(line),
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

	if err = afero.WriteFile(m.getServices().fs, m.Data.FileToMigrate, []byte(output), 0644); err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *UpdateHeadingJSMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateHeadingJSMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newHeadingsJSRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[headingsTitleProp],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "value:"
		},
	}
}

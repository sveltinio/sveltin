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

const trailingSlashPattern = `trailingSlash`

//=============================================================================

// UpdateSvelteConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateSvelteConfigMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *UpdateSvelteConfigMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateSvelteConfigMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateSvelteConfigMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateSvelteConfigMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateSvelteConfigMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateSvelteConfigMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.getServices().fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	migrationTriggers := []string{trailingSlashPattern}
	if exists {
		if fileContent, ok := isMigrationRequired(m, migrationTriggers, findStringMatcher); ok {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
			if err := updateSvelteConfigFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateSvelteConfigMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateSvelteConfigMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updateSvelteConfigFile(m *UpdateSvelteConfigMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{newSvelteConfigRule(line)}
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

	if err = afero.WriteFile(m.getServices().fs, m.Data.PathToFile, []byte(output), 0644); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newSvelteConfigRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         trailingSlashPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

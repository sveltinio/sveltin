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

// Patterns used by MigrationRule
const (
	remarkAutolinkHeadingsPkgPattern = `"remark-external-links"`
)

//=============================================================================

// UpdatePkgJSONMigration is the struct representing the migration update the defaults.js.ts file.
type UpdatePkgJSONMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdatePkgJSONMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdatePkgJSONMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// MakeMigration implements IMigration interface.
func (m *UpdatePkgJSONMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdatePkgJSONMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdatePkgJSONMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdatePkgJSONMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.getServices().fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	migrationTriggers := []string{remarkAutolinkHeadingsPkgPattern}
	if exists {
		if fileContent, ok := isMigrationRequired(m, migrationTriggers, findStringMatcher); ok {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
			if err := updatePkgJSONFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdatePkgJSONMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdatePkgJSONMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updatePkgJSONFile(m *UpdatePkgJSONMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newRemarkExternalLinksRule(line),
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

	if err = afero.WriteFile(m.getServices().fs, m.Data.PathToFile, []byte(output), 0644); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newRemarkExternalLinksRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         remarkAutolinkHeadingsPkgPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "\"rehype-external-links\":\"^2.0.1\","
		},
	}
}

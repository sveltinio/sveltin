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
	configConstPattern  = `^const config`
	exportConfigPattern = `^export default config`
	namePropPattern     = `name:`
)

//=============================================================================

// UpdateThemeConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateThemeConfigMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdateThemeConfigMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateThemeConfigMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateThemeConfigMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateThemeConfigMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m UpdateThemeConfigMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateThemeConfigMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.getServices().fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	migrationTriggers := []string{configConstPattern, exportConfigPattern}
	if exists {
		if fileContent, ok := isMigrationRequired(m, migrationTriggers, findStringMatcher); ok {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
			if err := updateThemeFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateThemeConfigMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateThemeConfigMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updateThemeFile(m *UpdateThemeConfigMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var prevLine string
		if i > 0 {
			prevLine = lines[i-1]
		}

		rules := []*migrationRule{
			newConstNameRule(line),
			newExportLineRule(line),
			newThemeNameRule(line, prevLine),
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

func newConstNameRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         configConstPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return `import { theme } from '../../sveltin.json';

const themeConfig = {`
		},
	}
}

func newExportLineRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         exportConfigPattern,
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "export { themeConfig }"
		},
	}
}

func newThemeNameRule(line, prevLine string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         namePropPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			if !strings.Contains(prevLine, "author:") && strings.Contains(line, "name:") {
				return "\tname: theme.name,"
			}
			return line
		},
	}
}

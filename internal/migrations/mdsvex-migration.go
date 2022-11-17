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
	remarkExternalLinksImportStrPattern = `^import remarkExternalLinks`
	remarkExternalLinksUsagePattern     = `\[remarkExternalLinks`
	rehypePluginPattern                 = `rehypePlugins:[\t\s]+\[`
)

//=============================================================================

// UpdateMDsveXMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateMDsveXMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdateMDsveXMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateMDsveXMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateMDsveXMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateMDsveXMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateMDsveXMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXMigration) up() error {

	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.getServices().fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	migrationTriggers := []string{remarkExternalLinksImportStrPattern, remarkExternalLinksUsagePattern}
	if exists {
		if fileContent, ok := isMigrationRequired(m, migrationTriggers, findStringMatcher); ok {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
			if err := updateMDsveXFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateMDsveXMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updateMDsveXFile(m *UpdateMDsveXMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newReplaceRemarkExternalLinksImportStrRule(line),
			newReplaceRemarkExternalLinksUsageRule(line),
			newReplaceRehypePluginUsageRule(line),
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

func newReplaceRemarkExternalLinksImportStrRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         remarkExternalLinksImportStrPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "import rehypeExternalLinks from 'rehype-external-links';"
		},
	}
}

func newReplaceRemarkExternalLinksUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         remarkExternalLinksUsagePattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newReplaceRehypePluginUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		pattern:         rehypePluginPattern,
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return `rehypePlugins: [
		[rehypeExternalLinks, { target: '_blank', rel: ['noopener', 'noreferrer'] }],`
		},
	}
}

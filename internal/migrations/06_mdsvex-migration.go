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
			patterns[remarkExtLinksImport],
			patterns[remarkSlugImport],
			patterns[headingsImport],
			patterns[remarkSlugUsage],
			patterns[remarkExtLinksUsage],
		}

		if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if _, err := m.migrate(fileContent); err != nil {
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

func (m *UpdateMDsveXMigration) migrate(content []byte) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newReplaceHeadingsImportStrRule(line),
			newReplaceRemarkExternalLinksImportStrRule(line),
			newReplaceRemarkExternalLinksUsageRule(line),
			newReplaceRemarkSlugImportStrRule(line),
			newReplaceRemarkSlugUsageRule(line),
			newReplaceRehypePluginUsageRule(line),
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

//=============================================================================

func newReplaceHeadingsImportStrRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[headingsImport],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "import headings from '@sveltinio/remark-headings';"
		},
	}
}

func newReplaceRemarkExternalLinksImportStrRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[remarkExtLinksImport],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "import rehypeExternalLinks from 'rehype-external-links';"
		},
	}
}

func newReplaceRemarkExternalLinksUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[remarkExtLinksUsage],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newReplaceRemarkSlugImportStrRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[remarkSlugImport],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newReplaceRemarkSlugUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[remarkSlugUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newReplaceRehypePluginUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[rehypePlugins],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return `rehypePlugins: [
		[rehypeExternalLinks, { target: '_blank', rel: ['noopener', 'noreferrer'] }],`
		},
	}
}

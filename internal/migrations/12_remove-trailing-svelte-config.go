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

// RemoveTrailingFromSvelteConfig is the struct representing the migration update the defaults.js.ts file.
type RemoveTrailingFromSvelteConfig struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *RemoveTrailingFromSvelteConfig) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &RemoveTrailingFromSvelteConfig{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *RemoveTrailingFromSvelteConfig) getServices() *MigrationServices { return m.Services }
func (m *RemoveTrailingFromSvelteConfig) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m RemoveTrailingFromSvelteConfig) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *RemoveTrailingFromSvelteConfig) up() error {
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
			patterns[trailingSlash],
			patterns[prerenderEnabled],
		}
		if patternsMatched(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if _, err := m.runMigration(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *RemoveTrailingFromSvelteConfig) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *RemoveTrailingFromSvelteConfig) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *RemoveTrailingFromSvelteConfig) runMigration(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newSvelteConfigTrailingSlashRule(line),
			newSvelteConfigPrerenderEnabledRule(line),
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

func newSvelteConfigTrailingSlashRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[trailingSlash],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newSvelteConfigPrerenderEnabledRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[prerenderEnabled],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

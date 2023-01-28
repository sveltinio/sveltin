/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// UpdateResourceLibsMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateResourceLibsMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *UpdateResourceLibsMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateResourceLibsMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateResourceLibsMigration) getServices() *MigrationServices { return m.Services }
func (m *UpdateResourceLibsMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateResourceLibsMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateResourceLibsMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := afero.DirExists(m.getServices().fs, m.Data.TargetPath)
	if err != nil {
		return err
	}

	if exists {
		files := []string{}

		walkFunc := func(file string, info os.FileInfo, err error) error {
			if filepath.Ext(file) == ".ts" {
				files = append(files, file)
			}
			return nil
		}

		err := afero.Walk(m.getServices().fs, m.Data.TargetPath, walkFunc)
		if err != nil {
			m.getServices().logger.Fatalf("Something went wrong visiting the folder %s. Are you sure it exists?", m.Data.TargetPath)
		}

		migrationTriggers := []string{
			patterns[sveltinNamespace],
			patterns[importIWebSiteSeoType],
			patterns[icontententryTypeUsage],
			patterns[iwebsiteSeoTypeUsage],
		}

		for _, file := range files {
			fileContent, err := retrieveFileContent(m.getServices().fs, file)
			if err != nil {
				return err
			}

			if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
				localFilePath :=
					strings.Replace(file, m.getServices().pathMaker.GetRootFolder(), "", 1)
				m.getServices().logger.Info(fmt.Sprintf("Migrating %s", localFilePath))
				if _, err := m.migrate(fileContent, file); err != nil {
					return err
				}
			}

		}

	}

	return nil
}

func (m *UpdateResourceLibsMigration) migrate(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")

	// It must be executed twice to replace multiple triggers on the same line
	for loopCounter := 0; loopCounter <= 1; loopCounter++ {
		for i, line := range lines {
			rules := []*migrationRule{
				newSveltinNamespaceRule(line),
				newIWebSiteImportRule(line),
				newContentEntryUsageRule(line),
				newIWebSiteUsageRule(line),
			}

			if res, ok := applyMigrationRules(rules); ok {
				lines[i] = res
			} else {
				lines[i] = line
			}

		}
	}
	output := strings.Join(lines, "\n")
	err := m.getServices().fs.Remove(file)
	if err != nil {
		return nil, err
	}

	if err = afero.WriteFile(m.getServices().fs, file, []byte(output), 0644); err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *UpdateResourceLibsMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateResourceLibsMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newSveltinNamespaceRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[sveltinNamespace],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "import type { Sveltin } from '$sveltin';"
		},
	}
}

func newIWebSiteImportRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[importIWebSiteSeoType],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newContentEntryUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[icontententryTypeUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "ResourceContent"
		},
	}
}

func newIWebSiteUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[iwebsiteSeoTypeUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "Sveltin.WebSite"
		},
	}
}

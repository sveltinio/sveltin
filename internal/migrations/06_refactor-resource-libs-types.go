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

// RefactorResourcesLibsTypes is the struct representing the migration update the defaults.js.ts file.
type RefactorResourcesLibsTypes struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *RefactorResourcesLibsTypes) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &RefactorResourcesLibsTypes{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *RefactorResourcesLibsTypes) getServices() *MigrationServices { return m.Services }
func (m *RefactorResourcesLibsTypes) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m RefactorResourcesLibsTypes) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *RefactorResourcesLibsTypes) up() error {
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
			patterns[capitalizeAll],
			patterns[capitalizeFirstLetter],
			patterns[camelToKebabCase],
			patterns[toTitle],
			patterns[toSlug],
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
				if _, err := m.runMigration(fileContent, file); err != nil {
					return err
				}
			}

		}

	}

	return nil
}

func (m *RefactorResourcesLibsTypes) runMigration(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")

	// It must be executed twice to replace multiple triggers on the same line
	for loopCounter := 0; loopCounter <= 1; loopCounter++ {
		for i, line := range lines {
			rules := []*migrationRule{
				newSveltinNamespaceRule(line),
				newIWebSiteImportRule(line),
				newContentEntryUsageRule(line),
				newIWebSiteUsageRule(line),
				newCapitalizeAllRule(line),
				newCapitalizeFirstLetterRule(line),
				newCamelToKebabCaseRule(line),
				newToTitleRule(line),
				newToSlug(line),
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

func (m *RefactorResourcesLibsTypes) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *RefactorResourcesLibsTypes) allowUp() error {
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

func newCapitalizeAllRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[capitalizeAll],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "capitalizeAll"
		},
	}
}

func newCapitalizeFirstLetterRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[capitalizeFirstLetter],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "capitalizeFirstLetter"
		},
	}
}

func newCamelToKebabCaseRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[camelToKebabCase],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "camelToKebabCase"
		},
	}
}

func newToTitleRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[toTitle],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "toTitle"
		},
	}
}

func newToSlug(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[toSlug],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "toSlug"
		},
	}
}

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
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
)

// RefactorSvelteFilesTypes is the struct representing the migration update the defaults.js.ts file.
type RefactorSvelteFilesTypes struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *RefactorSvelteFilesTypes) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &RefactorSvelteFilesTypes{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *RefactorSvelteFilesTypes) getServices() *MigrationServices { return m.Services }
func (m *RefactorSvelteFilesTypes) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m RefactorSvelteFilesTypes) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *RefactorSvelteFilesTypes) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := afero.DirExists(m.getServices().fs, m.Data.TargetPath)
	if err != nil {
		return err
	}

	if exists {
		files := []string{}
		targetFiles := []string{"+layout.svelte", "+page.svelte", "+page.svx"}

		walkFunc := func(file string, info os.FileInfo, err error) error {
			if common.Contains(targetFiles, info.Name()) {
				files = append(files, file)
			}
			return nil
		}

		err := afero.Walk(m.getServices().fs, m.Data.TargetPath, walkFunc)
		if err != nil {
			m.getServices().logger.Fatalf("Something went wrong visiting the folder %s. Are you sure it exists?", m.Data.TargetPath)
		}

		migrationTriggers := []string{
			patterns[iwebpagemedataImport],
			patterns[jsonLdCurrentTitle],
			patterns[jsonLdWebsiteData],
			patterns[svelteKitPrefetch],
		}

		for _, file := range files {
			fileContent, err := retrieveFileContent(m.getServices().fs, file)
			if err != nil {
				return err
			}

			if patternsMatched(fileContent, migrationTriggers, findStringMatcher) {
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

func (m *RefactorSvelteFilesTypes) runMigration(content []byte, file string) ([]byte, error) {

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newReplaceIWebPageMedatadaRule(line),
			newReplaceJSONLdWebsiteDataRule(line),
			newReplaceJSONLdCurrentTitleRule(line),
			newReplaceSvelteKitPrefetchRule(line),
		}
		if res, ok := applyMigrationRules(rules); ok {
			lines[i] = res
		} else {
			lines[i] = line
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

func (m *RefactorSvelteFilesTypes) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *RefactorSvelteFilesTypes) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newReplaceIWebPageMedatadaRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[iwebpagemedataImport],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return `SEOWebPageMetadata`
		},
	}
}

func newReplaceJSONLdWebsiteDataRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[jsonLdWebsiteData],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return `data`
		},
	}
}

func newReplaceJSONLdCurrentTitleRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[jsonLdCurrentTitle],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return `current`
		},
	}
}

func newReplaceSvelteKitPrefetchRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[svelteKitPrefetch],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return `data-sveltekit-preload-data="hover"`
		},
	}
}

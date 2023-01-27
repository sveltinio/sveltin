/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// UnhandledMigration is the struct representing the migration update the defaults.js.ts file.
type UnhandledMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *UnhandledMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UnhandledMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UnhandledMigration) getServices() *MigrationServices { return m.Services }
func (m *UnhandledMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails (IMigration interface).
func (m UnhandledMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UnhandledMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := afero.DirExists(m.getServices().fs, m.Data.TargetPath)
	if err != nil {
		return err
	}

	fileContent, err := retrieveFileContent(m.getServices().fs, path.Join(m.getServices().pathMaker.GetRootFolder(), "package.json"))
	if err != nil {
		return err
	}

	// @sveltinio/essentials min version
	const minEssentialsVersion = 0.5
	// Retrieve @sveltinio/essentials version
	currentEssentialsVersionStr, _ := getDevDependency(fileContent, "@sveltinio/essentials")
	currentEssentialsVersion, _ := versionAsNum(currentEssentialsVersionStr)

	// @sveltinio/widgets min version
	const minWidgetsVersion = 0.5
	// Retrieve @sveltinio/widgets version
	currentWidgetsVersionStr, _ := getDevDependency(fileContent, "@sveltinio/widgets")
	currentWidgetsVersion, _ := versionAsNum(currentWidgetsVersionStr)

	if exists &&
		(currentEssentialsVersion < minEssentialsVersion ||
			currentWidgetsVersion < minWidgetsVersion) {
		files := []string{}
		walkFunc := func(file string, info os.FileInfo, err error) error {
			if filepath.Ext(file) == ".svelte" {
				files = append(files, file)
			}
			return nil
		}

		err := afero.Walk(m.getServices().fs, m.Data.TargetPath, walkFunc)
		if err != nil {
			m.getServices().logger.Fatalf("Something went wrong visiting the folder %s. Are you sure it exists?", m.Data.TargetPath)
		}

		migrationTriggers := []string{
			patterns[essentialsImport],
			patterns[widgetsImport],
		}

		for _, file := range files {
			fileContent, err := retrieveFileContent(m.getServices().fs, file)
			if err != nil {
				return err
			}

			if isMigrationRequired(fileContent, migrationTriggers, findStringMatcher) {
				if !bytes.Contains(fileContent, []byte("[sveltin migrate] @IMPORTANT")) {
					localFilePath :=
						strings.Replace(file, m.getServices().pathMaker.GetRootFolder(), "", 1)
					m.getServices().logger.Info(fmt.Sprintf("Migrating %s", localFilePath))
					if _, err := m.migrate(fileContent, file); err != nil {
						return err
					}
				}
			}

		}

	}

	return nil
}

func (m *UnhandledMigration) migrate(content []byte, file string) ([]byte, error) {

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		rules := []*migrationRule{
			newEssentialsImportRule(line),
			newWidgetsImportRule(line),
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

func (m *UnhandledMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UnhandledMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newEssentialsImportRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[essentialsImport],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			message := `
	/**
	 * ! [sveltin migrate] @IMPORTANT
	 * We detected usage of components from @sveltinio/essentials.
	 *
	 * Latest versions of the package introduced changes the components interfaces.
	 *
	 * Check the updated documentation page and reflect the changes:
	 * https://github.com/sveltinio/components-library/tree/main/packages/essentials
	 */
`
			var sb strings.Builder
			sb.WriteString(message)
			sb.WriteString(line)
			return sb.String()
		},
	}
}

func newWidgetsImportRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[widgetsImport],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			message := `
	/**
	 * ! [sveltin migrate] @IMPORTANT
	 * We detected usage of components from @sveltinio/widgets.
	 *
	 * Latest versions of the package introduced changes the components interfaces.
	 *
	 * Check the updated documentation page and reflect the changes:
	 * https://github.com/sveltinio/components-library/tree/main/packages/widgets
	 */
`
			var sb strings.Builder
			sb.WriteString(message)
			sb.WriteString(line)
			return sb.String()
		},
	}
}

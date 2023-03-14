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

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/utils"
)

var npmPackagesMap = map[string]string{
	"@indaco/svelte-iconoir":   "^3.3.1",
	"@sveltinio/essentials":    "^0.7.0",
	"@sveltinio/media-content": "^0.4.0",
	"@sveltinio/seo":           "^0.4.0",
	"@sveltinio/services":      "^0.4.0",
	"@sveltinio/widgets":       "^0.7.0",
	"@sveltejs/adapter-static": "2.0.1",
	"@sveltejs/kit":            "1.11.0",
	"@types/gtag.js":           "^0.0.12",
	"rimraf":                   "^4.1.2",
	"svelte":                   "^3.55.1",
	"svelte-check":             "^3.0.4",
	"svelte-preprocess":        "^5.0.1",
	"tslib":                    "^2.5.0",
	"typescript":               "^4.9.5",
	"vite":                     "^4.1.4",
}

//=============================================================================

// UpdatePackageJson is the struct representing the migration update the defaults.js.ts file.
type UpdatePackageJson struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdatePackageJson) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdatePackageJson{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// MakeMigration implements IMigration interface.
func (m *UpdatePackageJson) getServices() *MigrationServices { return m.Services }
func (m *UpdatePackageJson) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdatePackageJson) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdatePackageJson) up() error {
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

		updatedContent := fileContent
		migrationTriggers := []string{
			patterns[remarkExtLinks],
			patterns[remarkSlug],
		}

		isMigrate := patternsMatched(fileContent, migrationTriggers, findStringMatcher)
		if isMigrate {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if updatedContent, err = m.runMigration(updatedContent, ""); err != nil {
				return err
			}
		}

		var updateVersion = false
		for name, nextVersion := range npmPackagesMap {
			currentVersion, ok := getDevDependency(fileContent, name)

			if ok && !isEqual(currentVersion, nextVersion) {
				m.getServices().logger.Info(fmt.Sprintf("Bump %s to %s", name, nextVersion))
				updateVersion = true

				updatedContent, err = utils.SetJsonStringValue(
					updatedContent,
					fmt.Sprintf("devDependencies.%s", name),
					nextVersion,
				)
				if err != nil {
					return err
				}
			}
		}

		if isMigrate || updateVersion {
			m.getServices().logger.Important(markup.Purple("Remember to run: sveltin install (or npm run install, pnpm install ...)"))
		}

		// save new package.json file
		if err = overwriteFile(m, updatedContent); err != nil {
			return err
		}
	}

	return nil
}

func (m *UpdatePackageJson) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdatePackageJson) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *UpdatePackageJson) runMigration(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newRemarkExternalLinksRule(line),
			newRemarkSlugRule(line),
			newRemoveMdastUtilToString(line),
			newRemoveUnistUtilVisit(line),
		}
		if res, ok := applyMigrationRules(rules); ok {
			lines[i] = res
		} else {
			lines[i] = line
		}
	}
	output := strings.Join(lines, "\n")
	return []byte(output), nil
}

//=============================================================================

func newRemarkExternalLinksRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[remarkExtLinks],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "\"rehype-external-links\":\"^2.0.1\","
		},
	}
}

func newRemarkSlugRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[remarkSlug],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "\"@sveltinio/remark-headings\":\"^1.0.1\","
		},
	}
}

func newRemoveMdastUtilToString(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[mdastUtilToString],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newRemoveUnistUtilVisit(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[unistUtilVisit],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

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
	"github.com/tidwall/sjson"
)

var npmPackagesMap = map[string]string{
	"@indaco/svelte-iconoir":   "^3.2.0",
	"@sveltinio/essentials":    "^0.5.2",
	"@sveltinio/media-content": "^0.3.6",
	"@sveltinio/seo":           "^0.2.1",
	"@sveltinio/services":      "^0.3.2",
	"@sveltinio/widgets":       "^0.5.1",
	"@sveltejs/adapter-static": "1.0.5",
	"@sveltejs/kit":            "1.3.2",
	"@types/gtag.js":           "^0.0.12",
	"rimraf":                   "^4.1.2",
	"svelte-check":             "^3.0.3",
	"svelte-preprocess":        "^5.0.1",
	"tslib":                    "^2.5.0",
	"typescript":               "^4.9.4",
	"vite":                     "^4.0.4",
}

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

		migrationTriggers := []string{patterns[remarkExtLinks], patterns[remarkSlug]}
		isMigrate := isMigrationRequired(fileContent, migrationTriggers, findStringMatcher)
		if isMigrate {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			if updatedContent, err = m.migrate(updatedContent, ""); err != nil {
				return err
			}
		}

		updateVersion := false
		for name, nextVersion := range npmPackagesMap {
			currentVersion, ok := getDevDependency(fileContent, name)
			if ok && !isEqual(currentVersion, nextVersion) {
				updateVersion = true
				m.getServices().logger.Info(fmt.Sprintf("Bump %s to %s", name, nextVersion))
				if updatedContent, err = updateDevDependency(m, updatedContent, name, nextVersion); err != nil {
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

func (m *UpdatePkgJSONMigration) migrate(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newRemarkExternalLinksRule(line),
			newRemarkSlugRule(line),
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

//=============================================================================

func updateDevDependency(m *UpdatePkgJSONMigration, content []byte, name, value string) ([]byte, error) {
	return sjson.SetBytes(content, fmt.Sprintf("devDependencies.%s", name), value)
}
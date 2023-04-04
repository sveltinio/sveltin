/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/utils"
)

// CleanDotEnv is the struct representing the migration update the defaults.js.ts file.
type CleanDotEnv struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *CleanDotEnv) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &CleanDotEnv{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *CleanDotEnv) getServices() *MigrationServices { return m.Services }
func (m *CleanDotEnv) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m CleanDotEnv) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *CleanDotEnv) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := utils.FileExists(m.getServices().fs, m.Data.TargetPath)
	if err != nil {
		return err
	}

	if exists {
		fileContent, err := retrieveFileContent(m.getServices().fs, m.getData().TargetPath)
		if err != nil {
			return err
		}

		migrationTriggers := []string{
			patterns[sitemapProp],
			patterns[svelteKitBuildFolder],
			patterns[svelteKitBuildComment],
		}
		if patternsMatched(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Infof("Migrating %s", filepath.Base(m.Data.TargetPath))
			if _, err := m.runMigration(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *CleanDotEnv) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *CleanDotEnv) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *CleanDotEnv) runMigration(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*migrationRule{
			newDotEnvSitemapRule(line),
			newDotEnvSvelteKitBuildCommentRule(line),
			newDotEnvSveltekitRule(line),
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

	cleanedOutput := removeMultiEmptyLines(output)
	if err = afero.WriteFile(m.getServices().fs, m.Data.TargetPath, cleanedOutput, 0644); err != nil {
		return nil, err
	}
	return nil, nil
}

//=============================================================================

func newDotEnvSitemapRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[sitemapProp],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newDotEnvSveltekitRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[svelteKitBuildFolder],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

func newDotEnvSvelteKitBuildCommentRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[svelteKitBuildComment],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return ""
		},
	}
}

// =============================================================================

func removeMultiEmptyLines(content string) []byte {
	rule := regexp.MustCompile(`^\n{1,}$`)
	output := rule.ReplaceAllString(strings.TrimSpace(content), "\n")
	return []byte(output)
}

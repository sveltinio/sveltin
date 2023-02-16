/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/utils"
)

// RefactorWebSiteTSTypes is the struct representing the migration update the defaults.js.ts file.
type RefactorWebSiteTSTypes struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface,
func (m *RefactorWebSiteTSTypes) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &RefactorWebSiteTSTypes{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *RefactorWebSiteTSTypes) getServices() *MigrationServices { return m.Services }
func (m *RefactorWebSiteTSTypes) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m RefactorWebSiteTSTypes) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *RefactorWebSiteTSTypes) up() error {
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
			patterns[importIWebSiteSeoType],
			patterns[iwebsiteSeoTypeUsage],
			patterns[keywordsProp],
			patterns[sitemap],
			patterns[webmasterProp],
			patterns[contactEmailProp],
		}
		if patternsMatched(fileContent, migrationTriggers, findStringMatcher) {
			localFilePath :=
				strings.Replace(m.Data.TargetPath, m.getServices().pathMaker.GetRootFolder(), "", 1)
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", localFilePath))
			if _, err := m.runMigration(fileContent, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *RefactorWebSiteTSTypes) runMigration(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var prevLine string
		if i > 0 {
			prevLine = lines[i-1]
		}
		rules := []*migrationRule{
			newWebSiteTSImportRule(line),
			newWebSiteTSUsageRule(line),
			replaceKeywordsPropRule(line),
			addCommentToSitemapPropRule(line),
			replaceWebmasterPropRule(line),
			replaceContactEmailPropRule(line, prevLine),
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

func (m *RefactorWebSiteTSTypes) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *RefactorWebSiteTSTypes) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newWebSiteTSImportRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[importIWebSiteSeoType],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			return "import type { Sveltin } from '$sveltin';"
		},
	}
}

func newWebSiteTSUsageRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[iwebsiteSeoTypeUsage],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "Sveltin.WebSite"
		},
	}
}

func replaceKeywordsPropRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[keywordsProp],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			splitted := strings.Split(line, ":")
			key, value := splitted[0], splitted[1]

			pattern := regexp.MustCompile(`[',"].*[',"]`)
			match := pattern.FindStringSubmatch(value)
			if len(match) == 1 {
				return fmt.Sprintf("%s: %s,", key, utils.ConvertJSStringToStringArray(value))
			}
			return line
		},
	}
}

func addCommentToSitemapPropRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[sitemap],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			message := `
	/**
	 * ! [sveltin migrate] @IMPORTANT
	 * sitemap has been moved as prop out from WebSite types.
	 *
	 * It is now configured in sveltin.json file. Reflect your sitemap config there.
	 */
`
			var sb strings.Builder
			sb.WriteString(message)
			sb.WriteString(line)
			return sb.String()
		},
	}
}

func replaceWebmasterPropRule(line string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[webmasterProp],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			return "creator"
		},
	}
}

func replaceContactEmailPropRule(line, prevLine string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[contactEmailProp],
		replaceFullLine: false,
		replacerFunc: func(string) string {
			prevLinePattern := regexp.MustCompile(`\baddress\b`)
			match := prevLinePattern.FindStringSubmatch(prevLine)
			if len(match) == 1 {
				return "email"
			}
			return "contactEmail"
		},
	}
}

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
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/utils"
)

// UpdateMDsveXPlugins is the struct representing the migration update the defaults.js.ts file.
type UpdateMDsveXPlugins struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *UpdateMDsveXPlugins) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &UpdateMDsveXPlugins{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// implements IMigration interface.
func (m *UpdateMDsveXPlugins) getServices() *MigrationServices { return m.Services }
func (m *UpdateMDsveXPlugins) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails (IMigration interface).
func (m UpdateMDsveXPlugins) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXPlugins) up() error {
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

		gatekeeper := "rehypeExternalLinks"
		migrationTriggers := []string{
			patterns[remarkExtLinksImport],
			patterns[remarkSlugImport],
			patterns[headingsImport],
			patterns[remarkSlugUsage],
			patterns[remarkExtLinksUsage],
			patterns[rehypeSlugUsage],
		}

		if mustMigrate(fileContent, gatekeeper) && patternsMatched(fileContent, migrationTriggers, findStringMatcher) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			updatedContent := []byte(fixRehypeAutoLinkHeadingsUsage(fileContent))
			if _, err := m.runMigration(updatedContent, ""); err != nil {
				return err
			}

		}
	}

	return nil
}

func (m *UpdateMDsveXPlugins) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXPlugins) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXPlugins) runMigration(content []byte, file string) ([]byte, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var prevLine string
		if i > 0 {
			prevLine = lines[i-1]
		}
		rules := []*migrationRule{
			newReplaceHeadingsImportStrRule(line),
			newReplaceRemarkExternalLinksImportStrRule(line),
			newReplaceRemarkExternalLinksUsageRule(line),
			newReplaceRemarkSlugImportStrRule(line),
			newReplaceRemarkSlugUsageRule(line),
			newReplaceRehypePluginUsageRule(line),
			newReplaceRehypeSlugUsageRule(line, prevLine),
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

func newReplaceRehypeSlugUsageRule(line, prevLine string) *migrationRule {
	return &migrationRule{
		value:           line,
		trigger:         patterns[rehypeSlugUsage],
		replaceFullLine: true,
		replacerFunc: func(string) string {
			if strings.Contains(prevLine, "'noopener', 'noreferrer'") {
				return `rehypeSlug,
				[`
			}
			return line
		},
	}
}

//=============================================================================

func fixRehypeAutoLinkHeadingsUsage(content []byte) []byte {
	data := string(content)
	var start string
	var end string

	reStart := regexp.MustCompile(`\(rehypeAutoLinkHeadings`)
	matchStart := reStart.FindStringSubmatch(data)
	if !utils.IsEmptySlice(matchStart) {
		start = matchStart[0]
	}

	reEnd := regexp.MustCompile(`(behavior:\s+'wrap'\s+\}\))`)
	matchEnd := reEnd.FindStringSubmatch(data)
	if !utils.IsEmptySlice(matchEnd) {
		end = matchEnd[0]
	}

	if !utils.IsEmpty(start) && !utils.IsEmpty(end) {
		newStr := "rehypeAutoLinkHeadings, { behavior: 'wrap' }"
		updatedContent := replaceTextInBetween(data, newStr, start, end)
		return []byte(updatedContent)
	}

	return []byte("")
}

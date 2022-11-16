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
	"github.com/sveltinio/sveltin/internal/fsm"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/yinlog"
)

// Patterns used by MigrationRule
const (
	remarkExternalLinksImportStrPattern = `^import remarkExternalLinks`
	remarkExternalLinksUsagePattern     = `\[remarkExternalLinks`
	rehypePluginPattern                 = `rehypePlugins:[\t\s]+\[`
)

//=============================================================================

// UpdateMDsveXMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateMDsveXMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

func (m *UpdateMDsveXMigration) getFs() afero.Fs { return m.Fs }
func (m *UpdateMDsveXMigration) getPathMaker() *pathmaker.SveltinPathMaker {
	return m.PathMaker
}
func (m *UpdateMDsveXMigration) getLogger() *yinlog.Logger { return m.Logger }
func (m *UpdateMDsveXMigration) getData() *MigrationData   { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m UpdateMDsveXMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXMigration) up() error {

	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	if exists {
		if fileContent, ok := isMigrationRequired(m, remarkExternalLinksImportStrPattern, findStringMatcher); ok {
			m.Logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
			if err := updateMDsveXFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateMDsveXMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMDsveXMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updateMDsveXFile(m *UpdateMDsveXMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*MigrationRule{
			newReplaceRemarkExternalLinksImportStrRule(line),
			newReplaceRemarkExternalLinksUsageRule(line),
			newReplaceRehypePluginUsageRule(line),
		}
		if res, ok := applyMigrationRules(rules); ok {
			lines[i] = res
		} else {
			lines[i] = line
		}
	}
	output := strings.Join(lines, "\n")
	err := m.Fs.Remove(m.Data.PathToFile)
	if err != nil {
		return err
	}

	if err = afero.WriteFile(m.Fs, m.Data.PathToFile, []byte(output), 0644); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func newReplaceRemarkExternalLinksImportStrRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           remarkExternalLinksImportStrPattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return "import rehypeExternalLinks from 'rehype-external-links';"
		},
	}
}

func newReplaceRemarkExternalLinksUsageRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           remarkExternalLinksUsagePattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return ""
		},
	}
}

func newReplaceRehypePluginUsageRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           rehypePluginPattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return `rehypePlugins: [
		[rehypeExternalLinks, { target: '_blank', rel: ['noopener', 'noreferrer'] }],`
		},
	}
}

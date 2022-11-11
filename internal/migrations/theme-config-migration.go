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

// UpdateThemeConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateThemeConfigMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

func (m *UpdateThemeConfigMigration) getFs() afero.Fs { return m.Fs }
func (m *UpdateThemeConfigMigration) getPathMaker() *pathmaker.SveltinPathMaker {
	return m.PathMaker
}
func (m *UpdateThemeConfigMigration) getLogger() *yinlog.Logger { return m.Logger }
func (m *UpdateThemeConfigMigration) getData() *MigrationData   { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m UpdateThemeConfigMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateThemeConfigMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	if exists {
		if fileContent, ok := isMigrationRequired(m, "const config = {", findStringMatcher); ok {
			m.Logger.Info(fmt.Sprintf("Migrating %s file", filepath.Base(m.Data.PathToFile)))
			if err := updateThemeFile(m, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateThemeConfigMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateThemeConfigMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func updateThemeFile(m *UpdateThemeConfigMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var prevLine string
		if i > 0 {
			prevLine = lines[i-1]
		}

		rules := []*MigrationRule{newConstNameRule(line), newExportLineRule(line), newThemeNameRule(line, prevLine)}
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

func newConstNameRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           "const config = {",
		IsReplaceFullLine: false,
		GetMatchReplacer: func(string) string {
			return `import { theme } from '../../sveltin.json';

const themeConfig = {`
		},
	}
}

func newExportLineRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           "export default config",
		IsReplaceFullLine: false,
		GetMatchReplacer: func(string) string {
			return "export { themeConfig }"
		},
	}
}

func newThemeNameRule(line, prevLine string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           "name:",
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			if !strings.Contains(prevLine, "author:") && strings.Contains(line, "name:") {
				return "\tname: theme.name,"
			}
			return line
		},
	}
}

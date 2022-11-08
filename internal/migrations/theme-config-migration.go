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
		pattern := regexp.MustCompile("const config = {")
		if isThemeConfigMigrationRequired(m, pattern) {
			m.Logger.Info(fmt.Sprintf("Migrating %s file", filepath.Base(m.Data.PathToFile)))
			if err := updateThemeFile(m); err != nil {
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

func isThemeConfigMigrationRequired(m *UpdateThemeConfigMigration, pattern *regexp.Regexp) bool {
	content, err := afero.ReadFile(m.Fs, m.Data.PathToFile)
	if err != nil {
		return false
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		matches := pattern.FindString(line)
		if len(matches) > 0 {
			return true
		}
	}
	return false
}

func updateThemeFile(m *UpdateThemeConfigMigration) error {
	content, err := afero.ReadFile(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var prevLine string
		if i != 0 {
			prevLine = line
		}
		if strings.Contains(line, "const config = {") {
			newLineContent := `import { theme } from '../../sveltin.config.json';

const themeConfig = {`
			lines[i] = newLineContent
		} else if strings.Contains(line, "name:") && !strings.Contains(prevLine, "author") {
			lines[i] = "  name: theme.name,"
		} else if strings.Contains(line, "export default config") {
			lines[i] = "export { themeConfig }"
		}
	}

	output := strings.Join(lines, "\n")
	err = m.Fs.Remove(m.Data.PathToFile)
	if err != nil {
		return err
	}

	if err = afero.WriteFile(m.Fs, m.Data.PathToFile, []byte(output), 0644); err != nil {
		return err
	}
	return nil
}

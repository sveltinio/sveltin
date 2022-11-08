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

// UpdateDefaultsConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateDefaultsConfigMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

// Execute return error if migration execution over up and down methods fails.
func (m UpdateDefaultsConfigMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDefaultsConfigMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	if exists {
		// regex for semantic versioning - https://ihateregex.io/expr/semver/
		pattern := regexp.MustCompile(`(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)

		if isDefaultsConfigMigrationRequired(m, pattern) {
			m.Logger.Info(fmt.Sprintf("Migrating %s file", filepath.Base(m.Data.PathToFile)))
			if err := updateConfigFile(m, pattern); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateDefaultsConfigMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDefaultsConfigMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func isDefaultsConfigMigrationRequired(m *UpdateDefaultsConfigMigration, pattern *regexp.Regexp) bool {
	content, err := afero.ReadFile(m.Fs, m.Data.PathToFile)
	if err != nil {
		return false
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		matches := pattern.FindStringSubmatch(line)
		if len(matches) > 0 && matches[1] != m.Data.CliVersion {
			return true
		}
	}
	return false
}

func updateConfigFile(m *UpdateDefaultsConfigMigration, pattern *regexp.Regexp) error {
	content, err := afero.ReadFile(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if pattern.MatchString(line) {
			newContent := `import { sveltin } from '../sveltin.config.json';

const sveltinVersion = sveltin.version;`
			lines[i] = newContent
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

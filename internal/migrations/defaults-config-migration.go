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

// SemVersionRegExp is the regexp pattern for semantic versioning - https://ihateregex.io/expr/semver/
const SemVersionRegExp = `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`

// UpdateDefaultsConfigMigration is the struct representing the migration update the defaults.js.ts file.
type UpdateDefaultsConfigMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

func (m *UpdateDefaultsConfigMigration) getFs() afero.Fs { return m.Fs }
func (m *UpdateDefaultsConfigMigration) getPathMaker() *pathmaker.SveltinPathMaker {
	return m.PathMaker
}
func (m *UpdateDefaultsConfigMigration) getLogger() *yinlog.Logger { return m.Logger }
func (m *UpdateDefaultsConfigMigration) getData() *MigrationData   { return m.Data }

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
		if fileContent, ok := isMigrationRequired(m, SemVersionRegExp, testAsOne); ok {
			m.Logger.Info(fmt.Sprintf("Migrating %s file", filepath.Base(m.Data.PathToFile)))
			if err := updateConfigFile(m, fileContent); err != nil {
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

func updateConfigFile(m *UpdateDefaultsConfigMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*MigrationRule{newSveltinVersionRule(line)}
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

func newSveltinVersionRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           SemVersionRegExp,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return `import { sveltin } from '../sveltin.json';

const sveltinVersion = sveltin.version;`
		},
	}
}

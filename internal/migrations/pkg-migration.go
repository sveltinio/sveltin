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
	remarkAutolinkHeadingsPkgPattern = `"remark-external-links"`
)

//=============================================================================

// UpdatePkgJSONMigration is the struct representing the migration update the defaults.js.ts file.
type UpdatePkgJSONMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

func (m *UpdatePkgJSONMigration) getFs() afero.Fs { return m.Fs }
func (m *UpdatePkgJSONMigration) getPathMaker() *pathmaker.SveltinPathMaker {
	return m.PathMaker
}
func (m *UpdatePkgJSONMigration) getLogger() *yinlog.Logger { return m.Logger }
func (m *UpdatePkgJSONMigration) getData() *MigrationData   { return m.Data }

// Execute return error if migration execution over up and down methods fails.
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

	exists, err := common.FileExists(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	if exists {
		if fileContent, ok := isMigrationRequired(m, remarkAutolinkHeadingsPkgPattern, findStringMatcher); ok {
			m.Logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.PathToFile)))
			if err := updatePkgJSONFile(m, fileContent); err != nil {
				return err
			}
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

//=============================================================================

func updatePkgJSONFile(m *UpdatePkgJSONMigration, content []byte) error {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		rules := []*MigrationRule{
			newRemarkExternalLinksRule(line),
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

func newRemarkExternalLinksRule(line string) *MigrationRule {
	return &MigrationRule{
		Value:             line,
		Pattern:           remarkAutolinkHeadingsPkgPattern,
		IsReplaceFullLine: true,
		GetMatchReplacer: func(string) string {
			return "\"rehype-external-links\":\"^2.0.1\","
		},
	}
}

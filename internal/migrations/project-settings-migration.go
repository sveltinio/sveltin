/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/fsm"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
	"github.com/sveltinio/yinlog"
)

// AddProjectSettingsMigration is the struct representing the migration add the sveltin.json file.
type AddProjectSettingsMigration struct {
	Mediator  MigrationMediator
	Fs        afero.Fs
	FsManager *fsm.SveltinFSManager
	PathMaker *pathmaker.SveltinPathMaker
	Logger    *yinlog.Logger
	Data      *MigrationData
}

func (m *AddProjectSettingsMigration) getFs() afero.Fs { return m.Fs }
func (m *AddProjectSettingsMigration) getPathMaker() *pathmaker.SveltinPathMaker {
	return m.PathMaker
}
func (m *AddProjectSettingsMigration) getLogger() *yinlog.Logger { return m.Logger }
func (m *AddProjectSettingsMigration) getData() *MigrationData   { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m AddProjectSettingsMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *AddProjectSettingsMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, _ := common.FileExists(m.Fs, m.Data.PathToFile)
	if !exists {
		m.Logger.Info(fmt.Sprintf("Creating %s", filepath.Base(m.Data.PathToFile)))
		return addProjectSettingsFile(m)
	} else if exists && m.Data.ProjectCliVersion != m.Data.CliVersion {
		m.Logger.Info(fmt.Sprintf("Bumping Sveltin CLI version in %s", filepath.Base(m.Data.PathToFile)))
		return updateFileContent(m)
	}

	return nil
}

func (m *AddProjectSettingsMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *AddProjectSettingsMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

//=============================================================================

func addProjectSettingsFile(m *AddProjectSettingsMigration) error {
	pathToPkgFile := filepath.Join(m.PathMaker.GetRootFolder(), "package.json")

	projectName, err := utils.RetrieveProjectName(m.Fs, pathToPkgFile)
	if err != nil {
		return err
	}

	cssLibName, err := utils.RetrieveCSSLib(m.Fs, pathToPkgFile)
	if err != nil {
		return err
	}

	themeData, err := makeThemeData(m)
	if err != nil {
		return err
	}
	themeData.CSSLib = cssLibName

	// NEW FILE: .sveltin.json
	sveltinConfigTplData := &config.TemplateData{
		Name: "sveltin.json",
		ProjectSettings: &tpltypes.ProjectSettings{
			Name:    projectName,
			BaseURL: fmt.Sprintf("http://%s.com", projectName),
			Sitemap: tpltypes.SitemapData{
				ChangeFreq: "monthly",
				Priority:   0.5,
			},
			Sveltin: tpltypes.SveltinCLIData{
				Version: m.Data.CliVersion,
			},
			Theme: *themeData,
		},
	}
	sveltinJSONConfigFile := m.FsManager.NewJSONConfigFile(sveltinConfigTplData)

	cwd, _ := os.Getwd()
	projectFolderName := filepath.Base(cwd)
	projectFolder := m.FsManager.GetFolder(projectFolderName)
	projectFolder.Add(sveltinJSONConfigFile)

	sfs := factory.NewProjectArtifact(&resources.SveltinFS, m.Fs)
	err = projectFolder.Create(sfs)
	if err != nil {
		return err
	}
	return nil
}

func makeThemeData(m *AddProjectSettingsMigration) (*tpltypes.ThemeData, error) {
	const (
		blankThemeId   string = "blank"
		sveltinThemeId string = "sveltin"
	)

	themeData := &tpltypes.ThemeData{}
	files, err := afero.ReadDir(m.Fs, m.PathMaker.GetThemesFolder())
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			if strings.HasPrefix(file.Name(), sveltinThemeId) {
				themeData.ID = sveltinThemeId
			} else {
				themeData.ID = blankThemeId
			}
			themeData.Name = file.Name()
		}
	}
	return themeData, nil
}

func updateFileContent(m *AddProjectSettingsMigration) error {
	content, err := afero.ReadFile(m.Fs, m.Data.PathToFile)
	if err != nil {
		return err
	}

	newContent := bytes.Replace(content, []byte(m.Data.ProjectCliVersion), []byte(m.Data.CliVersion), -1)

	if err = afero.WriteFile(m.Fs, m.Data.PathToFile, newContent, 0666); err != nil {
		return err
	}
	return nil
}

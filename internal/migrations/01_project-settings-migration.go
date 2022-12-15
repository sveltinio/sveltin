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
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

// ProjectSettingsMigration is the struct representing the migration add the sveltin.json file.
type ProjectSettingsMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *ProjectSettingsMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &ProjectSettingsMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// MakeMigration implements IMigration interface.
func (m *ProjectSettingsMigration) getServices() *MigrationServices { return m.Services }
func (m *ProjectSettingsMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m ProjectSettingsMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *ProjectSettingsMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, _ := common.FileExists(m.getServices().fs, m.Data.FileToMigrate)
	if !exists {
		m.getServices().logger.Info(fmt.Sprintf("Creating %s", filepath.Base(m.Data.FileToMigrate)))
		return addProjectSettingsFile(m)
	} else if exists && m.Data.ProjectCliVersion != m.Data.CliVersion {
		m.getServices().logger.Info(fmt.Sprintf("Bumping Sveltin CLI version in %s", filepath.Base(m.Data.FileToMigrate)))
		return updateFileContent(m)
	}

	return nil
}

func (m *ProjectSettingsMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *ProjectSettingsMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *ProjectSettingsMigration) migrate(content []byte) ([]byte, error) {
	return nil, nil
}

//=============================================================================

func addProjectSettingsFile(m *ProjectSettingsMigration) error {
	pathToPkgFile := filepath.Join(m.getServices().pathMaker.GetRootFolder(), "package.json")

	projectName, err := utils.RetrieveProjectName(m.getServices().fs, pathToPkgFile)
	if err != nil {
		return err
	}

	cssLibName, err := utils.RetrieveCSSLib(m.getServices().fs, pathToPkgFile)
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
	sveltinJSONConfigFile := m.getServices().fsManager.NewJSONConfigFile(sveltinConfigTplData)

	cwd, _ := os.Getwd()
	projectFolderName := filepath.Base(cwd)
	projectFolder := m.getServices().fsManager.GetFolder(projectFolderName)
	projectFolder.Add(sveltinJSONConfigFile)

	sfs := factory.NewProjectArtifact(&resources.SveltinTemplatesFS, m.getServices().fs)
	err = projectFolder.Create(sfs)
	if err != nil {
		return err
	}
	return nil
}

func makeThemeData(m *ProjectSettingsMigration) (*tpltypes.ThemeData, error) {
	const (
		blankThemeId   string = "blank"
		sveltinThemeId string = "sveltin"
	)

	themeData := &tpltypes.ThemeData{}
	files, err := afero.ReadDir(m.getServices().fs, m.getServices().pathMaker.GetThemesFolder())
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

func updateFileContent(m *ProjectSettingsMigration) error {
	content, err := afero.ReadFile(m.getServices().fs, m.Data.FileToMigrate)
	if err != nil {
		return err
	}

	newContent := bytes.Replace(content, []byte(m.Data.ProjectCliVersion), []byte(m.Data.CliVersion), -1)

	if err = afero.WriteFile(m.getServices().fs, m.Data.FileToMigrate, newContent, 0666); err != nil {
		return err
	}
	return nil
}

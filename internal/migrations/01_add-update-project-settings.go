/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
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

// AddUpdateProjectSettings is the struct representing the migration add the sveltin.json file.
type AddUpdateProjectSettings struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *AddUpdateProjectSettings) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &AddUpdateProjectSettings{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// MakeMigration implements IMigration interface.
func (m *AddUpdateProjectSettings) getServices() *MigrationServices { return m.Services }
func (m *AddUpdateProjectSettings) getData() *MigrationData         { return m.Data }

// Migrate return error if migration execution over up and down methods fails.
func (m AddUpdateProjectSettings) Migrate() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *AddUpdateProjectSettings) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, _ := common.FileExists(m.getServices().fs, m.Data.TargetPath)
	if !exists {
		m.getServices().logger.Info(fmt.Sprintf("Creating %s", filepath.Base(m.Data.TargetPath)))
		return addProjectSettingsFile(m)
	} else if exists && m.Data.ProjectCliVersion != m.Data.CliVersion {
		m.getServices().logger.Info(fmt.Sprintf("Bumping Sveltin CLI version in %s", filepath.Base(m.Data.TargetPath)))
		return updateFileContent(m)
	}

	return nil
}

func (m *AddUpdateProjectSettings) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *AddUpdateProjectSettings) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *AddUpdateProjectSettings) runMigration(content []byte, file string) ([]byte, error) {
	return nil, nil
}

//=============================================================================

func addProjectSettingsFile(m *AddUpdateProjectSettings) error {
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

func makeThemeData(m *AddUpdateProjectSettings) (*tpltypes.ThemeData, error) {
	themeData := &tpltypes.ThemeData{}
	files, err := afero.ReadDir(m.getServices().fs, m.getServices().pathMaker.GetThemesFolder())
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			if strings.HasPrefix(file.Name(), tpltypes.SveltinTheme) {
				themeData.ID = tpltypes.SveltinTheme
			} else {
				themeData.ID = tpltypes.BlankTheme
			}
			themeData.Name = file.Name()
		}
	}
	return themeData, nil
}

func updateFileContent(m *AddUpdateProjectSettings) error {
	content, err := afero.ReadFile(m.getServices().fs, m.Data.TargetPath)
	if err != nil {
		return err
	}

	newContent, err := utils.SetJsonStringValue(content, "sveltin.version", m.Data.CliVersion)
	if err != nil {
		return err
	}

	if err = afero.WriteFile(m.getServices().fs, m.Data.TargetPath, newContent, 0666); err != nil {
		return err
	}
	return nil
}

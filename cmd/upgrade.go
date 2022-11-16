/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/migrations"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

// Migration identifiers.
const (
	ProjectSettingsMigrationID string = "projectSettings"
	DefaultsConfigMigrationID  string = "defaultsConfig"
	ThemeConfigMigrationID     string = "themeConfig"
	DotEnvMigrationID          string = "dotenv"
	PackageJSONID              string = "packagejson"
	MDsveXID                   string = "mdsvex"
)

//=============================================================================

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your project to the latest Sveltin version.",
	Long: resources.GetASCIIArt() + `
Command used to upgrade your existing project to be compliant
with the latest Sveltin version.

`,
	Run: RunUpgradeCmd,
}

// RunUpgradeCmd is the actual work function.
func RunUpgradeCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject(false)
	cwd, _ := os.Getwd()

	cfg.log.Plain(markup.H1(fmt.Sprintf("Upgrading your project to sveltin v%s", CliVersion)))
	migrationManager := migrations.NewMigrationManager()

	// FILE: <project_root>/sveltin.json
	pathToFile := path.Join(cwd, ProjectSettingsFile)
	migration := handleMigration(ProjectSettingsMigrationID, migrationManager, cfg, pathToFile)
	err := migration.Execute()
	utils.ExitIfError(err)

	// FILE: <project_root>/config/defaults.js.ts
	pathToFile = path.Join(cwd, cfg.pathMaker.GetConfigFolder(), DefaultsConfigFile)
	migration = handleMigration(DefaultsConfigMigrationID, migrationManager, cfg, pathToFile)
	err = migration.Execute()
	utils.ExitIfError(err)

	// FILE: <project_root>/themes/<theme_name>/theme.config.js
	cfg.projectSettings, err = loadProjectSettings(ProjectSettingsFile)
	utils.ExitIfError(err)
	pathToFile = path.Join(cwd, cfg.pathMaker.GetThemesFolder(), cfg.projectSettings.Theme.Name, cfg.settings.GetThemeConfigFilename())
	migration = handleMigration(ThemeConfigMigrationID, migrationManager, cfg, pathToFile)
	err = migration.Execute()
	utils.ExitIfError(err)

	// File: <project_root>/.env.production
	pathToFile = path.Join(cwd, DotEnvProdFile)
	migration = handleMigration(DotEnvMigrationID, migrationManager, cfg, pathToFile)
	err = migration.Execute()
	utils.ExitIfError(err)

	// File: <project_root>/package.json
	pathToFile = path.Join(cwd, PackageJSONFile)
	migration = handleMigration(PackageJSONID, migrationManager, cfg, pathToFile)
	err = migration.Execute()
	utils.ExitIfError(err)

	// File: <project_root/mdsvex.config.json
	pathToFile = path.Join(cwd, MDsveXFile+".config.js")
	migration = handleMigration(MDsveXID, migrationManager, cfg, pathToFile)
	err = migration.Execute()
	utils.ExitIfError(err)

	cfg.log.Success(fmt.Sprintf("Your project is ready for sveltin v%s\n", CliVersion))
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

//=============================================================================

func handleMigration(migrationType string, migrationManager *migrations.MigrationManager, config appConfig, pathToFile string) migrations.Migration {
	switch migrationType {
	case ProjectSettingsMigrationID:
		return newAddProjectSettingsMigration(migrationManager, config, pathToFile)
	case DefaultsConfigMigrationID:
		return newUpdateDefaultsConfigMigration(migrationManager, config, pathToFile)
	case ThemeConfigMigrationID:
		return newUpdateThemeConfigMigration(migrationManager, config, pathToFile)
	case DotEnvMigrationID:
		return newDotEnvMigration(migrationManager, config, pathToFile)
	case PackageJSONID:
		return newPackageJSONMigration(migrationManager, config, pathToFile)
	case MDsveXID:
		return newMDsveXMigration(migrationManager, config, pathToFile)
	default:
		return nil
	}
}

//=============================================================================

func newAddProjectSettingsMigration(migrationManager *migrations.MigrationManager, config appConfig, pathTofile string) *migrations.AddProjectSettingsMigration {
	return &migrations.AddProjectSettingsMigration{
		Mediator:  migrationManager,
		Fs:        config.fs,
		FsManager: config.fsManager,
		PathMaker: config.pathMaker,
		Logger:    config.log,
		Data: &migrations.MigrationData{
			PathToFile:        pathTofile,
			CliVersion:        CliVersion,
			ProjectCliVersion: config.projectSettings.Sveltin.Version,
		},
	}
}

func newUpdateDefaultsConfigMigration(migrationManager *migrations.MigrationManager, config appConfig, pathTofile string) *migrations.UpdateDefaultsConfigMigration {
	return &migrations.UpdateDefaultsConfigMigration{
		Mediator:  migrationManager,
		Fs:        config.fs,
		FsManager: config.fsManager,
		PathMaker: config.pathMaker,
		Logger:    config.log,
		Data: &migrations.MigrationData{
			PathToFile:        pathTofile,
			CliVersion:        CliVersion,
			ProjectCliVersion: config.projectSettings.Sveltin.Version,
		},
	}
}

func newUpdateThemeConfigMigration(migrationManager *migrations.MigrationManager, config appConfig, pathTofile string) *migrations.UpdateThemeConfigMigration {
	return &migrations.UpdateThemeConfigMigration{
		Mediator:  migrationManager,
		Fs:        config.fs,
		FsManager: config.fsManager,
		PathMaker: config.pathMaker,
		Logger:    config.log,
		Data: &migrations.MigrationData{
			PathToFile:        pathTofile,
			CliVersion:        CliVersion,
			ProjectCliVersion: config.projectSettings.Sveltin.Version,
		},
	}
}

func newDotEnvMigration(migrationManager *migrations.MigrationManager, config appConfig, pathTofile string) *migrations.UpdateDotEnvMigration {
	return &migrations.UpdateDotEnvMigration{
		Mediator:  migrationManager,
		Fs:        config.fs,
		FsManager: config.fsManager,
		PathMaker: config.pathMaker,
		Logger:    config.log,
		Data: &migrations.MigrationData{
			PathToFile: pathTofile,
		},
	}
}

func newPackageJSONMigration(migrationManager *migrations.MigrationManager, config appConfig, pathTofile string) *migrations.UpdatePkgJSONMigration {
	return &migrations.UpdatePkgJSONMigration{
		Mediator:  migrationManager,
		Fs:        config.fs,
		FsManager: config.fsManager,
		PathMaker: config.pathMaker,
		Logger:    config.log,
		Data: &migrations.MigrationData{
			PathToFile: pathTofile,
		},
	}
}

func newMDsveXMigration(migrationManager *migrations.MigrationManager, config appConfig, pathTofile string) *migrations.UpdateMDsveXMigration {
	return &migrations.UpdateMDsveXMigration{
		Mediator:  migrationManager,
		Fs:        config.fs,
		FsManager: config.fsManager,
		PathMaker: config.pathMaker,
		Logger:    config.log,
		Data: &migrations.MigrationData{
			PathToFile: pathTofile,
		},
	}
}

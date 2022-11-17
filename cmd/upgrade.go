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
	migrationServices := migrations.NewMigrationServices(cfg.fs, cfg.fsManager, cfg.pathMaker, cfg.log)

	/** FILE: <project_root>/sveltin.json */
	pathToFile := path.Join(cwd, ProjectSettingsFile)
	migrationData := &migrations.MigrationData{
		PathToFile:        pathToFile,
		CliVersion:        CliVersion,
		ProjectCliVersion: cfg.projectSettings.Sveltin.Version,
	}
	migrationFactory, err := migrations.GetMigrationFactory(migrations.ProjectSettingsMigrationID)
	migration := migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
	utils.ExitIfError(err)
	// execute the migration.
	err = migration.Execute()
	utils.ExitIfError(err)

	/** FILE: <project_root>/config/defaults.js.ts */
	pathToFile = path.Join(cwd, cfg.pathMaker.GetConfigFolder(), DefaultsConfigFile)
	migrationData = &migrations.MigrationData{
		PathToFile:        pathToFile,
		CliVersion:        CliVersion,
		ProjectCliVersion: cfg.projectSettings.Sveltin.Version,
	}
	migrationFactory, err = migrations.GetMigrationFactory(migrations.DefaultsConfigMigrationID)
	migration = migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
	utils.ExitIfError(err)
	// execute the migration.
	err = migration.Execute()
	utils.ExitIfError(err)

	// Load project settings file after sveltin.json file creation
	cfg.projectSettings, err = loadProjectSettings(ProjectSettingsFile)
	utils.ExitIfError(err)

	/** FILE: <project_root>/themes/<theme_name>/theme.config.js */
	pathToFile = path.Join(cwd, cfg.pathMaker.GetThemesFolder(), cfg.projectSettings.Theme.Name, cfg.settings.GetThemeConfigFilename())
	migrationData = &migrations.MigrationData{
		PathToFile:        pathToFile,
		CliVersion:        CliVersion,
		ProjectCliVersion: cfg.projectSettings.Sveltin.Version,
	}
	migrationFactory, err = migrations.GetMigrationFactory(migrations.ThemeConfigMigrationID)
	migration = migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
	utils.ExitIfError(err)
	// execute the migration.
	err = migration.Execute()
	utils.ExitIfError(err)

	/** File: <project_root>/.env.production */
	pathToFile = path.Join(cwd, DotEnvProdFile)
	migrationData = &migrations.MigrationData{
		PathToFile: pathToFile,
	}
	migrationFactory, err = migrations.GetMigrationFactory(migrations.DotEnvMigrationID)
	migration = migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
	utils.ExitIfError(err)
	// execute the migration.
	err = migration.Execute()
	utils.ExitIfError(err)

	/** File: <project_root>/package.json */
	pathToFile = path.Join(cwd, PackageJSONFile)
	migrationData = &migrations.MigrationData{
		PathToFile: pathToFile,
	}
	migrationFactory, err = migrations.GetMigrationFactory(migrations.PackageJSONMigrationID)
	migration = migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
	utils.ExitIfError(err)
	// execute the migration.
	err = migration.Execute()
	utils.ExitIfError(err)

	/** File: <project_root/mdsvex.config.json */
	pathToFile = path.Join(cwd, MDsveXFile)
	migrationData = &migrations.MigrationData{
		PathToFile: pathToFile,
	}
	migrationFactory, err = migrations.GetMigrationFactory(migrations.MDsveXMigrationID)
	migration = migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
	utils.ExitIfError(err)
	// execute the migration.
	err = migration.Execute()
	utils.ExitIfError(err)

	cfg.log.Success(fmt.Sprintf("Your project is ready for sveltin v%s\n", CliVersion))
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

package cmd

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/confirm"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/migrations"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/feedbacks"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate your project to the latest Sveltin version",
	Long: resources.GetASCIIArt() + `
Command used to migrate your project files to the latest Sveltin version.
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(0),
	Run:                   RunMigrateCmd,
}

// RunMigrateCmd is the actual work function.
func RunMigrateCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject(false)

	feedbacks.ShowUpgradeCommandMessage()

	isConfirm, err := confirm.Run(&confirm.Config{Question: "Continue?"})
	utils.ExitIfError(err)

	if isConfirm {
		cwd, _ := os.Getwd()
		cfg.log.Plain(markup.H1(fmt.Sprintf("Migrating your project to sveltin v%s", CliVersion)))

		migrationManager := migrations.NewMigrationManager()
		migrationServices := migrations.NewMigrationServices(cfg.fs, cfg.fsManager, cfg.pathMaker, cfg.log)

		/** FILE: <project_root>/sveltin.json */
		pathToFile := path.Join(cwd, ProjectSettingsFile)
		migrationData := &migrations.MigrationData{
			TargetPath:        pathToFile,
			CliVersion:        CliVersion,
			ProjectCliVersion: cfg.projectSettings.Sveltin.Version,
		}
		migrationFactory, err := migrations.GetMigrationFactory(migrations.ProjectSettingsMigrationId)
		utils.ExitIfError(err)
		migration := migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
		// execute the migration.
		err = migration.Execute()
		utils.ExitIfError(err)

		// Load project settings file after sveltin.json file creation
		cfg.projectSettings, err = loadProjectSettings(ProjectSettingsFile)
		utils.ExitIfError(err)

		migrationIdPathToFileMap := map[string]string{
			migrations.PackageJSONMigrationId:    path.Join(cwd, PackageJSONFile),
			migrations.MDsveXMigrationId:         path.Join(cwd, MDsveXFile),
			migrations.SvelteConfigMigrationId:   path.Join(cwd, SvelteConfigFile),
			migrations.DotEnvMigrationId:         path.Join(cwd, DotEnvProdFile),
			migrations.SveltinDTSMigrationId:     path.Join(cwd, cfg.pathMaker.GetSrcFolder(), SveltinDTSFile),
			migrations.LayoutMigrationId:         path.Join(cwd, cfg.pathMaker.GetRoutesFolder(), LayoutTSFile),
			migrations.DefaultsConfigMigrationId: path.Join(cwd, cfg.pathMaker.GetConfigFolder(), DefaultsConfigFile),
			migrations.WebSiteTSMigrationId:      path.Join(cwd, cfg.pathMaker.GetConfigFolder(), WebSiteTSFile),
			migrations.MenuTSMigrationId:         path.Join(cwd, cfg.pathMaker.GetConfigFolder(), MenuTSFile),
			migrations.HeadingsMigrationId:       path.Join(cwd, cfg.pathMaker.GetLibFolder(), "utils", HeadingsJSFile),
			migrations.StringsTSMigrationId:      path.Join(cwd, cfg.pathMaker.GetLibFolder(), "utils", StringsTSFile),
			migrations.ThemeConfigMigrationId:    path.Join(cwd, cfg.pathMaker.GetThemesFolder(), cfg.projectSettings.Theme.Name, cfg.settings.GetThemeConfigFilename()),
		}

		for id, pathToFile := range migrationIdPathToFileMap {
			migrationData := &migrations.MigrationData{
				TargetPath: pathToFile,
			}
			migrationFactory, err := migrations.GetMigrationFactory(id)
			utils.ExitIfError(err)
			migration := migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
			// execute the migration.
			err = migration.Execute()
			utils.ExitIfError(err)
		}

		cfg.log.Success(fmt.Sprintf("Your project is ready for sveltin v%s\n", CliVersion))
	}
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
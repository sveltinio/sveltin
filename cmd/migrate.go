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
	"sort"

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
		migrationFactory, err := migrations.GetMigrationFactory(migrations.ProjectSettings)
		utils.ExitIfError(err)
		migration := migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
		// execute the migration.
		err = migration.Execute()
		utils.ExitIfError(err)

		// Load project settings file after sveltin.json file creation
		cfg.projectSettings, err = loadProjectSettings(ProjectSettingsFile)
		utils.ExitIfError(err)

		migrationIdPathToTargetMap := map[migrations.Migration]string{
			migrations.DefaultsConfig:     path.Join(cwd, cfg.pathMaker.GetConfigFolder(), DefaultsConfigFile),
			migrations.WebSiteTS:          path.Join(cwd, cfg.pathMaker.GetConfigFolder(), WebSiteTSFile),
			migrations.MenuTS:             path.Join(cwd, cfg.pathMaker.GetConfigFolder(), MenuTSFile),
			migrations.SveltinDTS:         path.Join(cwd, cfg.pathMaker.GetSrcFolder(), SveltinDTSFile),
			migrations.ResourceLibs:       path.Join(cwd, cfg.pathMaker.GetLibFolder()),
			migrations.Layout:             path.Join(cwd, cfg.pathMaker.GetRoutesFolder(), LayoutTSFile),
			migrations.SvelteFiles:        path.Join(cwd, cfg.pathMaker.GetRoutesFolder()),
			migrations.PageServerTS:       path.Join(cwd, cfg.pathMaker.GetRoutesFolder()),
			migrations.SveltinioComponent: path.Join(cwd, cfg.pathMaker.GetRoutesFolder()),
			migrations.ThemeConfig: path.Join(cwd, cfg.pathMaker.GetThemesFolder(),
				cfg.projectSettings.Theme.Name, cfg.settings.GetThemeConfigFilename()),
			migrations.ThemeSveltinioComponents: path.Join(cwd, cfg.pathMaker.GetThemesFolder()),
			migrations.MDsveXConfig:             path.Join(cwd, MDsveXFile),
			migrations.SvelteConfig:             path.Join(cwd, SvelteConfigFile),
			migrations.DotEnv:                   path.Join(cwd, DotEnvProdFile),
			migrations.PackageJSON:              path.Join(cwd, PackageJSONFile),
		}

		// Ensure the migrations execution order
		migrationKeys := sortedMigrationMap(migrationIdPathToTargetMap)

		for _, k := range migrationKeys {
			_id := migrations.Migration(k)
			_pathToFile := migrationIdPathToTargetMap[_id]
			migrationData := &migrations.MigrationData{
				TargetPath: _pathToFile,
			}
			migrationFactory, err := migrations.GetMigrationFactory(_id)
			utils.ExitIfError(err)
			migration := migrationFactory.MakeMigration(migrationManager, migrationServices, migrationData)
			// execute the migration.
			err = migration.Execute()
			utils.ExitIfError(err)
		}

		cfg.log.Success(markup.Green(fmt.Sprintf("Your project is ready for sveltin v%s\n", CliVersion)))
	}
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func sortedMigrationMap(m map[migrations.Migration]string) []int {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	return keys
}

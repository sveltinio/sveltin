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

	pathToProjectSettingsFile := path.Join(cwd, ProjectSettingsFile)
	addProjectSettingsMigration := handleMigration(ProjectSettingsMigrationID, migrationManager, pathToProjectSettingsFile)
	err := addProjectSettingsMigration.Execute()
	utils.ExitIfError(err)

	pathToDefaultsConfigFile := path.Join(cwd, cfg.pathMaker.GetConfigFolder(), DefaultsConfigFile)
	updateDefaultsConfigMigration := handleMigration(DefaultsConfigMigrationID, migrationManager, pathToDefaultsConfigFile)
	err = updateDefaultsConfigMigration.Execute()
	utils.ExitIfError(err)

	cfg.log.Success(fmt.Sprintf("Your project is ready for sveltin v%s\n", CliVersion))
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

//=============================================================================

func handleMigration(migrationType string, migrationManager *migrations.MigrationManager, pathToFile string) migrations.Migration {
	switch migrationType {
	case ProjectSettingsMigrationID:
		return newAddProjectSettingsMigration(migrationManager, pathToFile)
	case DefaultsConfigMigrationID:
		return newUpdateDefaultsConfigMigration(migrationManager, pathToFile)
	default:
		return nil
	}
}

func newAddProjectSettingsMigration(migrationManager *migrations.MigrationManager, pathTofile string) *migrations.AddProjectSettingsMigration {
	return &migrations.AddProjectSettingsMigration{
		Mediator:  migrationManager,
		Fs:        cfg.fs,
		FsManager: cfg.fsManager,
		PathMaker: cfg.pathMaker,
		Logger:    cfg.log,
		Data: &migrations.MigrationData{
			PathToFile:        pathTofile,
			CliVersion:        CliVersion,
			ProjectCliVersion: cfg.projectSettings.CLI.Version,
		},
	}
}

func newUpdateDefaultsConfigMigration(migrationManager *migrations.MigrationManager, pathTofile string) *migrations.UpdateDefaultsConfigMigration {
	return &migrations.UpdateDefaultsConfigMigration{
		Mediator:  migrationManager,
		Fs:        cfg.fs,
		FsManager: cfg.fsManager,
		PathMaker: cfg.pathMaker,
		Logger:    cfg.log,
		Data: &migrations.MigrationData{
			PathToFile:        pathTofile,
			CliVersion:        CliVersion,
			ProjectCliVersion: cfg.projectSettings.CLI.Version,
		},
	}
}

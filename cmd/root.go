/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd contains all the commands and subcommands for sveltin.
package cmd

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/fsm"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	logger "github.com/sveltinio/yinlog"
)

//=============================================================================

type appConfig struct {
	log             *logger.Logger
	settings        *config.SveltinSettings
	projectSettings tpltypes.ProjectSettings
	prodData        tpltypes.EnvProductionData
	pathMaker       *pathmaker.SveltinPathMaker
	fsManager       *fsm.SveltinFSManager
	startersMap     map[string]config.StarterTemplate
	fs              afero.Fs
}

//=============================================================================

const (
	// SvelteKitStarter is a string representing the project starter id.
	SvelteKitStarter string = "starter"
	// ThemeStarter is a string representing the project starter id for new themes.
	ThemeStarter string = "blog-theme-starter"
)

// Folder names for a Sveltin project structure.
const (
	RootFolder    string = "root"
	BackupsFolder string = "backups"
	ConfigFolder  string = "config"
	ContentFolder string = "content"
	RoutesFolder  string = "routes"
	ApiFolder     string = "api"
	ParamsFolder  string = "params"
	LibFolder     string = "lib"
	StaticFolder  string = "static"
	ThemesFolder  string = "themes"
)

// File names for a Sveltin project structure.
const (
	StringMatcher       string = "string_matcher"
	GenericMatcher      string = "generic_matcher"
	ApiIndexFile        string = "api_index"
	ApiMetadataIndex    string = "api_metadata_index"
	ApiSlugFile         string = "api_slug"
	IndexFile           string = "index"
	IndexEndpointFile   string = "indexendpoint"
	SlugFile            string = "slug"
	SlugEndpointFile    string = "slugendpoint"
	SlugLayoutFile      string = "sluglayout"
	MDsveXFile          string = "mdsvex"
	ProjectSettingsFile string = "sveltin.json"
	DefaultsConfigFile  string = "defaults.js.ts"
	DotEnvProdFile      string = ".env.production"
)

var (
	// YamlConfig is used by yaml.Unmarshal to decode the YAML file.
	YamlConfig    []byte
	npmClientName string
	cfg           appConfig
)

//=============================================================================

var rootCmd = &cobra.Command{
	Use:              "sveltin",
	Version:          CliVersion,
	TraverseChildren: true,
	Short:            "sveltin is the main command to work with SvelteKit powered static websites.",
	Long: resources.GetASCIIArt() + `
A powerful CLI for your SvelteKit powered static website!

sveltin is the main command used to boost your productivity
while creating a new production-ready SvelteKit project.

Resources:
  Documentation           -> https://docs.sveltin.io
  A helpful quick-start   -> https://docs.sveltin.io/quick-start

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It is called by main.main().
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(loadSveltinSettings, initAppConfig)
}

//=============================================================================

func initAppConfig() {
	cfg.log = logger.New()
	cfg.log.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    false,
		Icons:     true,
	})
	cfg.pathMaker = pathmaker.NewSveltinPathMaker(cfg.settings)
	cfg.fsManager = fsm.NewSveltinFSManager(cfg.pathMaker)
	cfg.startersMap = helpers.InitStartersTemplatesMap()
	cfg.projectSettings, _ = loadProjectSettings(ProjectSettingsFile)
	cfg.prodData, _ = loadEnvFile(DotEnvProdFile)
	cfg.fs = afero.NewOsFs()
}

func loadSveltinSettings() {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(YamlConfig))
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg.settings)
	if err != nil {
		cfg.log.Fatal(err.Error())
	}

}

func loadProjectSettings(filename string) (prjConfig tpltypes.ProjectSettings, err error) {
	currentDir, _ := os.Getwd()
	viper.AddConfigPath(currentDir)
	viper.SetConfigName(filename)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&prjConfig)

	validate := validator.New()
	if err := validate.Struct(&prjConfig); err != nil {
		nErr := sveltinerr.NewNotValidProjectSettingsError(err)
		cfg.log.Fatalf("%s\n", nErr)
	}
	return
}

func loadEnvFile(filename string) (tplData tpltypes.EnvProductionData, err error) {
	currentDir, _ := os.Getwd()
	viper.AddConfigPath(currentDir)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&tplData)
	return
}

/** isValidProject returns an error if sveltin cannot find:
 * - package.json file within the current folder
 * - sveltin.json within the current folder
 */
func isValidProject(checkIfLatestVersion bool) {
	cwd, _ := os.Getwd()
	pathToFile := filepath.Join(cwd, "package.json")
	exists, _ := afero.Exists(cfg.fs, pathToFile)
	if !exists {
		err := sveltinerr.NewNotValidProjectError(pathToFile)
		cfg.log.Fatalf("\n%s", err.Error())
	}

	if checkIfLatestVersion {
		pathToFile = filepath.Join(cwd, ProjectSettingsFile)
		exists, _ = afero.Exists(cfg.fs, pathToFile)
		if !exists {
			err := sveltinerr.NewNotLatestVersionError(pathToFile)
			cfg.log.Fatalf("\n%s", err.Error())
		}
	}

}

//=============================================================================

// GetSveltinCommands returns an array of pointers to the implemented cobra.Command
func GetSveltinCommands() []*cobra.Command {
	return []*cobra.Command{
		initCmd, newCmd, addCmd, generateCmd, installCmd, updateCmd, serverCmd, buildCmd, previewCmd, deployCmd, upgradeCmd,
	}
}

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
	"github.com/sveltinio/sveltin/utils"
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
	// CliVersion is the current sveltin cli version number.
	CliVersion string = "0.10.0"
)

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

// File IDs for a Sveltin project structure.
const (
	ApiIndexFileId      string = "api_index"
	ApiMetadataIndexId  string = "api_metadata_index"
	ApiSlugFileId       string = "api_slug"
	IndexFileId         string = "index"
	IndexEndpointFileId string = "indexendpoint"
	SlugFileId          string = "slug"
	SlugEndpointFileId  string = "slugendpoint"
	SlugLayoutFileId    string = "sluglayout"
	MDsveXFileId        string = "mdsvex"
	DummyImgFileId      string = "dummy"
	SveltinDTSFileId    string = "sveltin_d_ts"
)

// File names for a Sveltin project structure.
const (
	ProjectSettingsFile string = "sveltin.json"
	DefaultsConfigFile  string = "defaults.js.ts"
	DotEnvProdFile      string = ".env.production"
	WebSiteTSFile       string = "website.js.ts"
	MenuTSFile          string = "menu.js.ts"
	PackageJSONFile     string = "package.json"
	MDsveXFile          string = "mdsvex.config.js"
	SvelteConfigFile    string = "svelte.config.js"
	LayoutTSFile        string = "+layout.ts"
	HeadingsJSFile      string = "headings.js"
	StringsTSFile       string = "strings.js.ts"
	SveltinDTSFile      string = "sveltin.d.ts"
	ViteConfigFile      string = "vite.config.ts"
	TSConfigFile        string = "tsconfig.json"
)

// Matchers IDs
const (
	StringMatcher  string = "string_matcher"
	GenericMatcher string = "generic_matcher"
)

//=============================================================================

var (
	// Short description shown in the 'help' output.
	rootCmdShortMsg = "sveltin is the main command to work with SvelteKit powered static websites."
	// Long message shown in the 'help <this-command>' output.
	rootCmdLongMsg = utils.MakeCmdLongMsg(`A powerful CLI for your SvelteKit powered static website!

sveltin is the main command used to boost your productivity
while creating a new production-ready SvelteKit project.

Resources:
  Documentation           -> https://docs.sveltin.io
  A helpful quick-start   -> https://docs.sveltin.io/quick-start`)
)

var (
	npmClientName string
	cfg           appConfig
)

// YamlConfig is used by yaml.Unmarshal to decode the YAML file.
var YamlConfig []byte

//=============================================================================

var rootCmd = &cobra.Command{
	Use:              "sveltin",
	Version:          CliVersion,
	TraverseChildren: true,
	Short:            rootCmdShortMsg,
	Long:             rootCmdLongMsg,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It is called by main.main().
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// Command initialization.
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

//=============================================================================

var preRunHook = func(cmd *cobra.Command, args []string) {
	isValidProject()
	isPre011VersionProject()
}

// Exit if cannot find a package.json file within the current folder.
func isValidProject() {
	cwd, _ := os.Getwd()
	pathToFile := filepath.Join(cwd, PackageJSONFile)
	exists, _ := afero.Exists(cfg.fs, pathToFile)
	if !exists {
		err := sveltinerr.NewNotValidProjectError(pathToFile)
		cfg.log.Fatalf("%s", err.Error())
	}
}

// Exit if cannot find a sveltin.json within the current folder (sveltin < v0.11.0).
func isPre011VersionProject() {
	cwd, _ := os.Getwd()
	pathToFile := filepath.Join(cwd, ProjectSettingsFile)
	exists, _ := afero.Exists(cfg.fs, pathToFile)
	if !exists {
		err := sveltinerr.NewNotLatestVersionError(pathToFile)
		cfg.log.Fatalf("\n%s", err.Error())
	}
}

//=============================================================================

// GetSveltinCommands returns an array of pointers to the implemented cobra.Command.
// Used to generate command documentations.
func GetSveltinCommands() []*cobra.Command {
	return []*cobra.Command{
		initCmd, newCmd, addCmd, generateCmd, installCmd, updateCmd, serverCmd, buildCmd, previewCmd, deployCmd, migrateCmd,
	}
}

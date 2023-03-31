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
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/fsm"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/notifier"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	projectvalidator "github.com/sveltinio/sveltin/internal/validator"
	"github.com/sveltinio/sveltin/utils"
	logger "github.com/sveltinio/yinlog"
)

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

type sveltinCmdConfig struct {
	NpmClient  string `mapstructure:"npmClient" env:"SVELTIN_NPM_CLIENT"`
	CssLib     string `mapstructure:"css" env:"SVELTIN_CSS_LIB"`
	PortNumber string `mapstructure:"port" env:"SVELTIN_SERVER_PORT"`
	InitGit    bool   `mapstructure:"git" env:"SVELTIN_INIT_GIT"`
}

// CliVersion is the current sveltin cli version number.
const CliVersion string = "0.12.0"

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

// File IDs for the Sveltin project structure.
const (
	ApiIndexFileId      string = "api_index"
	ApiMetadataIndexId  string = "api_metadata_index"
	ApiSlugFileId       string = "api_slug"
	IndexPageFileId     string = "index"
	IndexPageLoadFileId string = "index_pageload"
	SlugPageFileId      string = "slug"
	SlugPageLoadFileId  string = "slug_pageload"
	SlugLayoutFileId    string = "slug_layout"
	MDsveXFileId        string = "mdsvex"
	DummyImgFileId      string = "dummy"
	SveltinDTSFileId    string = "sveltin_d_ts"
)

// File names for the Sveltin project structure.
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

// Bind choice to continue with installed version when a newer one is available.
var releaseNotifierHandler bool

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
	cobra.OnInitialize(loadSveltinSettings, initAppConfig, sveltinCmdConfigSetup)
}

//=============================================================================

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

func sveltinCmdConfigSetup() {
	err := bindSveltinCmdFlagsAndEnv()
	if err != nil {
		return
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
	validate := projectvalidator.Init()
	if err := validate.Struct(&prjConfig); err != nil {
		nErr := sveltinerr.NewNotValidProjectSettingsError(err)
		cfg.log.Fatalf("%s\n", nErr)
	}
	return
}

// Used to load .env.production file and make the data available to the deploy cmd.
func loadEnvFile(filename string) (tpltypes.EnvProductionData, error) {
	var tplData tpltypes.EnvProductionData

	currentDir, _ := os.Getwd()
	viper.AddConfigPath(currentDir)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return tpltypes.EnvProductionData{}, err
	}

	err = viper.Unmarshal(&tplData)
	if err != nil {
		return tpltypes.EnvProductionData{}, err
	}

	return tplData, nil
}

//=============================================================================

// Set the dependency between the config sources.
func bindSveltinCmdFlagsAndEnv() error {
	tags, err := utils.GetStructTags(sveltinCmdConfig{})
	if err != nil {
		return err
	}

	for _, t := range tags {
		canonicalTag, err := t.Get("mapstructure")
		if err != nil {
			return err
		}
		envTag, err := t.Get("env")
		if err != nil {
			return err
		}

		err = viper.BindPFlag(canonicalTag.Name, initCmd.Flags().Lookup(canonicalTag.Name))
		if err != nil {
			return err
		}

		err = viper.BindEnv(canonicalTag.Name, envTag.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
 * Configuration options priority:
 * - 1: CLI flags
 * - 2: env variables
 * - 3: sveltin_config.yaml
 */
func setCmdsDefaultConfigs(defaultsConfig *sveltinCmdConfig) error {
	config, err := readSveltinCmdConfig()
	if err != nil {
		return err
	}

	// Set defaultsConfig
	defaultsConfig.NpmClient = config.NpmClient
	defaultsConfig.CssLib = config.CssLib
	defaultsConfig.PortNumber = config.PortNumber
	defaultsConfig.InitGit = config.InitGit

	return nil
}

func readSveltinCmdConfig() (*sveltinCmdConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	viper.AddConfigPath(homeDir)
	viper.SetConfigName("sveltin_config")
	viper.SetConfigType("yaml")

	_ = viper.ReadInConfig()

	cfg := &sveltinCmdConfig{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

//=============================================================================

func isPreRelease() bool {
	return strings.Contains(CliVersion, "-pre")
}

func allExceptInitCmdPreRunHook(cmd *cobra.Command, args []string) {
	isValidProject()
	handleReleaseNotifier(notifier.AllExceptInitCmd)
	isPre011VersionProject()
}

// Exit if cannot find the current directory is not a valid folder structure and some files do not exists.
func isValidProject() {
	foldersToCheck := []string{ConfigFolder, StaticFolder, ThemesFolder}
	filesToCheck := []string{
		PackageJSONFile,
		ProjectSettingsFile,
		MDsveXFile,
		SvelteConfigFile,
		ViteConfigFile,
	}
	cwd, _ := os.Getwd()

	for _, folder := range foldersToCheck {
		pathToFolder := filepath.Join(cwd, folder)
		exists, _ := afero.DirExists(cfg.fs, pathToFolder)
		if !exists {
			err := sveltinerr.NewNotValidProjectError(pathToFolder, "folder")
			cfg.log.Fatalf("%s", err.Error())
		}
	}

	for _, file := range filesToCheck {
		pathToFile := filepath.Join(cwd, file)
		exists, _ := afero.Exists(cfg.fs, pathToFile)
		if !exists {
			err := sveltinerr.NewNotValidProjectError(pathToFile, "file")
			cfg.log.Fatalf("%s", err.Error())
		}
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

func handleReleaseNotifier(cmdName string) {
	if !isPreRelease() {
		cwd, _ := os.Getwd()
		var isCheckUpdates *bool
		var lastCheck string

		switch cmdName {
		case notifier.InitCmd:
			isCheckUpdates = utils.NewTrue()
		case notifier.AllExceptInitCmd:
			isCheckUpdates = cfg.projectSettings.Sveltin.CheckUpdates
			lastCheck = cfg.projectSettings.Sveltin.LastCheck
		}

		// Create a new ReleaseMonitor object.
		newReleaseMonitor := notifier.NewReleaseMonitor(isCheckUpdates, CliVersion, lastCheck)
		// Create observer/subscribe objects.
		cmdObserver := notifier.CommandObserver{
			Id: cmdName,
		}
		sveltinJsonObserver := notifier.SveltinJsonObserver{
			Id:         cmdName,
			Fs:         cfg.fs,
			TargetPath: filepath.Join(cwd, ProjectSettingsFile),
		}

		// Register the Observers to the ReleaseMonitor.
		_, err := newReleaseMonitor.Register(&cmdObserver)
		if err != nil {
			cfg.log.Important(markup.Faint(err.Error()))
		}
		_, err = newReleaseMonitor.Register(&sveltinJsonObserver)
		if err != nil {
			cfg.log.Important(markup.Faint(err.Error()))
		}

		// Start the ReleaseNotifier.
		if releaseNotifierHandler, err = newReleaseMonitor.Notify(); err != nil {
			cfg.log.Fatal(err.Error())
		}

		if !releaseNotifierHandler {
			os.Exit(1)
		}
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

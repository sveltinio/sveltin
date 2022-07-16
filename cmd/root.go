/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/pkg/fsm"
	"github.com/sveltinio/sveltin/pkg/logger"
	"github.com/sveltinio/sveltin/pkg/pathmaker"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/resources"
	"gopkg.in/yaml.v3"
)

//=============================================================================

type appConfig struct {
	log         *logger.Logger
	sveltin     *config.SveltinConfig
	project     config.ProjectConfig
	pathMaker   *pathmaker.SveltinPathMaker
	fsManager   *fsm.SveltinFSManager
	startersMap map[string]config.StarterTemplate
	fs          afero.Fs
}

//=============================================================================

const (
	// CliVersion is the current sveltin cli version number.
	CliVersion string = "0.8.10"
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
	StringMatcher     string = "string_matcher"
	GenericMatcher    string = "generic_matcher"
	ApiIndexFile      string = "api_index"
	ApiMetadataIndex  string = "api_metadata_index"
	ApiSlugFile       string = "api_slug"
	IndexFile         string = "index"
	IndexEndpointFile string = "indexendpoint"
	SlugFile          string = "slug"
	SlugEndpointFile  string = "slugendpoint"
	SettingsFile      string = ".sveltin-settings.yaml"
	DotEnvProdFile    string = ".env.production"
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
while creating a new production-ready project.

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
	cobra.OnInitialize(loadSveltinConfig, initAppConfig)
}

//=============================================================================

func loadSveltinConfig() {
	err := yaml.Unmarshal(YamlConfig, &cfg.sveltin)
	if err != nil {
		jww.FATAL.Fatal(err)
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
	cfg.pathMaker = pathmaker.NewSveltinPathMaker(cfg.sveltin)
	cfg.fsManager = fsm.NewSveltinFSManager(cfg.pathMaker)
	cfg.startersMap = helpers.InitStartersTemplatesMap()
	cfg.project, _ = loadEnvFile(DotEnvProdFile)
	cfg.fs = afero.NewOsFs()
}

func loadEnvFile(filename string) (config config.ProjectConfig, err error) {
	currentDir, _ := os.Getwd()
	viper.AddConfigPath(currentDir)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// isValidProject returns error if cannot find the package.json file within the current folder.
func isValidProject() {
	pwd, _ := os.Getwd()
	pathToPkgJSON := filepath.Join(pwd, "package.json")
	exists, _ := afero.Exists(cfg.fs, pathToPkgJSON)
	if !exists {
		err := sveltinerr.NewNotValidProjectError(pathToPkgJSON)
		jww.FATAL.Fatalf("\x1b[31;1m✘ %s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
}

//=============================================================================

// GetSveltinCommands returns an array of pointers to the implemented cobra.Command
func GetSveltinCommands() []*cobra.Command {
	return []*cobra.Command{
		newCmd, generateCmd, installCmd, updateCmd, serverCmd, buildCmd, previewCmd, deployCmd, newThemeCmd,
	}
}

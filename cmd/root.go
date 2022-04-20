/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
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
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/fsm"
	"github.com/sveltinio/sveltin/sveltinlib/logger"
	"github.com/sveltinio/sveltin/sveltinlib/pathmaker"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"gopkg.in/yaml.v2"
)

//=============================================================================

const (
	// CliVersion is the current sveltin clie version number.
	CliVersion string = "0.7.3"
	// SvelteKitStarter a string rapresenting the project template id.
	SvelteKitStarter string = "starter"
)

// folder and file names for a Sveltin project structure.
const (
	Root          string = "root"
	Backups       string = "backups"
	Config        string = "config"
	Content       string = "content"
	Routes        string = "routes"
	Api           string = "api"
	Lib           string = "lib"
	Static        string = "static"
	Themes        string = "themes"
	Index         string = "index"
	IndexEndpoint string = "indexendpoint"
	Slug          string = "slug"
	SlugEndpoint  string = "slugendpoint"
	SettingsFile  string = ".sveltin-settings.yaml"
	DotEnvProd    string = ".env.production"
)

var (
	// AppFs is the Afero wrapper around the native OS calls.
	AppFs = afero.NewOsFs()
	// YamlConfig is used by yaml.Unmarshal to decode the YAML file.
	YamlConfig []byte
)

var (
	log             = logger.New()
	npmClientName   string
	appTemplatesMap map[string]config.AppTemplate
	pathMaker       pathmaker.SveltinPathMaker
	conf            config.SveltinConfig
	projectConfig   config.ProjectConfig
	fsManager       *fsm.SveltinFSManager
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
	cobra.OnInitialize(setDefaultLoggerOptions, loadSveltinConfig, initSveltin)
}

//=============================================================================

func setDefaultLoggerOptions() {
	log.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    false,
		Icons:     true,
	})
}

func initSveltin() {
	pathMaker = pathmaker.NewSveltinPathMaker(&conf)
	fsManager = fsm.NewSveltinFSManager(&pathMaker)
	appTemplatesMap = helpers.InitAppTemplatesMap()
	projectConfig, _ = loadEnvFile(DotEnvProd)
}

func loadSveltinConfig() {
	err := yaml.Unmarshal(YamlConfig, &conf)
	if err != nil {
		jww.FATAL.Fatal(err)
	}
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
	exists, _ := afero.Exists(AppFs, pathToPkgJSON)
	if !exists {
		err := sveltinerr.NewNotValidProjectError(pathToPkgJSON)
		jww.FATAL.Fatalf("\x1b[31;1m✘ %s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
}

func isSveltinStyles(style string) bool {
	return style == StyleDefault
}

//=============================================================================

// GetSveltinCommands returns an array of pointers to the implemented cobra.Command
func GetSveltinCommands() []*cobra.Command {
	return []*cobra.Command{
		newCmd, generateCmd, installCmd, updateCmd, serverCmd, buildCmd, previewCmd, deployCmd,
	}
}

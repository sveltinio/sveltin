/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"os"

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
	"gopkg.in/yaml.v2"
)

//=============================================================================

const (
	// CLI_VERSION is the current sveltin clie version number.
	CLI_VERSION string = "0.6.0"
	// SVELTEKIT_STARTER a string rapresenting the project template id.
	SVELTEKIT_STARTER string = "starter"
)

// folder and file names for a Sveltin project structure.
const (
	ROOT           string = "root"
	BACKUPS        string = "backups"
	CONFIG         string = "config"
	CONTENT        string = "content"
	ROUTES         string = "routes"
	API            string = "api"
	LIB            string = "lib"
	STATIC         string = "static"
	THEMES         string = "themes"
	INDEX          string = "index"
	INDEX_ENDPOINT string = "indexendpoint"
	SLUG           string = "slug"
	SLUG_ENDPOINT  string = "slugendpoint"
	SETTINGS_FILE  string = ".sveltin-settings.yaml"
	DOTENV_PROD    string = ".env.production"
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
	Version:          CLI_VERSION,
	TraverseChildren: true,
	Short:            "sveltin is the main command to work with SvelteKit powered static websites.",
	Long: resources.GetAsciiArt() + `
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

func setListLoggerOptions() {
	log.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    false,
		Icons:     false,
	})
}

func initSveltin() {
	pathMaker = pathmaker.NewSveltinPathMaker(&conf)
	fsManager = fsm.NewSveltinFSManager(&pathMaker)
	appTemplatesMap = helpers.InitAppTemplatesMap()
	projectConfig, _ = loadEnvFile(DOTENV_PROD)
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

//=============================================================================

// GetSveltinCommands returns an array of pointers to the implemented cobra.Command
func GetSveltinCommands() []*cobra.Command {
	return []*cobra.Command{
		newCmd, generateCmd, installCmd, updateCmd, serverCmd, buildCmd, previewCmd, deployCmd,
	}
}

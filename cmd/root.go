/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
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
	"github.com/sveltinio/sveltin/sveltinlib/pathmaker"
	"github.com/sveltinio/sveltin/utils"
	"gopkg.in/yaml.v2"
)

//=============================================================================

const (
	CLI_VERSION       string = "0.3.1"
	SVELTEKIT_STARTER string = "starter"
)

const (
	DOTENV_DEV  string = ".env.development"
	DOTENV_PROD string = ".env.production"
)

const (
	ROOT          string = "root"
	CONFIG        string = "config"
	CONTENT       string = "content"
	ROUTES        string = "routes"
	API           string = "api"
	LIB           string = "lib"
	STATIC        string = "static"
	THEMES        string = "themes"
	INDEX         string = "index"
	SLUG          string = "slug"
	SETTINGS_FILE string = ".sveltin-settings.yaml"
)

var (
	AppFs      = afero.NewOsFs()
	YamlConfig []byte
)

var (
	textLogger      = utils.NewTextLogger()
	listLogger      = utils.NewListWriter()
	npmClientName   string
	appTemplatesMap map[string]config.AppTemplate
	pathMaker       pathmaker.SveltinPathMaker
	conf            config.SveltinConfig
	siteConfig      config.SiteConfig
	fsManager       *fsm.SveltinFSManager
)

//=============================================================================

var rootCmd = &cobra.Command{
	Use:              "sveltin",
	Version:          CLI_VERSION,
	TraverseChildren: true,
	Short:            "sveltin is the main command to work with SvelteKit powered static websites.",
	Long: resources.GetAsciiArt() + `
sveltin is the main command to work with SvelteKit powered static website.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(loadSveltinConfig, initSveltin)
}

//=============================================================================

func initSveltin() {
	pathMaker = pathmaker.NewSveltinPathMaker(&conf)
	fsManager = fsm.NewSveltinFSManager(&pathMaker)
	appTemplatesMap = helpers.InitAppTemplatesMap()
	siteConfig, _ = loadEnvFile(DOTENV_PROD)
}

func loadSveltinConfig() {
	err := yaml.Unmarshal(YamlConfig, &conf)
	if err != nil {
		jww.FATAL.Fatal(err)
	}
}

func loadEnvFile(filename string) (config config.SiteConfig, err error) {
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

func GetSveltinCommands() []*cobra.Command {
	return []*cobra.Command{
		newCmd, generateCmd, prepareCmd, updateCmd, serverCmd, buildCmd, previewCmd,
	}
}

package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func TestPages(t *testing.T) {
	is := is.New(t)

	var settings SveltinSettings
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	is.NoErr(err)

	err = viper.Unmarshal(&settings)
	is.NoErr(err)

	is.Equal("index.svx", settings.GetContentPageFilename())
	is.Equal("+page.svelte", settings.GetIndexPageFilename())
	is.Equal("+page.svelte", settings.GetSlugPageFilename())
	is.Equal("+page.ts", settings.GetSlugEndpointFilename())
}

func TestPaths(t *testing.T) {
	is := is.New(t)
	var settings SveltinSettings
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	is.NoErr(err)

	err = viper.Unmarshal(&settings)
	is.NoErr(err)

	pwd, _ := os.Getwd()

	tests := []struct {
		path string
		want string
	}{
		{path: settings.GetBuildPath(), want: filepath.Join(pwd, "build")},
		{path: settings.GetConfigPath(), want: filepath.Join("config")},
		{path: settings.GetContentPath(), want: filepath.Join("content")},
		{path: settings.GetStaticPath(), want: filepath.Join("static")},
		{path: settings.GetSrcPath(), want: filepath.Join("src")},
		{path: settings.GetRoutesPath(), want: filepath.Join("src", "routes")},
		{path: settings.GetLibPath(), want: filepath.Join("src", "lib")},
		{path: settings.GetParamsPath(), want: filepath.Join("src", "params")},
		{path: settings.GetAPIPath(), want: filepath.Join("src", "routes", "api")},
		{path: settings.GetThemesPath(), want: filepath.Join(pwd, "themes")},
	}

	for _, tc := range tests {
		is := is.New(t)
		is.Equal(tc.want, tc.path)
	}

}

func TestAPIs(t *testing.T) {
	is := is.New(t)

	var settings SveltinSettings
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	if err != nil {
		return
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		log.Fatal(err)
	}

	/*err = yaml.Unmarshal(yamlFile, &settings)
	is.NoErr(err)
	*/

	is.Equal("v1", settings.GetAPIVersion())
	is.Equal("+server.ts", settings.GetAPIFilename())
}

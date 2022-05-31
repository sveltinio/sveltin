package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

func TestPages(t *testing.T) {
	is := is.New(t)

	var conf SveltinConfig
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	is.Equal("index.svx", conf.GetContentPageFilename())
	is.Equal("index.svelte", conf.GetIndexPageFilename())
	is.Equal("[slug].svelte", conf.GetSlugPageFilename())
}

func TestPaths(t *testing.T) {
	is := is.New(t)
	var conf SveltinConfig
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	pwd, _ := os.Getwd()

	tests := []struct {
		path string
		want string
	}{
		{path: conf.GetBuildPath(), want: filepath.Join(pwd, "build")},
		{path: conf.GetConfigPath(), want: filepath.Join("config")},
		{path: conf.GetContentPath(), want: filepath.Join("content")},
		{path: conf.GetStaticPath(), want: filepath.Join("static")},
		{path: conf.GetSrcPath(), want: filepath.Join("src")},
		{path: conf.GetRoutesPath(), want: filepath.Join("src", "routes")},
		{path: conf.GetLibPath(), want: filepath.Join("src", "lib")},
		{path: conf.GetAPIPath(), want: filepath.Join("src", "routes", "api")},
		{path: conf.GetThemesPath(), want: filepath.Join(pwd, "themes")},
	}

	for _, tc := range tests {
		is := is.New(t)
		is.Equal(tc.want, tc.path)
	}

}

func TestAPIs(t *testing.T) {
	is := is.New(t)

	var conf SveltinConfig
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	is.Equal("v1", conf.GetAPIVersion())
	is.Equal("published.json.ts", conf.GetAPIFilename())
	is.Equal("published.json", conf.GetPublicAPIFilename())
	is.Equal("category.json.ts", conf.GetMetadataAPIFilename("category"))
}

// Package helpers ...
package helpers

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/internal/tpltypes"
)

func TestPublicPageFilename(t *testing.T) {
	tests := []struct {
		pageData *config.TemplateData
		want     string
	}{
		{
			pageData: &config.TemplateData{
				Page: &tpltypes.PageData{
					Name:     "index",
					Language: "svelte"},
			},
			want: "+page.svelte",
		},
		{
			pageData: &config.TemplateData{
				Page: &tpltypes.PageData{
					Name:     "about",
					Language: "markdown",
				},
			},
			want: "+page.svx",
		},
	}

	for _, tc := range tests {
		is := is.New(t)
		is.Equal(PublicPageFilename(tc.pageData.Page.Language), tc.want)
	}
}

func TestGetResourceRouteFilename(t *testing.T) {
	conf := loadConfigFile(filepath.Join("..", "resources", "sveltin.yaml"))

	tests := []struct {
		name string
		want string
	}{
		{
			name: "index",
			want: "+page.svelte",
		},
		{
			name: "indexendpoint",
			want: "+page.ts",
		},
		{
			name: "slug",
			want: "+page.svelte",
		},
		{
			name: "slugendpoint",
			want: "+page.ts",
		},
	}
	for _, tc := range tests {
		is := is.New(t)
		is.Equal(GetRouteFilename(tc.name, &conf), tc.want)
	}
}

func loadConfigFile(filepath string) config.SveltinSettings {
	var settings config.SveltinSettings
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath)
	if err != nil {
		os.Exit(0)
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		log.Fatal(err)
	}
	return settings
}

// Package helpers ...
package helpers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"gopkg.in/yaml.v3"
)

func TestPublicPageFilename(t *testing.T) {
	tests := []struct {
		pageData *config.TemplateData
		want     string
	}{
		{
			pageData: &config.TemplateData{
				Page: &tpltypes.PageData{
					Name: "index",
					Type: "svelte"},
			},
			want: "+page.svelte",
		},
		{
			pageData: &config.TemplateData{
				Page: &tpltypes.PageData{
					Name: "about",
					Type: "markdown",
				},
			},
			want: "+page.svx",
		},
	}

	for _, tc := range tests {
		is := is.New(t)
		is.Equal(PublicPageFilename(tc.pageData.Page.Type), tc.want)
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
			want: "+page.server.ts",
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
		is.Equal(GetResourceRouteFilename(tc.name, &conf), tc.want)
	}
}

func loadConfigFile(filepath string) config.SveltinConfig {
	var conf config.SveltinConfig
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath)
	if err != nil {
		os.Exit(0)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		os.Exit(0)
	}
	return conf
}

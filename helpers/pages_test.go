// Package helpers ...
package helpers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"gopkg.in/yaml.v2"
)

func TestGetPublicPages(t *testing.T) {
	is := is.New(t)
	path := "src/routes"

	tests := []struct {
		filename string
	}{
		{filename: "index.svelte"},
		{filename: "about.svx"},
		{filename: "contact.mdx"},
	}

	memFS := afero.NewMemMapFs()
	for _, tc := range tests {
		is.NoErr(common.TouchFile(memFS, filepath.Join(path, tc.filename)))
	}

	pages := GetAllPublicPages(memFS, path)
	is.Equal(2, len(pages))
}

func TestPublicPageFilename(t *testing.T) {
	tests := []struct {
		pageData *config.TemplateData
		want     string
	}{
		{
			pageData: &config.TemplateData{
				Name: "index",
				Type: "svelte",
			},
			want: "index.svelte",
		},
		{
			pageData: &config.TemplateData{
				Name: "about",
				Type: "markdown",
			},
			want: "about.svx",
		},
	}

	for _, tc := range tests {
		is := is.New(t)
		is.Equal(PublicPageFilename(tc.pageData.Name, tc.pageData.Type), tc.want)
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
			want: "index.svelte",
		},
		{
			name: "indexendpoint",
			want: "index.ts",
		},
		{
			name: "slug",
			want: "[slug].svelte",
		},
		{
			name: "slugendpoint",
			want: "[slug].ts",
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

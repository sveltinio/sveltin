package helpers

import (
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
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

package helpers

import (
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
)

func TestValidFileForContent(t *testing.T) {
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

	files, err := afero.ReadDir(memFS, path)
	is.NoErr(err)
	for _, f := range files {
		is.True(IsValidFileForContent(f))
	}

}

func TestNotValidFileForContent(t *testing.T) {
	is := is.New(t)
	path := "src/routes"

	notValidFileExtTests := []struct {
		filename string
	}{
		{filename: "index.html"},
		{filename: "about.json"},
	}

	memFS := afero.NewMemMapFs()
	for _, tc := range notValidFileExtTests {
		is.NoErr(common.TouchFile(memFS, filepath.Join(path, tc.filename)))
	}

	files, err := afero.ReadDir(memFS, path)
	is.NoErr(err)
	for _, f := range files {
		is.Equal(false, IsValidFileForContent(f))
	}

}

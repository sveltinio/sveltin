package common

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/resources"
)

func TestMkDir(t *testing.T) {
	tests := []struct {
		folder string
	}{
		{folder: "config"},
		{folder: "themes/basic"},
		{folder: "src/routes/api"},
	}

	memFS := afero.NewMemMapFs()
	for _, tc := range tests {
		is := is.New(t)

		err := MkDir(memFS, tc.folder)
		is.NoErr(err)
		exists, err := afero.DirExists(memFS, tc.folder)
		is.NoErr(err)
		is.True(exists)

	}
}

func TestDirExists(t *testing.T) {
	tests := []struct {
		folder string
	}{
		{folder: "config"},
		{folder: "themes/basic"},
		{folder: "src/routes/api"},
	}

	memFS := afero.NewMemMapFs()
	for _, tc := range tests {
		is := is.New(t)

		err := MkDir(memFS, tc.folder)
		is.NoErr(err)

		is.True(DirExists(memFS, tc.folder))
	}

	dirNotFoundTests := []struct {
		folder string
	}{
		{folder: "configs"},
	}

	for _, tc := range dirNotFoundTests {
		is := is.New(t)
		is.Equal(false, DirExists(memFS, tc.folder))
	}
}

func TestFileExists(t *testing.T) {
	tests := []struct {
		pathToFile string
	}{
		{pathToFile: "config/website.js"},
		{pathToFile: "themes/basic/theme.config.js"},
	}

	memFS := afero.NewMemMapFs()
	for _, tc := range tests {
		is := is.New(t)

		err := TouchFile(memFS, tc.pathToFile)
		is.NoErr(err)
		exists, err := FileExists(memFS, tc.pathToFile)
		is.NoErr(err)
		is.True(exists)

	}

	fileNotFoundTests := []struct {
		pathToFile string
	}{
		{pathToFile: "config/website.json"},
	}

	for _, tc := range fileNotFoundTests {
		is := is.New(t)

		exists, err := FileExists(memFS, tc.pathToFile)
		re := err.(*sveltinerr.SveltinError)
		is.Equal(10, int(re.Code))
		is.Equal("FileNotFoundError", re.Name)
		is.Equal(false, exists)
	}

	dirInsteadOfFileTests := []struct {
		pathToFile string
	}{
		{pathToFile: "config"},
	}

	for _, tc := range dirInsteadOfFileTests {
		is := is.New(t)

		exists, err := FileExists(memFS, tc.pathToFile)
		re := err.(*sveltinerr.SveltinError)
		is.Equal(11, int(re.Code))
		is.Equal("DirInsteadOfFileError", re.Name)
		is.Equal(false, exists)
	}
}

func TestCopyFileFromEmbeddedFS(t *testing.T) {
	tests := []struct {
		pathToFile string
		saveTo     string
	}{
		{pathToFile: resources.SveltinProjectFS["website"], saveTo: "website.js.gotxt"},
		{pathToFile: resources.SveltinProjectFS["menu"], saveTo: "menu.js.gotxt"},
	}

	memFS := afero.NewMemMapFs()
	for _, tc := range tests {
		is := is.New(t)

		err := CopyFileFromEmbeddedFS(&resources.SveltinFS, memFS, tc.pathToFile, tc.saveTo)
		is.NoErr(err)

		exists, err := afero.Exists(memFS, tc.saveTo)
		is.NoErr(err)
		is.Equal(nil, err)
		is.True(exists)
	}

	fileNotFoundTests := []struct {
		pathToFile string
		saveTo     string
	}{
		{pathToFile: filepath.Join("..", "resources", "internal", "templates", "site", "website"), saveTo: "website.js"},
	}

	for _, tc := range fileNotFoundTests {
		is := is.New(t)

		err := CopyFileFromEmbeddedFS(&resources.SveltinFS, memFS, tc.pathToFile, tc.saveTo)
		re := err.(*sveltinerr.SveltinError)
		is.Equal(10, int(re.Code))
		placeholderText := `file not found! Please, check the file path:

%s`
		msg := fmt.Sprintf(placeholderText, tc.pathToFile)
		is.Equal(msg, re.Message)

	}

}

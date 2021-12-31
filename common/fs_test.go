package common

import (
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
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
		is.Equal(11, re.Code)
		is.Equal("SVELTIN FileNotFoundError", re.Message)
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
		is.Equal(12, re.Code)
		is.Equal("SVELTIN DirInsteadOfFileError", re.Message)
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
		is.Equal(11, re.Code)
		is.Equal("SVELTIN FileNotFoundError: please, check the file path", re.Error())

	}

}

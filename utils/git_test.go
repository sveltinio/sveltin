package utils

import (
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
)

func TestGit(t *testing.T) {
	is := is.New(t)
	osFs := afero.NewOsFs()

	appTemplatesMap := make(map[string]config.AppTemplate)
	appTemplatesMap["hello-world"] = config.AppTemplate{
		Name: "golang-example",
		URL:  "https://github.com/golang/example.git",
	}

	helloWorld := appTemplatesMap["hello-world"]
	GitClone(&helloWorld, helloWorld.Name)

	exists, err := afero.DirExists(osFs, helloWorld.Name)
	is.NoErr(err)
	is.True(exists)

	_, err = afero.ReadDir(osFs, helloWorld.Name)
	is.NoErr(err)

	osFs.RemoveAll(helloWorld.Name)
}

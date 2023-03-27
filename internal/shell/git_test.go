package shell

import (
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
)

func TestGitClone(t *testing.T) {
	is := is.New(t)
	osFs := afero.NewOsFs()

	appTemplatesMap := make(map[string]config.StarterTemplate)
	appTemplatesMap["hello-world"] = config.StarterTemplate{
		Name: "golang-example",
		URL:  "https://github.com/golang/example.git",
	}

	helloWorld := appTemplatesMap["hello-world"]
	gitClient := NewGitClient()
	err := gitClient.RunClone(helloWorld.URL, helloWorld.Name, true)
	is.NoErr(err)

	exists, err := afero.DirExists(osFs, helloWorld.Name)
	is.NoErr(err)
	is.True(exists)

	_, err = afero.ReadDir(osFs, helloWorld.Name)
	is.NoErr(err)

	is.NoErr(osFs.RemoveAll(helloWorld.Name))
}

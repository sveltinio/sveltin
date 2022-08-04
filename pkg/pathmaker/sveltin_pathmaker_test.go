// Package pathmaker
package pathmaker

import (
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
	"gopkg.in/yaml.v3"
)

func TestPages(t *testing.T) {
	is := is.New(t)

	var conf config.SveltinConfig
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	pathMaker := NewSveltinPathMaker(&conf)

	artifact := "posts"
	is.Equal(filepath.Join("index.svx"), pathMaker.GetResourceContentFilename())
	is.Equal("loadPosts.ts", pathMaker.GetResourceLibFilename(artifact))
}

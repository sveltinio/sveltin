// Package pathmaker
package pathmaker

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/sveltinio/sveltin/config"
)

func TestPages(t *testing.T) {
	is := is.New(t)

	var settings config.SveltinSettings
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	is.NoErr(err)

	err = viper.Unmarshal(&settings)
	is.NoErr(err)

	pathMaker := NewSveltinPathMaker(&settings)

	artifact := "posts"
	is.Equal(filepath.Join("index.svx"), pathMaker.GetResourceContentFilename())
	is.Equal("loadPosts.ts", pathMaker.GetResourceLibFilename(artifact))
}

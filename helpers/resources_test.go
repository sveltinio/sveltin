// Package helpers ...
package helpers

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
)

func TestGetResources(t *testing.T) {
	var settings config.SveltinSettings
	is := is.New(t)
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	is.NoErr(err)

	err = viper.Unmarshal(&settings)
	is.NoErr(err)

	is.Equal(filepath.Join("content"), settings.GetContentPath())

	memFs := afero.NewMemMapFs()
	is.NoErr(common.MkDir(memFs, settings.GetContentPath()))

	// Create dummy folders for resources
	dummyResources := []string{"posts", "projects", "testimonials"}
	for _, r := range dummyResources {
		is.NoErr(common.MkDir(memFs, filepath.Join(settings.GetContentPath(), r)))
	}

	resources := GetAllResources(memFs, settings.GetContentPath())
	is.Equal(3, len(resources))

	tests := []struct {
		resName string
	}{
		{resName: "posts"},
		{resName: "projects"},
		{resName: "testimonials"},
	}

	for _, tc := range tests {
		is.True(lo.Contains(resources, tc.resName))
	}
}

func TestGetResourceContentMap(t *testing.T) {

	var settings config.SveltinSettings
	is := is.New(t)
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	is.NoErr(err)

	err = viper.Unmarshal(&settings)
	is.NoErr(err)

	is.Equal(filepath.Join("content"), settings.GetContentPath())

	memFs := afero.NewMemMapFs()
	is.NoErr(common.MkDir(memFs, settings.GetContentPath()))

	// Create dummy folders for resources
	dummyResources := []string{"posts", "projects", "testimonials"}
	for _, r := range dummyResources {
		is.NoErr(common.MkDir(memFs, filepath.Join(settings.GetContentPath(), r)))
	}

	resources := GetAllResources(memFs, settings.GetContentPath())
	is.Equal(3, len(resources))

	// Create dummy folders for content
	dummyContents := []string{"first", "second", "third"}
	for _, r := range dummyResources {
		for _, c := range dummyContents {
			is.NoErr(common.MkDir(memFs, filepath.Join(settings.GetContentPath(), r, c)))
		}
	}
	retrievedContents := GetResourceContentMap(memFs, resources, settings.GetContentPath())

	for res, content := range retrievedContents {
		is.True(lo.Contains(resources, res))
		for _, c := range content {
			is.True(lo.Contains(dummyContents, c))
		}
	}

}

func TestGetResourceMetadataMap(t *testing.T) {
	pwd, _ := os.Getwd()
	var settings config.SveltinSettings
	is := is.New(t)
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))
	is.NoErr(err)

	err = viper.Unmarshal(&settings)
	is.NoErr(err)

	is.Equal(filepath.Join("content"), settings.GetContentPath())

	is.Equal(filepath.Join("src", "routes"), settings.GetRoutesPath())

	memFs := afero.NewMemMapFs()
	is.NoErr(common.MkDir(memFs, settings.GetRoutesPath()))

	// Create dummy folders for resources
	dummyResources := []string{"posts", "projects", "testimonials"}
	for _, r := range dummyResources {
		is.NoErr(common.MkDir(memFs, filepath.Join(settings.GetContentPath(), r)))
		is.NoErr(common.MkDir(memFs, filepath.Join(settings.GetRoutesPath(), r)))
	}

	resources := GetAllResources(memFs, settings.GetContentPath())
	is.Equal(3, len(resources))

	// Create dummy folders for resources
	dummyMetadata := []string{"author", "category"}
	for _, r := range resources {
		for _, m := range dummyMetadata {
			is.NoErr(common.MkDir(memFs, filepath.Join(pwd, settings.GetRoutesPath(), r, m)))
		}
	}

	retrievedMetadata := GetResourceMetadataMap(memFs, resources, settings.GetRoutesPath())

	for res, metadata := range retrievedMetadata {
		is.True(lo.Contains(dummyResources, res))
		for _, m := range metadata {
			is.True(lo.Contains(dummyMetadata, m))
		}
	}

}

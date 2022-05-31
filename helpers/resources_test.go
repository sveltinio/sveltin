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
	"gopkg.in/yaml.v3"
)

func TestGetResources(t *testing.T) {
	var conf config.SveltinConfig
	is := is.New(t)
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	is.Equal(filepath.Join("content"), conf.GetContentPath())

	memFs := afero.NewMemMapFs()
	is.NoErr(common.MkDir(memFs, conf.GetContentPath()))

	// Create dummy folders for resources
	dummyResources := []string{"posts", "projects", "testimonials"}
	for _, r := range dummyResources {
		is.NoErr(common.MkDir(memFs, filepath.Join(conf.GetContentPath(), r)))
	}

	resources := GetAllResources(memFs, conf.GetContentPath())
	is.Equal(3, len(resources))

	tests := []struct {
		resName string
	}{
		{resName: "posts"},
		{resName: "projects"},
		{resName: "testimonials"},
	}

	for _, tc := range tests {
		is.True(common.Contains(resources, tc.resName))
	}
}

func TestGetResourceContentMap(t *testing.T) {

	var conf config.SveltinConfig
	is := is.New(t)
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	is.Equal(filepath.Join("content"), conf.GetContentPath())

	memFs := afero.NewMemMapFs()
	is.NoErr(common.MkDir(memFs, conf.GetContentPath()))

	// Create dummy folders for resources
	dummyResources := []string{"posts", "projects", "testimonials"}
	for _, r := range dummyResources {
		is.NoErr(common.MkDir(memFs, filepath.Join(conf.GetContentPath(), r)))
	}

	resources := GetAllResources(memFs, conf.GetContentPath())
	is.Equal(3, len(resources))

	// Create dummy folders for content
	dummyContents := []string{"first", "second", "third"}
	for _, r := range dummyResources {
		for _, c := range dummyContents {
			is.NoErr(common.MkDir(memFs, filepath.Join(conf.GetContentPath(), r, c)))
		}
	}
	retrievedContents := GetResourceContentMap(memFs, resources, conf.GetContentPath())

	for res, content := range retrievedContents {
		is.True(common.Contains(resources, res))
		for _, c := range content {
			is.True(common.Contains(dummyContents, c))
		}
	}

}

func TestGetResourceMetadataMap(t *testing.T) {
	pwd, _ := os.Getwd()
	var conf config.SveltinConfig
	is := is.New(t)
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	is.Equal(filepath.Join("content"), conf.GetContentPath())

	is.Equal(filepath.Join("src", "routes"), conf.GetRoutesPath())

	memFs := afero.NewMemMapFs()
	is.NoErr(common.MkDir(memFs, conf.GetRoutesPath()))

	// Create dummy folders for resources
	dummyResources := []string{"posts", "projects", "testimonials"}
	for _, r := range dummyResources {
		is.NoErr(common.MkDir(memFs, filepath.Join(conf.GetContentPath(), r)))
		is.NoErr(common.MkDir(memFs, filepath.Join(conf.GetRoutesPath(), r)))
	}

	resources := GetAllResources(memFs, conf.GetContentPath())
	is.Equal(3, len(resources))

	// Create dummy folders for resources
	dummyMetadata := []string{"author", "category"}
	for _, r := range resources {
		for _, m := range dummyMetadata {
			is.NoErr(common.MkDir(memFs, filepath.Join(pwd, conf.GetRoutesPath(), r, m)))
		}
	}

	retrievedMetadata := GetResourceMetadataMap(memFs, resources, conf.GetRoutesPath())

	for res, metadata := range retrievedMetadata {
		is.True(common.Contains(dummyResources, res))
		for _, m := range metadata {
			is.True(common.Contains(dummyMetadata, m))
		}
	}

}

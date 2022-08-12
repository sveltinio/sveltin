/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers ...
package helpers

import (
	"path/filepath"

	"github.com/spf13/afero"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

// ResourceExists return nil if a Resource identified by name exists.
func ResourceExists(fs afero.Fs, name string, c *config.SveltinConfig) error {
	availableResources := GetAllResources(fs, c.GetContentPath())
	if !common.Contains(availableResources, name) {
		return sveltinerr.NewResourceNotFoundError()
	}
	return nil
}

// GetAllResources returns a slice of resource names as string.
func GetAllResources(fs afero.Fs, path string) []string {
	resources := []string{}
	if common.DirExists(fs, path) {
		files, err := afero.ReadDir(fs, path)
		if err != nil {
			jww.FATAL.Fatalf("Something went wrong visiting dir %s. Are you sure it exists?", path)
		}
		for _, f := range files {
			if f.IsDir() {
				resources = append(resources, f.Name())
			}
		}
	}
	return resources
}

// GetAllResourcesWithContentName traverses the resource contents and returns a slice of ResourceItem.
func GetAllResourcesWithContentName(fs afero.Fs, path string, children bool) []*config.ResourceItem {
	var result []*config.ResourceItem
	exists, _ := afero.DirExists(fs, path)
	if exists {
		files, err := afero.ReadDir(fs, path)
		if err != nil {
			jww.FATAL.Fatalf("Something went wrong visiting dir %s. Are you sure it exists?", path)
		}
		for _, f := range files {
			if f.IsDir() {
				item := config.NewResourceItem(f.Name())
				if children {
					subFolders, err := afero.ReadDir(fs, filepath.Join(path, f.Name()))
					if err != nil {
						jww.FATAL.Fatalf("Something went wrong visiting subfolder %s. Are you sure it exists?", f.Name())
					}
					for _, s := range subFolders {
						if s.IsDir() {
							item.AddChild(s.Name())
						}
					}
				}
				result = append(result, item)

			}
		}
	}

	return result
}

// GetResourceContentMap returns a map of resources and relative contents.
func GetResourceContentMap(fs afero.Fs, resources []string, path string) map[string][]string {
	content := make(map[string][]string)
	if common.DirExists(fs, path) {
		for _, resource := range resources {
			resourcePath := filepath.Join(path, resource)
			r, _ := afero.ReadDir(fs, resourcePath)
			for _, entry := range r {
				if entry.IsDir() {
					content[resource] = append(content[resource], entry.Name())
				}
			}
		}
	}

	return content
}

// GetResourceMetadataMap returns a map of metadata and relative name.
func GetResourceMetadataMap(fs afero.Fs, resources []string, path string) map[string][]string {
	metadata := make(map[string][]string)
	if common.DirExists(fs, path) {
		for _, resource := range resources {
			resourcePath := filepath.Join(path, resource)
			r, _ := afero.ReadDir(fs, resourcePath)
			for _, entry := range r {
				if entry.IsDir() {
					metadata[resource] = append(metadata[resource], entry.Name())
				}
			}
		}
	}
	return metadata
}

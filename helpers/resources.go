/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers ...
package helpers

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/afero"
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
			log.Fatalf("Something went wrong visiting the folder %s. Are you sure it exists?", path)
		}
		for _, f := range files {
			if f.IsDir() {
				resources = append(resources, f.Name())
			}
		}
	}
	return resources
}

// GetAllRoutes return a slice of all routes names as string.
func GetAllRoutes(fs afero.Fs, path string) []string {
	entries := []string{}
	if common.DirExists(fs, path) {
		walkFunc := func(filepath string, info os.FileInfo, err error) error {
			if info.IsDir() {
				replacer := strings.NewReplacer(path, "", "/[slug]", "")
				res := replacer.Replace(filepath)
				res = strings.TrimSpace(res)

				if !strings.HasPrefix(res, "/api") {
					// Match (group) name
					re := regexp.MustCompile(`\/\((.*?)\)`)
					if re.MatchString(res) {
						submatchall := re.FindAllString(res, -1)

						for _, element := range submatchall {
							element = strings.ReplaceAll(res, element, "")
							element = strings.Replace(element, "/", "", 1)

							if !common.Contains(entries, element) {
								entries = append(entries, element)
							}
						}
					} else {
						if !common.Contains(entries, res) {
							entries = append(entries, res)
						}
					}
				}
				return nil
			}
			return nil
		}

		err := afero.Walk(fs, path, walkFunc)
		if err != nil {
			log.Fatalf("Something went wrong visiting the folder %s. Are you sure it exists?", path)
		}
	}

	routes := []string{}
	for _, file := range entries {
		routes = append(routes, strings.Replace(file, "/", "", 1))
	}
	return common.Unique(routes)
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
				if entry.IsDir() && excludeIfNotValidEntry(entry.Name()) {
					metadata[resource] = append(metadata[resource], entry.Name())
				}
			}
		}
	}
	return metadata
}

func excludeIfNotValidEntry(s string) bool {
	return !(strings.HasPrefix(s, "[") || strings.Contains(s, "["))
}

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
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
)

// GetAllPublicPages return a slice of all available public page names as string.
func GetAllPublicPages(fs afero.Fs, path string) []string {
	files, err := afero.ReadDir(fs, path)
	pages := []string{}

	if err != nil {
		log.Fatalf("Something went wrong visiting the folder %s. Are you sure it exists?", path)
	}

	for _, f := range files {
		pageName := ""
		if IsValidFileForContent(f) {
			if f.Name() != "index.svelte" {
				pageName = strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				pages = append(pages, `"`+pageName+`"`)
			}
		}
	}

	return pages
}

// GetResourceRouteFilename returns a string representing the index and slug routes for a resource.
func GetResourceRouteFilename(txt string, c *config.SveltinConfig) string {
	switch txt {
	case "index":
		return c.GetIndexPageFilename()
	case "indexendpoint":
		return c.GetIndexEndpointFilename()
	case "slug":
		return c.GetSlugPageFilename()
	case "slugendpoint":
		return c.GetSlugEndpointFilename()
	default:
		return ""
	}
}

// PublicPageFilename returns the filename string for a public page based on the page type (svelte or markdown).
func PublicPageFilename(pageType string) string {
	switch pageType {
	case "svelte":
		return `+page.svelte`
	case "markdown":
		return `+page.svx`
	default:
		return ""
	}
}

// NewNoPageItems return a NoPageItems.
func NewNoPageItems(resources []string, content map[string][]string, metadata map[string][]string, pages []string) *config.NoPageItems {
	r := new(config.NoPageItems)
	r.Resources = resources
	r.Content = content
	r.Metadata = metadata
	r.Pages = pages
	return r
}

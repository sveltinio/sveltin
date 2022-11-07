/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package helpers

import (
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/internal/tpltypes"
)

// GetResourceRouteFilename returns a string representing the index and slug routes for a resource.
func GetResourceRouteFilename(txt string, s *config.SveltinSettings) string {
	switch txt {
	case "index":
		return s.GetIndexPageFilename()
	case "indexendpoint":
		return s.GetIndexEndpointFilename()
	case "slug":
		return s.GetSlugPageFilename()
	case "slugendpoint":
		return s.GetSlugEndpointFilename()
	case "sluglayout":
		return s.GetSlugLayoutFilename()
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
func NewNoPageItems(resources []string, content map[string][]string) *tpltypes.NoPageItems {
	r := new(tpltypes.NoPageItems)
	r.Resources = resources
	r.Content = content
	return r
}

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

// GetRouteFilename returns a string representing the index and slug routes filename.
func GetRouteFilename(txt string, s *config.SveltinSettings) string {
	switch txt {
	case "index":
		return s.GetIndexPageFilename()
	case "index_pageload":
		return s.GetIndexEndpointFilename()
	case "slug":
		return s.GetSlugPageFilename()
	case "slug_pageload":
		return s.GetSlugEndpointFilename()
	case "slug_layout":
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

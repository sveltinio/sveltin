/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package tpltypes

// NoPageData is the struct representing a no-public page (sitemap and rss) for a sveltin project.
type NoPageData struct {
	Config *ProjectData
	Items  *NoPageItems
}

// NoPageItems is the struct representing an item
// of no-public page (sitemap and rss) for a sveltin project.
type NoPageItems struct {
	Resources []string
	Content   map[string][]string
}

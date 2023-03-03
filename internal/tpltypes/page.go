/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package tpltypes

// PageData is the struct representing the user selection for the page.
type PageData struct {
	Name     string
	Language string
}

// Supported languages for pages.
const (
	Svelte   string = "svelte"
	Markdown string = "markdown"
)

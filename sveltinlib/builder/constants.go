/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder ...
package builder

// constants representing different file names.
const (
	Defaults         string = "defaults"
	Externals        string = "externals"
	Website          string = "website"
	Menu             string = "menu"
	InitMenu         string = "init_menu"
	DotEnv           string = "dotenv"
	Readme           string = "readme"
	License          string = "license"
	ThemeConfig      string = "theme_config"
	IndexPage        string = "index"
	IndexNoThemePage string = "index_notheme"
)

const (
	// Blank represents the fontmatter-only template id used when generating the content file.
	Blank string = "blank"
	// Sample represents the sample-content template id used when generating the content file.
	Sample string = "sample"

	// Svelte set svelte as the language used to scaffold a new page
	Svelte string = "svelte"
	// Markdown set markdown as the language used to scaffold a new page
	Markdown string = "markdown"

	// API is a string for the 'api' folder.
	API string = "api"
	// APISingle is a string representing the api template file
	// to be used when creating a metadata of type 'single'.
	APISingle string = "api_single"
	// APIList is a string representing the api template file
	// to be used when creating a metadata of type 'list'.
	APIList string = "api_list"

	// Index is a string for the 'index' file.
	Index string = "index"
	// IndexEndpoint is a string for the 'index.ts' file.
	IndexEndpoint string = "indexendpoint"
	// Slug is a string for the 'slug' file.
	Slug string = "slug"
	// SlugEndpoint is a string for the 'slug' file.
	SlugEndpoint string = "slugendpoint"

	// Lib is a string for the 'lib' folder.
	Lib string = "lib"
	// LibSingle is a string representing the template id used
	// for the lib file when metadata's type is single.
	LibSingle string = "lib_single"
	// LibList is a string representing the template id used
	// for the lib file when metadata's type is list.
	LibList string = "lib_list"
)

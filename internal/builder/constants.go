/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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

	//=============================================================================

	// Svelte set svelte as the language used to scaffold a new page
	Svelte string = "svelte"
	// Markdown set markdown as the language used to scaffold a new page
	Markdown string = "markdown"

	//=============================================================================

	// StringMatcher is the string for the string parameters matcher
	StringMatcher string = "string_matcher"
	// GenericMatcher is the string for the generic parameters matcher
	GenericMatcher string = "generic_matcher"

	//=============================================================================

	// ApiFolder is the string for the 'api' folder.
	ApiFolder string = "api"
	// ApiIndexFile is the string for the index api file.
	ApiIndexFile string = "api_index"
	// ApiSlugFile is the string for the slug api file.
	ApiSlugFile string = "api_slug"
	// ApiMetadataIndex is the string for the api template file
	// to get all resources grouped by metadata.
	ApiMetadataIndex string = "api_metadata_index"
	// ApiMetadataSingle is the string for the api template file
	// to be used when creating a metadata of type 'single' and
	// to get all resources filtered by metadata name.
	ApiMetadataSingle string = "api_metadata_single"
	// ApiMetadataList is the string for the api template file
	// to be used when creating a metadata of type 'list' and
	// to get all resources filtered by metadata name.
	ApiMetadataList string = "api_metadata_list"

	//=============================================================================

	// Index is the string for the 'index' file.
	Index string = "index"
	// IndexEndpoint is the string for the 'index.ts' file.
	IndexEndpoint string = "indexendpoint"
	// Slug is the string for the 'slug' file.
	Slug string = "slug"
	// SlugEndpoint is the string for the 'slug' file.
	SlugEndpoint string = "slugendpoint"

	//=============================================================================

	// Lib is the string for the 'lib' folder.
	Lib string = "lib"
	// LibSingle is the string representing the template id used
	// for the lib file when metadata's type is single.
	LibSingle string = "lib_single"
	// LibList is the string representing the template id used
	// for the lib file when metadata's type is list.
	LibList string = "lib_list"
)

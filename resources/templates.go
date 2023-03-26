/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package resources provides access to files embedded in the running Sveltin program.
package resources

import "embed"

const sveltinASCIIArt = `
                _ _   _
               | | | (_)
  _____   _____| | |_ _ _ __
 / __\ \ / / _ \ | __| | '_ \
 \__ \\ V /  __/ | |_| | | | |
 |___/ \_/ \___|_|\__|_|_| |_|

`

// GetASCIIArt returns the ascii art string.
func GetASCIIArt() string {
	return sveltinASCIIArt
}

// SveltinTemplatesFS is the name for the embedded FS used by Sveltin.
//
//go:embed templates/*
var SveltinTemplatesFS embed.FS

// EmbeddedFSEntry type is a map[string]string used to maps embedded files.
type EmbeddedFSEntry map[string]string

// ProjectFilesMap is the map for the project template files.
var ProjectFilesMap = EmbeddedFSEntry{
	"defaults":         "templates/site/defaults.js.ts.gotxt",
	"externals":        "templates/site/externals.js.ts.gotxt",
	"website":          "templates/site/website.js.ts.gotxt",
	"init_menu":        "templates/site/init_menu.js.ts.gotxt",
	"menu":             "templates/site/menu.js.ts.gotxt",
	"dotenv":           "templates/misc/env.gotxt",
	"project_settings": "templates/misc/sveltin.json.gotxt",
	"readme":           "templates/misc/README.md.gotxt",
	"license":          "templates/misc/LICENSE.gotxt",
	"index":            "templates/themes/index.svelte.gotxt",
	"index_notheme":    "templates/themes/index.notheme.svelte.gotxt",
	"index_pageload":   "templates/themes/index.server.ts.gotxt",
	"theme_config":     "templates/themes/theme.config.js.gotxt",
}

// ResourceFilesMap is the map for the resource template files.
var ResourceFilesMap = EmbeddedFSEntry{
	"lib":            "templates/resource/lib.gotxt",
	"index_blank":    "templates/resource/themes/blank/page.svelte.gotxt",
	"index_sveltin":  "templates/resource/themes/sveltin/page.svelte.gotxt",
	"index_pageload": "templates/resource/page.server.ts.gotxt",
	"slug_blank":     "templates/resource/themes/blank/slug.svelte.gotxt",
	"slug_sveltin":   "templates/resource/themes/sveltin/slug.svelte.gotxt",
	"slug_pageload":  "templates/resource/slug.server.ts.gotxt",
	"slug_layout":    "templates/resource/layout.svelte.gotxt",
}

// APIFilesMap is the map for the api template files.
var APIFilesMap = EmbeddedFSEntry{
	"api_index":           "templates/resource/api/apiIndex.gotxt",
	"api_slug":            "templates/resource/api/apiSlug.gotxt",
	"api_metadata_index":  "templates/resource/api/apiMetadataIndex.gotxt",
	"api_metadata_single": "templates/resource/api/apiMetadataSingle.gotxt",
	"api_metadata_list":   "templates/resource/api/apiMetadataList.gotxt",
}

// MatchersFilesMap is the map for the matchers template files.
var MatchersFilesMap = EmbeddedFSEntry{
	"string_matcher":  "templates/resource/matchers/string.js.gotxt",
	"generic_matcher": "templates/resource/matchers/generic.js.gotxt",
}

// MetadataFilesMap is the map for the metadata template files.
var MetadataFilesMap = EmbeddedFSEntry{
	"lib_single":     "templates/resource/metadata/libSingle.gotxt",
	"lib_list":       "templates/resource/metadata/libList.gotxt",
	"index_blank":    "templates/resource/metadata/themes/blank/page.svelte.gotxt",
	"index_sveltin":  "templates/resource/metadata/themes/sveltin/page.svelte.gotxt",
	"index_pageload": "templates/resource/metadata/page.ts.gotxt",
	"slug_blank":     "templates/resource/metadata/themes/blank/slug.svelte.gotxt",
	"slug_sveltin":   "templates/resource/metadata/themes/sveltin/slug.svelte.gotxt",
	"slug_pageload":  "templates/resource/metadata/slug.server.ts.gotxt",
}

// PageFilesMap is the map for the page template files.
var PageFilesMap = EmbeddedFSEntry{
	"svelte_blank":     "templates/page/themes/blank/page.svelte.gotxt",
	"svelte_sveltin":   "templates/page/themes/sveltin/page.svelte.gotxt",
	"markdown_blank":   "templates/page/themes/blank/page.svx.gotxt",
	"markdown_sveltin": "templates/page/themes/sveltin/page.svx.gotxt",
	"index_pageload":   "templates/page/page.server.ts.gotxt",
}

// ContentFilesMap is the map for the content template files.
var ContentFilesMap = EmbeddedFSEntry{
	"blank":  "templates/content/blank.svx.gotxt",
	"sample": "templates/content/sample.svx.gotxt",
}

// XMLFilesMap is a map for the xml (sitemap and rss) template files.
var XMLFilesMap = EmbeddedFSEntry{
	"sitemap_static": "templates/xml/sitemap.xml.gotxt",
	"rss_static":     "templates/xml/rss.xml.gotxt",
	"sitemap_ssr":    "templates/xml/ssr_sitemap.xml.ts.gotxt",
	"rss_ssr":        "templates/xml/ssr_rss.xml.ts.gotxt",
}

//=============================================================================

// BootstrapSveltinThemeFilesMap is a map for the styled templates file whe using bootstrap.
var BootstrapSveltinThemeFilesMap = EmbeddedFSEntry{
	"package_json":   "templates/themes/sveltin/bootstrap/package.json.gotxt",
	"svelte_config":  "templates/themes/sveltin/bootstrap/svelte.config.js",
	"vite_config":    "templates/themes/sveltin/bootstrap/vite.config.ts.gotxt",
	"layout":         "templates/themes/sveltin/bootstrap/layout.svelte.gotxt",
	"md_layout":      "templates/themes/sveltin/bootstrap/md-layout.svelte.gotxt",
	"layout_ts":      "templates/themes/layout.ts.gotxt",
	"mdsvex_config":  "templates/themes/mdsvex.config.js.gotxt",
	"app_html":       "templates/themes/sveltin/bootstrap/app.html",
	"reset_css":      "templates/themes/tw-preflight.css",
	"app_css":        "templates/themes/sveltin/bootstrap/app.scss",
	"variables_scss": "templates/themes/sveltin/bootstrap/variables.scss",
	"cta":            "templates/themes/sveltin/bootstrap/CTA.svelte",
	"footer":         "templates/themes/sveltin/bootstrap/Footer.svelte",
	"error":          "templates/themes/error.styled.svelte",
}

// BootstrapBlankThemeFilesMap is the map for the unstyled templates file whe using bootstrap.
var BootstrapBlankThemeFilesMap = EmbeddedFSEntry{
	"package_json":   "templates/themes/blank/bootstrap/package.json.gotxt",
	"svelte_config":  "templates/themes/blank/bootstrap/svelte.config.js",
	"vite_config":    "templates/themes/blank/bootstrap/vite.config.ts.gotxt",
	"app_html":       "templates/themes/blank/bootstrap/app.html",
	"layout":         "templates/themes/blank/bootstrap/layout.svelte.gotxt",
	"md_layout":      "templates/themes/blank/bootstrap/md-layout.svelte.gotxt",
	"layout_ts":      "templates/themes/layout.ts.gotxt",
	"mdsvex_config":  "templates/themes/mdsvex.config.js.gotxt",
	"reset_css":      "templates/themes/tw-preflight.css",
	"app_css":        "templates/themes/blank/bootstrap/app.scss",
	"variables_scss": "templates/themes/blank/bootstrap/variables.scss",
	"cta":            "templates/themes/blank/bootstrap/CTA.svelte",
	"error":          "templates/themes/error.unstyled.svelte",
}

//=============================================================================

// BulmaSveltinThemeFilesMap is the map for the styled templates file whe using bulma.
var BulmaSveltinThemeFilesMap = EmbeddedFSEntry{
	"package_json":   "templates/themes/sveltin/bulma/package.json.gotxt",
	"svelte_config":  "templates/themes/sveltin/bulma/svelte.config.js",
	"vite_config":    "templates/themes/sveltin/bulma/vite.config.ts.gotxt",
	"layout":         "templates/themes/sveltin/bulma/layout.svelte.gotxt",
	"md_layout":      "templates/themes/sveltin/bulma/md-layout.svelte.gotxt",
	"layout_ts":      "templates/themes/layout.ts.gotxt",
	"mdsvex_config":  "templates/themes/mdsvex.config.js.gotxt",
	"app_html":       "templates/themes/sveltin/bulma/app.html",
	"reset_css":      "templates/themes/tw-preflight.css",
	"app_css":        "templates/themes/sveltin/bulma/app.scss",
	"variables_scss": "templates/themes/sveltin/bulma/variables.scss",
	"cta":            "templates/themes/sveltin/bulma/CTA.svelte",
	"footer":         "templates/themes/sveltin/bulma/Footer.svelte",
	"error":          "templates/themes/error.styled.svelte",
}

// BulmaBlankThemeFilesMap is the map for the unstyled templates file whe using bulma.
var BulmaBlankThemeFilesMap = EmbeddedFSEntry{
	"package_json":   "templates/themes/blank/bulma/package.json.gotxt",
	"svelte_config":  "templates/themes/blank/bulma/svelte.config.js",
	"vite_config":    "templates/themes/blank/bulma/vite.config.ts.gotxt",
	"app_html":       "templates/themes/blank/bulma/app.html",
	"layout":         "templates/themes/blank/bulma/layout.svelte.gotxt",
	"md_layout":      "templates/themes/blank/bulma/md-layout.svelte.gotxt",
	"layout_ts":      "templates/themes/layout.ts.gotxt",
	"mdsvex_config":  "templates/themes/mdsvex.config.js.gotxt",
	"reset_css":      "templates/themes/tw-preflight.css",
	"app_css":        "templates/themes/blank/bulma/app.scss",
	"variables_scss": "templates/themes/blank/bulma/variables.scss",
	"cta":            "templates/themes/blank/bulma/CTA.svelte",
	"error":          "templates/themes/error.unstyled.svelte",
}

//=============================================================================

// SassSveltinThemeFilesMap is the map for the styled templates file whe using scss/sass.
var SassSveltinThemeFilesMap = EmbeddedFSEntry{
	"package_json":   "templates/themes/sveltin/scss/package.json.gotxt",
	"svelte_config":  "templates/themes/sveltin/scss/svelte.config.js",
	"vite_config":    "templates/themes/sveltin/scss/vite.config.ts.gotxt",
	"app_html":       "templates/themes/sveltin/scss/app.html",
	"layout":         "templates/themes/sveltin/scss/layout.svelte.gotxt",
	"md_layout":      "templates/themes/sveltin/scss/md-layout.svelte.gotxt",
	"layout_ts":      "templates/themes/layout.ts.gotxt",
	"mdsvex_config":  "templates/themes/mdsvex.config.js.gotxt",
	"reset_css":      "templates/themes/tw-preflight.css",
	"app_css":        "templates/themes/sveltin/scss/app.scss",
	"variables_scss": "templates/themes/sveltin/scss/variables.scss",
	"cta":            "templates/themes/sveltin/scss/CTA.svelte",
	"footer":         "templates/themes/sveltin/scss/Footer.svelte",
	"error":          "templates/themes/error.styled.svelte",
}

// SassBlankThemeFilesMap is the map for the unstyled templates file whe using scss/sass.
var SassBlankThemeFilesMap = EmbeddedFSEntry{
	"package_json":   "templates/themes/blank/scss/package.json.gotxt",
	"svelte_config":  "templates/themes/blank/scss/svelte.config.js",
	"vite_config":    "templates/themes/blank/scss/vite.config.ts.gotxt",
	"layout":         "templates/themes/blank/scss/layout.svelte.gotxt",
	"md_layout":      "templates/themes/blank/scss/md-layout.svelte.gotxt",
	"layout_ts":      "templates/themes/layout.ts.gotxt",
	"mdsvex_config":  "templates/themes/mdsvex.config.js.gotxt",
	"app_html":       "templates/themes/blank/scss/app.html",
	"reset_css":      "templates/themes/tw-preflight.css",
	"app_css":        "templates/themes/blank/scss/app.scss",
	"variables_scss": "templates/themes/blank/scss/variables.scss",
	"cta":            "templates/themes/blank/scss/CTA.svelte",
	"error":          "templates/themes/error.unstyled.svelte",
}

//=============================================================================

// TailwindSveltinThemeFilesMap is the map for the styled templates file whe using tailwind css.
var TailwindSveltinThemeFilesMap = EmbeddedFSEntry{
	"package_json":        "templates/themes/sveltin/tailwindcss/package.json.gotxt",
	"svelte_config":       "templates/themes/sveltin/tailwindcss/svelte.config.js",
	"vite_config":         "templates/themes/sveltin/tailwindcss/vite.config.ts.gotxt",
	"tailwind_css_config": "templates/themes/sveltin/tailwindcss/tailwind.config.cjs",
	"layout":              "templates/themes/sveltin/tailwindcss/layout.svelte.gotxt",
	"md_layout":           "templates/themes/sveltin/tailwindcss/md-layout.svelte.gotxt",
	"layout_ts":           "templates/themes/layout.ts.gotxt",
	"mdsvex_config":       "templates/themes/mdsvex.config.js.gotxt",
	"app_html":            "templates/themes/sveltin/tailwindcss/app.html",
	"postcss":             "templates/themes/sveltin/tailwindcss/postcss.config.cjs",
	"app_css":             "templates/themes/sveltin/tailwindcss/app.css",
	"cta":                 "templates/themes/sveltin/tailwindcss/CTA.svelte",
	"footer":              "templates/themes/sveltin/tailwindcss/Footer.svelte",
	"error":               "templates/themes/error.styled.svelte",
}

// TailwindBlankThemeFilesMap is the map for the unstyled templates file whe using tailwind css.
var TailwindBlankThemeFilesMap = EmbeddedFSEntry{
	"package_json":        "templates/themes/blank/tailwindcss/package.json.gotxt",
	"svelte_config":       "templates/themes/blank/tailwindcss/svelte.config.js",
	"vite_config":         "templates/themes/blank/tailwindcss/vite.config.ts.gotxt",
	"tailwind_css_config": "templates/themes/blank/tailwindcss/tailwind.config.cjs",
	"postcss":             "templates/themes/blank/tailwindcss/postcss.config.cjs",
	"layout":              "templates/themes/blank/tailwindcss/layout.svelte.gotxt",
	"md_layout":           "templates/themes/blank/tailwindcss/md-layout.svelte.gotxt",
	"layout_ts":           "templates/themes/layout.ts.gotxt",
	"mdsvex_config":       "templates/themes/mdsvex.config.js.gotxt",
	"app_html":            "templates/themes/blank/tailwindcss/app.html",
	"app_css":             "templates/themes/blank/tailwindcss/app.css",
	"cta":                 "templates/themes/blank/tailwindcss/CTA.svelte",
	"error":               "templates/themes/error.unstyled.svelte",
}

//=============================================================================

// UnoCSSSveltinThemeFilesMap is the map for the styled templates file whe using tailwind css.
var UnoCSSSveltinThemeFilesMap = EmbeddedFSEntry{
	"package_json":  "templates/themes/sveltin/unocss/package.json.gotxt",
	"svelte_config": "templates/themes/sveltin/unocss/svelte.config.js",
	"vite_config":   "templates/themes/sveltin/unocss/vite.config.ts.gotxt",
	"unocss_config": "templates/themes/sveltin/unocss/unocss.config.ts",
	"layout":        "templates/themes/sveltin/unocss/layout.svelte.gotxt",
	"md_layout":     "templates/themes/sveltin/unocss/md-layout.svelte.gotxt",
	"layout_ts":     "templates/themes/layout.ts.gotxt",
	"mdsvex_config": "templates/themes/mdsvex.config.js.gotxt",
	"app_html":      "templates/themes/sveltin/unocss/app.html",
	"postcss":       "templates/themes/sveltin/unocss/postcss.config.cjs",
	"app_css":       "templates/themes/sveltin/unocss/app.css",
	"cta":           "templates/themes/sveltin/unocss/CTA.svelte",
	"footer":        "templates/themes/sveltin/unocss/Footer.svelte",
	"error":         "templates/themes/error.styled.svelte",
}

// UnoCSSBlankThemeFilesMap is the map for the unstyled templates file whe using tailwind css.
var UnoCSSBlankThemeFilesMap = EmbeddedFSEntry{
	"package_json":  "templates/themes/blank/unocss/package.json.gotxt",
	"svelte_config": "templates/themes/blank/unocss/svelte.config.js",
	"vite_config":   "templates/themes/blank/unocss/vite.config.ts.gotxt",
	"unocss_config": "templates/themes/blank/unocss/unocss.config.ts",
	"postcss":       "templates/themes/blank/unocss/postcss.config.cjs",
	"layout":        "templates/themes/blank/unocss/layout.svelte.gotxt",
	"md_layout":     "templates/themes/blank/unocss/md-layout.svelte.gotxt",
	"layout_ts":     "templates/themes/layout.ts.gotxt",
	"mdsvex_config": "templates/themes/mdsvex.config.js.gotxt",
	"app_html":      "templates/themes/blank/unocss/app.html",
	"app_css":       "templates/themes/blank/unocss/app.css",
	"cta":           "templates/themes/blank/unocss/CTA.svelte",
	"error":         "templates/themes/error.unstyled.svelte",
}

//=============================================================================

// VanillaSveltinThemeFilesMap is the map for the styled templates file whe using vanilla css.
var VanillaSveltinThemeFilesMap = EmbeddedFSEntry{
	"package_json":  "templates/themes/sveltin/vanillacss/package.json.gotxt",
	"svelte_config": "templates/themes/sveltin/vanillacss/svelte.config.js",
	"vite_config":   "templates/themes/sveltin/vanillacss/vite.config.ts.gotxt",
	"app_html":      "templates/themes/sveltin/vanillacss/app.html",
	"layout":        "templates/themes/sveltin/vanillacss/layout.svelte.gotxt",
	"md_layout":     "templates/themes/sveltin/vanillacss/md-layout.svelte.gotxt",
	"layout_ts":     "templates/themes/layout.ts.gotxt",
	"mdsvex_config": "templates/themes/mdsvex.config.js.gotxt",
	"reset_css":     "templates/themes/tw-preflight.css",
	"app_css":       "templates/themes/sveltin/vanillacss/app.css",
	"cta":           "templates/themes/sveltin/vanillacss/CTA.svelte",
	"footer":        "templates/themes/sveltin/vanillacss/Footer.svelte",
	"error":         "templates/themes/error.styled.svelte",
}

// VanillaBlankThemeFilesMap is the map for the unstyled templates file whe using vanilla css.
var VanillaBlankThemeFilesMap = EmbeddedFSEntry{
	"package_json":  "templates/themes/blank/vanillacss/package.json.gotxt",
	"svelte_config": "templates/themes/blank/vanillacss/svelte.config.js",
	"vite_config":   "templates/themes/blank/vanillacss/vite.config.ts.gotxt",
	"app_html":      "templates/themes/blank/vanillacss/app.html",
	"layout":        "templates/themes/blank/vanillacss/layout.svelte.gotxt",
	"md_layout":     "templates/themes/blank/vanillacss/md-layout.svelte.gotxt",
	"layout_ts":     "templates/themes/layout.ts.gotxt",
	"mdsvex_config": "templates/themes/mdsvex.config.js.gotxt",
	"reset_css":     "templates/themes/tw-preflight.css",
	"app_css":       "templates/themes/blank/vanillacss/app.css",
	"cta":           "templates/themes/blank/vanillacss/CTA.svelte",
	"error":         "templates/themes/error.unstyled.svelte",
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package resources ...
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

// SveltinFS is the name for the embedded FS used by Sveltin.
//
//go:embed internal/templates/*
var SveltinFS embed.FS

// SveltinFSItem represents an entry for the embedded FS.
type SveltinFSItem map[string]string

// SveltinProjectFS is the map for the project template files.
var SveltinProjectFS = map[string]string{
	"defaults":      "internal/templates/site/defaults.js.ts.gotxt",
	"externals":     "internal/templates/site/externals.js.ts.gotxt",
	"website":       "internal/templates/site/website.js.ts.gotxt",
	"init_menu":     "internal/templates/site/init_menu.js.ts.gotxt",
	"menu":          "internal/templates/site/menu.js.ts.gotxt",
	"dotenv":        "internal/templates/misc/env.gotxt",
	"readme":        "internal/templates/misc/README.md.gotxt",
	"license":       "internal/templates/misc/LICENSE.gotxt",
	"index":         "internal/templates/themes/index.svelte.gotxt",
	"index_notheme": "internal/templates/themes/index.notheme.svelte.gotxt",
	"theme_config":  "internal/templates/themes/theme.config.js.gotxt",
}

// SveltinResourceFS is the map for the resource template files.
var SveltinResourceFS = map[string]string{
	"lib":           "internal/templates/resource/lib.gotxt",
	"index":         "internal/templates/resource/page.svelte.gotxt",
	"indexendpoint": "internal/templates/resource/page.server.ts.gotxt",
	"slug":          "internal/templates/resource/slug.svelte.gotxt",
	"slugendpoint":  "internal/templates/resource/slug.ts.gotxt",
	"sluglayout":    "internal/templates/resource/layout.svelte.gotxt",
}

// SveltinAPIFS is the map for the api template files.
var SveltinAPIFS = map[string]string{
	"api_index":           "internal/templates/resource/api/apiIndex.gotxt",
	"api_slug":            "internal/templates/resource/api/apiSlug.gotxt",
	"api_metadata_index":  "internal/templates/resource/api/apiMetadataIndex.gotxt",
	"api_metadata_single": "internal/templates/resource/api/apiMetadataSingle.gotxt",
	"api_metadata_list":   "internal/templates/resource/api/apiMetadataList.gotxt",
}

// SveltinMatchersFS is the map for the matchers template files.
var SveltinMatchersFS = map[string]string{
	"string_matcher":  "internal/templates/resource/matchers/string.js.gotxt",
	"generic_matcher": "internal/templates/resource/matchers/generic.js.gotxt",
}

// SveltinMetadataFS is the map for the metadata template files.
var SveltinMetadataFS = map[string]string{
	"lib_single":    "internal/templates/resource/metadata/libSingle.gotxt",
	"lib_list":      "internal/templates/resource/metadata/libList.gotxt",
	"index":         "internal/templates/resource/metadata/page.svelte.gotxt",
	"indexendpoint": "internal/templates/resource/metadata/page.server.ts.gotxt",
	"slug":          "internal/templates/resource/metadata/slug.svelte.gotxt",
	"slugendpoint":  "internal/templates/resource/metadata/slug.ts.gotxt",
}

// SveltinPageFS is the map for the page template files.
var SveltinPageFS = map[string]string{
	"svelte":   "internal/templates/page/page.svelte.gotxt",
	"markdown": "internal/templates/page/page.svx.gotxt",
}

// SveltinContentFS is the map for the content template files.
var SveltinContentFS = map[string]string{
	"blank":  "internal/templates/content/blank.svx.gotxt",
	"sample": "internal/templates/content/sample.svx.gotxt",
}

// SveltinXMLFS is a map for the xml (sitemap and rss) template files.
var SveltinXMLFS = map[string]string{
	"sitemap_static": "internal/templates/xml/sitemap.xml.gotxt",
	"rss_static":     "internal/templates/xml/rss.xml.gotxt",
	"sitemap_ssr":    "internal/templates/xml/ssr_sitemap.xml.ts.gotxt",
	"rss_ssr":        "internal/templates/xml/ssr_rss.xml.ts.gotxt",
}

//=============================================================================

// BootstrapSveltinThemeFS is a map for the styled templates file whe using bootstrap.
var BootstrapSveltinThemeFS = SveltinFSItem{
	"package_json":   "internal/templates/themes/sveltin/bootstrap/package.json.gotxt",
	"svelte_config":  "internal/templates/themes/sveltin/bootstrap/svelte.config.js",
	"vite_config":    "internal/templates/themes/sveltin/bootstrap/vite.config.ts.gotxt",
	"layout":         "internal/templates/themes/sveltin/bootstrap/layout.svelte.gotxt",
	"layout_ts":      "internal/templates/themes/layout.ts.gotxt",
	"app_html":       "internal/templates/themes/sveltin/bootstrap/app.html",
	"app_css":        "internal/templates/themes/sveltin/bootstrap/app.scss",
	"variables_scss": "internal/templates/themes/sveltin/bootstrap/variables.scss",
	"hero":           "internal/templates/themes/sveltin/bootstrap/Hero.svelte",
	"footer":         "internal/templates/themes/sveltin/bootstrap/Footer.svelte",
	"error":          "internal/templates/themes/error.styled.svelte",
}

// BootstrapBlankThemeFS is the map for the unstyled templates file whe using bootstrap.
var BootstrapBlankThemeFS = SveltinFSItem{
	"package_json":   "internal/templates/themes/blank/bootstrap/package.json.gotxt",
	"svelte_config":  "internal/templates/themes/blank/bootstrap/svelte.config.js",
	"vite_config":    "internal/templates/themes/blank/bootstrap/vite.config.ts.gotxt",
	"app_html":       "internal/templates/themes/blank/bootstrap/app.html",
	"layout":         "internal/templates/themes/blank/bootstrap/layout.svelte.gotxt",
	"layout_ts":      "internal/templates/themes/layout.ts.gotxt",
	"app_css":        "internal/templates/themes/blank/bootstrap/app.scss",
	"variables_scss": "internal/templates/themes/blank/bootstrap/variables.scss",
	"hero":           "internal/templates/themes/blank/bootstrap/Hero.svelte",
	"error":          "internal/templates/themes/error.unstyled.svelte",
}

//=============================================================================

// BulmaSveltinThemeFS is the map for the styled templates file whe using bulma.
var BulmaSveltinThemeFS = SveltinFSItem{
	"package_json":   "internal/templates/themes/sveltin/bulma/package.json.gotxt",
	"svelte_config":  "internal/templates/themes/sveltin/bulma/svelte.config.js",
	"vite_config":    "internal/templates/themes/sveltin/bulma/vite.config.ts.gotxt",
	"layout":         "internal/templates/themes/sveltin/bulma/layout.svelte.gotxt",
	"layout_ts":      "internal/templates/themes/layout.ts.gotxt",
	"app_html":       "internal/templates/themes/sveltin/bulma/app.html",
	"app_css":        "internal/templates/themes/sveltin/bulma/app.scss",
	"variables_scss": "internal/templates/themes/sveltin/bulma/variables.scss",
	"hero":           "internal/templates/themes/sveltin/bulma/Hero.svelte",
	"footer":         "internal/templates/themes/sveltin/bulma/Footer.svelte",
	"error":          "internal/templates/themes/error.styled.svelte",
}

// BulmaBlankThemeFS is the map for the unstyled templates file whe using bulma.
var BulmaBlankThemeFS = SveltinFSItem{
	"package_json":   "internal/templates/themes/blank/bulma/package.json.gotxt",
	"svelte_config":  "internal/templates/themes/blank/bulma/svelte.config.js",
	"vite_config":    "internal/templates/themes/blank/bulma/vite.config.ts.gotxt",
	"app_html":       "internal/templates/themes/blank/bulma/app.html",
	"layout":         "internal/templates/themes/blank/bulma/layout.svelte.gotxt",
	"layout_ts":      "internal/templates/themes/layout.ts.gotxt",
	"app_css":        "internal/templates/themes/blank/bulma/app.scss",
	"variables_scss": "internal/templates/themes/blank/bulma/variables.scss",
	"hero":           "internal/templates/themes/blank/bulma/Hero.svelte",
	"error":          "internal/templates/themes/error.unstyled.svelte",
}

//=============================================================================

// SCSSSveltinThemeFS is the map for the styled templates file whe using scss/sass.
var SCSSSveltinThemeFS = SveltinFSItem{
	"package_json":   "internal/templates/themes/sveltin/scss/package.json.gotxt",
	"svelte_config":  "internal/templates/themes/sveltin/scss/svelte.config.js",
	"vite_config":    "internal/templates/themes/sveltin/scss/vite.config.ts.gotxt",
	"app_html":       "internal/templates/themes/sveltin/scss/app.html",
	"layout":         "internal/templates/themes/sveltin/scss/layout.svelte.gotxt",
	"layout_ts":      "internal/templates/themes/layout.ts.gotxt",
	"app_css":        "internal/templates/themes/sveltin/scss/app.scss",
	"variables_scss": "internal/templates/themes/sveltin/scss/variables.scss",
	"hero":           "internal/templates/themes/sveltin/scss/Hero.svelte",
	"footer":         "internal/templates/themes/sveltin/scss/Footer.svelte",
	"error":          "internal/templates/themes/error.styled.svelte",
}

// SCSSBlankThemeFS is the map for the unstyled templates file whe using scss/sass.
var SCSSBlankThemeFS = SveltinFSItem{
	"package_json":   "internal/templates/themes/blank/scss/package.json.gotxt",
	"svelte_config":  "internal/templates/themes/blank/scss/svelte.config.js",
	"vite_config":    "internal/templates/themes/blank/scss/vite.config.ts.gotxt",
	"layout":         "internal/templates/themes/blank/scss/layout.svelte.gotxt",
	"layout_ts":      "internal/templates/themes/layout.ts.gotxt",
	"app_html":       "internal/templates/themes/blank/scss/app.html",
	"app_css":        "internal/templates/themes/blank/scss/app.scss",
	"variables_scss": "internal/templates/themes/blank/scss/variables.scss",
	"hero":           "internal/templates/themes/blank/scss/Hero.svelte",
	"error":          "internal/templates/themes/error.unstyled.svelte",
}

//=============================================================================

// TailwindSveltinThemeFS is the map for the styled templates file whe using tailwind css.
var TailwindSveltinThemeFS = SveltinFSItem{
	"package_json":        "internal/templates/themes/sveltin/tailwindcss/package.json.gotxt",
	"svelte_config":       "internal/templates/themes/sveltin/tailwindcss/svelte.config.js",
	"vite_config":         "internal/templates/themes/sveltin/tailwindcss/vite.config.ts.gotxt",
	"tailwind_css_config": "internal/templates/themes/sveltin/tailwindcss/tailwind.config.cjs",
	"layout":              "internal/templates/themes/sveltin/tailwindcss/layout.svelte.gotxt",
	"layout_ts":           "internal/templates/themes/layout.ts.gotxt",
	"app_html":            "internal/templates/themes/sveltin/tailwindcss/app.html",
	"postcss":             "internal/templates/themes/sveltin/tailwindcss/postcss.config.cjs",
	"app_css":             "internal/templates/themes/sveltin/tailwindcss/app.css",
	"hero":                "internal/templates/themes/sveltin/tailwindcss/Hero.svelte",
	"footer":              "internal/templates/themes/sveltin/tailwindcss/Footer.svelte",
	"error":               "internal/templates/themes/error.styled.svelte",
}

// TailwindBlankThemeFS is the map for the unstyled templates file whe using tailwind css.
var TailwindBlankThemeFS = SveltinFSItem{
	"package_json":        "internal/templates/themes/blank/tailwindcss/package.json.gotxt",
	"svelte_config":       "internal/templates/themes/blank/tailwindcss/svelte.config.js",
	"vite_config":         "internal/templates/themes/blank/tailwindcss/vite.config.ts.gotxt",
	"tailwind_css_config": "internal/templates/themes/blank/tailwindcss/tailwind.config.cjs",
	"postcss":             "internal/templates/themes/blank/tailwindcss/postcss.config.cjs",
	"layout":              "internal/templates/themes/blank/tailwindcss/layout.svelte.gotxt",
	"layout_ts":           "internal/templates/themes/layout.ts.gotxt",
	"app_html":            "internal/templates/themes/blank/tailwindcss/app.html",
	"app_css":             "internal/templates/themes/blank/tailwindcss/app.css",
	"hero":                "internal/templates/themes/blank/tailwindcss/Hero.svelte",
	"error":               "internal/templates/themes/error.unstyled.svelte",
}

//=============================================================================

// VanillaSveltinThemeFS is the map for the styled templates file whe using vanilla css.
var VanillaSveltinThemeFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/sveltin/vanillacss/package.json.gotxt",
	"svelte_config": "internal/templates/themes/sveltin/vanillacss/svelte.config.js",
	"vite_config":   "internal/templates/themes/sveltin/vanillacss/vite.config.ts.gotxt",
	"app_html":      "internal/templates/themes/sveltin/vanillacss/app.html",
	"layout":        "internal/templates/themes/sveltin/vanillacss/layout.svelte.gotxt",
	"layout_ts":     "internal/templates/themes/layout.ts.gotxt",
	"app_css":       "internal/templates/themes/sveltin/vanillacss/app.css",
	"hero":          "internal/templates/themes/sveltin/vanillacss/Hero.svelte",
	"footer":        "internal/templates/themes/sveltin/vanillacss/Footer.svelte",
	"error":         "internal/templates/themes/error.styled.svelte",
}

// VanillaBlankThemeFS is the map for the unstyled templates file whe using vanilla css.
var VanillaBlankThemeFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/blank/vanillacss/package.json.gotxt",
	"svelte_config": "internal/templates/themes/blank/vanillacss/svelte.config.js",
	"vite_config":   "internal/templates/themes/blank/vanillacss/vite.config.ts.gotxt",
	"app_html":      "internal/templates/themes/blank/vanillacss/app.html",
	"layout":        "internal/templates/themes/blank/vanillacss/layout.svelte.gotxt",
	"layout_ts":     "internal/templates/themes/layout.ts.gotxt",
	"app_css":       "internal/templates/themes/blank/vanillacss/app.css",
	"hero":          "internal/templates/themes/blank/vanillacss/Hero.svelte",
	"error":         "internal/templates/themes/error.unstyled.svelte",
}

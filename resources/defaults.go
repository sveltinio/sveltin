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
	"index":         "internal/templates/resource/metadata/index.gotxt",
	"indexendpoint": "internal/templates/resource/metadata/index.ts.gotxt",
	"slug":          "internal/templates/resource/metadata/slug.gotxt",
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

// SveltinVanillaFS is themap for the default templates file whe using plain css.
var SveltinVanillaFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/vanillacss/package.json.gotxt",
	"svelte_config": "internal/templates/themes/vanillacss/svelte.config.js",
	"vite_config":   "internal/templates/themes/vanillacss/vite.config.js.gotxt",
	"app_html":      "internal/templates/themes/vanillacss/app.html",
}

// SveltinVanillaStyledFS is the map for the styled templates file whe using vanilla css.
var SveltinVanillaStyledFS = SveltinFSItem{
	"layout":  "internal/templates/themes/vanillacss/styled/layout.svelte.gotxt",
	"app_css": "internal/templates/themes/vanillacss/styled/app.css",
	"hero":    "internal/templates/themes/vanillacss/styled/Hero.svelte",
	"footer":  "internal/templates/themes/vanillacss/styled/Footer.svelte",
}

// SveltinVanillaUnstyledFS is the map for the unstyled templates file whe using vanilla css.
var SveltinVanillaUnstyledFS = SveltinFSItem{
	"layout":  "internal/templates/themes/vanillacss/unstyled/layout.svelte.gotxt",
	"app_css": "internal/templates/themes/vanillacss/unstyled/app.css",
	"hero":    "internal/templates/themes/vanillacss/unstyled/Hero.svelte",
}

//=============================================================================

// SveltinTailwindLibFS is the map for the default templates file whe using tailwind css.
var SveltinTailwindLibFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/tailwindcss/package.json.gotxt",
	"svelte_config": "internal/templates/themes/tailwindcss/svelte.config.js",
	"vite_config":   "internal/templates/themes/tailwindcss/vite.config.js.gotxt",
	"app_html":      "internal/templates/themes/tailwindcss/app.html",
	"postcss":       "internal/templates/themes/tailwindcss/postcss.config.cjs",
}

// SveltinTailwindLibStyledFS is the map for the styled templates file whe using tailwind css.
var SveltinTailwindLibStyledFS = SveltinFSItem{
	"layout":              "internal/templates/themes/tailwindcss/styled/layout.svelte.gotxt",
	"tailwind_css_config": "internal/templates/themes/tailwindcss/styled/tailwind.config.cjs",
	"app_css":             "internal/templates/themes/tailwindcss/styled/app.css",
	"hero":                "internal/templates/themes/tailwindcss/styled/Hero.svelte",
	"footer":              "internal/templates/themes/tailwindcss/styled/Footer.svelte",
}

// SveltinTailwindLibUnstyledFS is the map for the unstyled templates file whe using tailwind css.
var SveltinTailwindLibUnstyledFS = SveltinFSItem{
	"layout":              "internal/templates/themes/tailwindcss/unstyled/layout.svelte.gotxt",
	"tailwind_css_config": "internal/templates/themes/tailwindcss/unstyled/tailwind.config.cjs",
	"app_css":             "internal/templates/themes/tailwindcss/unstyled/app.css",
	"hero":                "internal/templates/themes/tailwindcss/unstyled/Hero.svelte",
}

//=============================================================================

// SveltinBulmaLibFS is the map for the default templates file whe using bulma.
var SveltinBulmaLibFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/bulma/package.json.gotxt",
	"svelte_config": "internal/templates/themes/bulma/svelte.config.js",
	"vite_config":   "internal/templates/themes/bulma/vite.config.js.gotxt",
	"app_html":      "internal/templates/themes/bulma/app.html",
}

// SveltinBulmaLibStyledFS is the map for the styled templates file whe using bulma.
var SveltinBulmaLibStyledFS = SveltinFSItem{
	"layout":         "internal/templates/themes/bulma/styled/layout.svelte.gotxt",
	"app_css":        "internal/templates/themes/bulma/styled/app.scss",
	"variables_scss": "internal/templates/themes/bulma/styled/variables.scss",
	"hero":           "internal/templates/themes/bulma/styled/Hero.svelte",
	"footer":         "internal/templates/themes/bulma/styled/Footer.svelte",
}

// SveltinBulmaLibUnstyledFS is the map for the unstyled templates file whe using bulma.
var SveltinBulmaLibUnstyledFS = SveltinFSItem{
	"layout":         "internal/templates/themes/bulma/unstyled/layout.svelte.gotxt",
	"app_css":        "internal/templates/themes/bulma/unstyled/app.scss",
	"variables_scss": "internal/templates/themes/bulma/unstyled/variables.scss",
	"hero":           "internal/templates/themes/bulma/unstyled/Hero.svelte",
}

//=============================================================================

// SveltinBootstrapLibFS is the map for the default templates file whe using bootstrap.
var SveltinBootstrapLibFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/bootstrap/package.json.gotxt",
	"svelte_config": "internal/templates/themes/bootstrap/svelte.config.js",
	"vite_config":   "internal/templates/themes/bootstrap/vite.config.js.gotxt",
	"app_html":      "internal/templates/themes/bootstrap/app.html",
}

// SveltinBootstrapLibStyledFS is a map for the styled templates file whe using bootstrap.
var SveltinBootstrapLibStyledFS = SveltinFSItem{
	"layout":         "internal/templates/themes/bootstrap/styled/layout.svelte.gotxt",
	"app_css":        "internal/templates/themes/bootstrap/styled/app.scss",
	"variables_scss": "internal/templates/themes/bootstrap/styled/variables.scss",
	"hero":           "internal/templates/themes/bootstrap/styled/Hero.svelte",
	"footer":         "internal/templates/themes/bootstrap/styled/Footer.svelte",
}

// SveltinBootstrapLibUnstyledFS is the map for the unstyled templates file whe using bootstrap.
var SveltinBootstrapLibUnstyledFS = SveltinFSItem{
	"layout":         "internal/templates/themes/bootstrap/unstyled/layout.svelte.gotxt",
	"app_css":        "internal/templates/themes/bootstrap/unstyled/app.scss",
	"variables_scss": "internal/templates/themes/bootstrap/unstyled/variables.scss",
	"hero":           "internal/templates/themes/bootstrap/unstyled/Hero.svelte",
}

//=============================================================================

// SveltinSCSSLibFS is a map for the default templates file whe using scss/sass.
var SveltinSCSSLibFS = SveltinFSItem{
	"package_json":  "internal/templates/themes/scss/package.json.gotxt",
	"svelte_config": "internal/templates/themes/scss/svelte.config.js",
	"vite_config":   "internal/templates/themes/scss/vite.config.js.gotxt",
	"app_html":      "internal/templates/themes/scss/app.html",
}

// SveltinSCSSLibStyledFS is the map for the styled templates file whe using scss/sass.
var SveltinSCSSLibStyledFS = SveltinFSItem{
	"layout":         "internal/templates/themes/scss/styled/layout.svelte.gotxt",
	"app_css":        "internal/templates/themes/scss/styled/app.scss",
	"variables_scss": "internal/templates/themes/scss/styled/variables.scss",
	"hero":           "internal/templates/themes/scss/styled/Hero.svelte",
	"footer":         "internal/templates/themes/scss/styled/Footer.svelte",
}

// SveltinSCSSLibUnstyledFS is the map for the unstyled templates file whe using scss/sass.
var SveltinSCSSLibUnstyledFS = SveltinFSItem{
	"layout":         "internal/templates/themes/scss/unstyled/layout.svelte.gotxt",
	"app_css":        "internal/templates/themes/scss/unstyled/app.scss",
	"variables_scss": "internal/templates/themes/scss/unstyled/variables.scss",
	"hero":           "internal/templates/themes/scss/unstyled/Hero.svelte",
}

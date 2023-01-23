/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

// Patterns used by MigrationRule
const (
	semVersion = "semversion"
	// used to trigger the config/defaults.js.ts and config/website.js.ts files migrations
	sveltinjson = "sveltinjson"
	// used to trigger the src/sveltin.d.ts file migration
	sveltindts = "sveltin_d_ts"
	// used to trigger the .env.production file migration
	svelteKitBuildFolder  = "sveltekit-build-folder"
	svelteKitBuildComment = "sveltekit-build-comment"
	sitemap               = "sitemap"
	// used to trigger the svelte.config.js file migration
	prerenderConst   = "prerender-const"
	prerenderEnabled = "prerender-enabled"
	trailingSlash    = "trailing-slash"
	// used to trigger the themes/<theme_name>/theme.config.js file migration
	themeConfigConst  = "theme-config-const"
	themeConfigExport = "theme-config-export"
	themeNameProp     = "theme-name-prop"
	// used to trigger the mdsvex.config.js file migration
	headingsImport       = "headings-import"
	remarkExtLinks       = "remark-extlinks"
	remarkExtLinksImport = "remark-extlinks-import"
	remarkExtLinksUsage  = "remark-extlinks-usage"
	remarkSlug           = "remark-slug"
	remarkSlugImport     = "remark-slug-import"
	remarkSlugUsage      = "remakr-slug-usage"
	rehypePlugins        = "rehype-plugins"
	// used to trigger the src/lib/utils/headings.js file migration
	headingsTitleProp = "headings-js"
	// used to trigger the config/website.js.ts
	importIWebSiteSeoType = "import-iwebsite"
	iwebsiteSeoTypeUsage  = "iwebsite"
	// used to trigger the config/menu.js.ts
	importIMenuItemSeoType = "import-imenuitem"
	imenuitemSeoTypeUsage  = "imenuitem"
	// used to trigger the src/lib/utils/strings.js.ts
	icontententryTypeUsage = "content-entry"
)

var patterns = map[string]string{
	// semantic versioning regex - https://ihateregex.io/expr/semver/ .
	semVersion:             `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`,
	sveltinjson:            `/sveltin.json`,
	sveltindts:             `export type ResourceContent`,
	svelteKitBuildFolder:   `SVELTEKIT_BUILD_FOLDER`,
	svelteKitBuildComment:  `^*# The folder where adapter-static`,
	sitemap:                `^sitemap`,
	prerenderConst:         `^export const prerender`,
	prerenderEnabled:       `enabled`,
	trailingSlash:          `trailingSlash`,
	themeConfigConst:       `^const config`,
	themeConfigExport:      `^export default config`,
	themeNameProp:          `name:`,
	headingsImport:         `import headings from './src/lib/utils/headings.js`,
	remarkExtLinks:         `"remark-external-links"`,
	remarkExtLinksImport:   `^import remarkExternalLinks`,
	remarkExtLinksUsage:    `\[remarkExternalLinks`,
	remarkSlug:             "remark-slug",
	remarkSlugImport:       `^import remarkSlug`,
	remarkSlugUsage:        `remarkSlug,`,
	rehypePlugins:          `rehypePlugins:[\t\s]+\[`,
	headingsTitleProp:      `title:`,
	importIWebSiteSeoType:  `^import type { IWebSite } from '@sveltinio/seo/types';`,
	iwebsiteSeoTypeUsage:   `IWebSite`,
	importIMenuItemSeoType: `^import type { IMenuItem } from '@sveltinio/seo/types';`,
	imenuitemSeoTypeUsage:  `IMenuItem`,
	icontententryTypeUsage: `ContentEntry`,
}

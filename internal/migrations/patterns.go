/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

type migrationTriggerId int

const (
	semVersion migrationTriggerId = iota
	// config/defaults.js.ts and config/website.js.ts files migrations
	sveltinjson
	// src/sveltin.d.ts file migration
	sveltindts
	// .env.production file migration
	svelteKitBuildFolder
	svelteKitBuildComment
	sitemap
	// svelte.config.js file migration
	prerenderConst
	prerenderEnabled
	trailingSlash
	// themes/<theme_name>/theme.config.js file migration
	themeConfigConst
	themeConfigExport
	themeNameProp
	// mdsvex.config.js file migration
	headingsImport
	remarkExtLinks
	remarkExtLinksImport
	remarkExtLinksUsage
	remarkSlug
	remarkSlugImport
	remarkSlugUsage
	rehypePlugins
	rehypeSlugUsage
	// src/lib/utils/headings.js file migration
	headingsTitleProp
	// config/website.js.ts migration
	importIWebSiteSeoType
	iwebsiteSeoTypeUsage
	// config/menu.js.ts migration
	importIMenuItemSeoType
	imenuitemSeoTypeUsage
	// src/lib/utils/strings.js.ts migration
	icontententryTypeUsage
	sveltinNamespace
	capitalizeAll
	capitalizeFirstLetter
	camelToKebabCase
	toTitle
	toSlug
	// +page.[svelte|svx] migration
	iwebpagemedataImport
	jsonLdWebsiteData
	jsonLdCurrentTitle
	svelteKitPrefetch
	// vite.config.ts migration
	viteAlias
	// tsconfig.json migration
	tsPath
	// unhandled migrations where @sveltinio/* componets are used
	essentialsImport
	seoImport
	widgetsImport
)

// Patterns used by MigrationRule
var patterns = map[migrationTriggerId]string{
	// semantic versioning regex - https://ihateregex.io/expr/semver/ .
	semVersion:             `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`,
	sveltinjson:            `\/sveltin.json`,
	sveltindts:             `export type ResourceContent`,
	svelteKitBuildFolder:   `SVELTEKIT_BUILD_FOLDER`,
	svelteKitBuildComment:  `^*# The folder where adapter-static`,
	sitemap:                `^sitemap`,
	prerenderConst:         `^export const prerender`,
	prerenderEnabled:       `enabled`,
	trailingSlash:          `\btrailingSlash\b`,
	themeConfigConst:       `^const config`,
	themeConfigExport:      `^export default config`,
	themeNameProp:          `\bname:\b`,
	headingsImport:         `import headings from './src/lib/utils/headings.js`,
	remarkExtLinks:         `"remark-external-links"`,
	remarkExtLinksImport:   `^import remarkExternalLinks`,
	remarkExtLinksUsage:    `\[remarkExternalLinks`,
	remarkSlug:             "remark-slug",
	remarkSlugImport:       `^import remarkSlug`,
	remarkSlugUsage:        `remarkSlug,`,
	rehypePlugins:          `rehypePlugins:[\t\s]+\[`,
	rehypeSlugUsage:        `rehypeSlug\[`,
	headingsTitleProp:      `title:`,
	importIWebSiteSeoType:  `\{\s+IWebSite\s+\}`,
	iwebsiteSeoTypeUsage:   `\bIWebSite\b`,
	importIMenuItemSeoType: `\{\s+IMenuItem\s+\}`,
	imenuitemSeoTypeUsage:  `\bIMenuItem\b`,
	icontententryTypeUsage: `\bContentEntry\b`,
	sveltinNamespace:       `'src\/sveltin';$`,
	capitalizeAll:          `\bCapitalizeAll\b`,
	capitalizeFirstLetter:  `\bCapitalizeFirstLetter\b`,
	camelToKebabCase:       `\bCamelToKebabCase\b`,
	toTitle:                `\bToTitle\b`,
	toSlug:                 `\bToSlug\b`,
	iwebpagemedataImport:   `\bIWebPageMetadata\b`,
	jsonLdWebsiteData:      `\bwebsiteData\b`,
	jsonLdCurrentTitle:     `\bcurrentTitle\b`,
	svelteKitPrefetch:      `\bdata-sveltekit-prefetch\b`,
	viteAlias:              `^\s+(alias)`,
	tsPath:                 `^\s+"(paths)"`,
	essentialsImport:       `(.*?)'@sveltinio\/essentials';$`,
	seoImport:              `(.*?)'@sveltinio\/seo';$`,
	widgetsImport:          `(.*?)'@sveltinio\/widgets';$`,
}

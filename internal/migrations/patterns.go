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
	remarkExtLinks       = "remark-extlinks"
	remarkExtLinksImport = "remark-extlinks-import"
	remarkExtLinksUsage  = "remark-extlinks-usage"
	rehypePlugins        = "rehype-plugins"
	// used to trigger the headings.js file migration
	headingsTitleProp = "headings-js"
)

var patterns = map[string]string{
	// semantic versioning regex - https://ihateregex.io/expr/semver/ .
	semVersion:            `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`,
	sveltinjson:           `/sveltin.json`,
	svelteKitBuildFolder:  `SVELTEKIT_BUILD_FOLDER`,
	svelteKitBuildComment: `^*# The folder where adapter-static`,
	sitemap:               `^sitemap`,
	prerenderConst:        `^export const prerender`,
	prerenderEnabled:      `enabled`,
	trailingSlash:         `trailingSlash`,
	themeConfigConst:      `^const config`,
	themeConfigExport:     `^export default config`,
	themeNameProp:         `name:`,
	remarkExtLinks:        `"remark-external-links"`,
	remarkExtLinksImport:  `^import remarkExternalLinks`,
	remarkExtLinksUsage:   `\[remarkExternalLinks`,
	rehypePlugins:         `rehypePlugins:[\t\s]+\[`,
	headingsTitleProp:     `title:`,
}

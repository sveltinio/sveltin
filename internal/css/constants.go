/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package css

// CSS lib names
const (
	Bootstrap   string = "bootstrap"
	Bulma       string = "bulma"
	Scss        string = "scss"
	TailwindCSS string = "tailwindcss"
	VanillaCSS  string = "vanillacss"
)

// AvailableCSSLib is the list of the available css libs.
var AvailableCSSLib = []string{Bootstrap, Bulma, Scss, TailwindCSS, VanillaCSS}

// template file ids
const (
	PackageJSONFileId    string = "package_json"
	AppHTMLFileId        string = "app_html"
	ResetCSSFileId       string = "reset_css"
	AppCSSFileId         string = "app_css"
	VariablesFileId      string = "variables_scss"
	LayoutFileId         string = "layout"
	MDLayoutFileId       string = "md_layout"
	LayoutTSFileId       string = "layout_ts"
	ErrorFileId          string = "error"
	SvelteConfigFileId   string = "svelte_config"
	MDsveXFileId         string = "mdsvex_config"
	ViteConfigFileId     string = "vite_config"
	HeroFileId           string = "hero"
	FooterFileId         string = "footer"
	TailwindConfigFileId string = "tailwind_css_config"
	PostCSSFileId        string = "postcss"
)

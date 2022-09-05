/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package css ...
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
	PackageJSONFileID    string = "package_json"
	AppHTMLFileID        string = "app_html"
	AppCSSFileID         string = "app_css"
	VariablesFileID      string = "variables_scss"
	LayoutFileID         string = "layout"
	LayoutTSFileID       string = "layout_ts"
	ErrorFileID          string = "error"
	SvelteConfigFileID   string = "svelte_config"
	ViteConfigFileID     string = "vite_config"
	HeroFileID           string = "hero"
	FooterFileID         string = "footer"
	TailwindConfigFileID string = "tailwind_css_config"
	PostCSSFileID        string = "postcss"
)

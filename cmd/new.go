/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/spf13/cobra"
)

//=============================================================================

var (
	withCSSLib     string
	withThemeName  string
	withPortNumber string
	withGit        bool
)

// names for the available style options
const (
	StyleDefault string = "default"
	StyleNone    string = "none"
)

// names for the available CSS Lib options
const (
	Bootstrap   string = "bootstrap"
	Bulma       string = "bulma"
	Scss        string = "scss"
	TailwindCSS string = "tailwindcss"
	VanillaCSS  string = "vanillacss"
)

// names for config files
const (
	Defaults  string = "defaults"
	Externals string = "externals"
	Website   string = "website"
	Menu      string = "menu"
	DotEnv    string = "dotenv"
)

//=============================================================================

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Create new resources, pages and themes.",
	Long: `Command used to creates SvelteKit routes in your project. A routes in Sveltin is both a public page or a resource.

Examples:

sveltin new page about
sveltin new resource posts
sveltin new theme paper --css tailwindcss --npmClient pnpm`,
	Run: func(cmd *cobra.Command, args []string) {
		// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
		isValidProject(true)

		cfg.log.Important("Run 'sveltin new -h'")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

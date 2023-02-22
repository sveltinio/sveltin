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

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Create new resources and pages",
	Long: `Command used to creates SvelteKit routes in your project. A route in Sveltin is both a public page and a resource.

Examples:

sveltin new page about
sveltin new resource posts
`,
	ValidArgs:             []string{"page", "resource"},
	ArgAliases:            []string{"p", "r"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/resources"
)

//=============================================================================

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Generate static files (sitemap, rss, menu)",
	Long: resources.GetASCIIArt() + `
Command used to generate static files through its own subcommands.

Run 'sveltin generate -h' for further details.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
		isValidProject(true)

		cfg.log.Important("Run 'sveltin generate -h'")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

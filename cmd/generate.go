/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/resources"
)

//=============================================================================

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g, gen"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Command to generate static files (sitemap, rss, menu)",
	Long: resources.GetASCIIArt() + `
Used to generate static files through its own subcommands.

Run 'sveltin generate -h' for further details.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Exit if running sveltin commands from a not valid directory.
		isValidProject()

		log.Important("Run 'sveltin generate -h'")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

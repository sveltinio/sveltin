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
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g, gen"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Command to generate static files (sitemap, rss, menu)",
	Long: resources.GetAsciiArt() + `
Used to generate static files through its own subcommands.

Run 'sveltin generate -h' for further details.
`,
	Run: func(cmd *cobra.Command, args []string) {
		textLogger.Reset()
		textLogger.SetTitle("generate command called")
		textLogger.SetContent("Run 'sveltin generate -h'")
		// LOG TO STDOUT
		utils.PrettyPrinter(textLogger).Print()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

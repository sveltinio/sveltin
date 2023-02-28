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
	Short:   "Generate static files (sitemap, rss, menu)",
	Long: resources.GetASCIIArt() + `
Command used to generate static files through its own subcommands.

Run 'sveltin generate -h' for further details.
`,
	ValidArgs:             []string{"menu", "rss", "sitemap"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
}

func init() {
	generateCmd.AddGroup(&cobra.Group{ID: "generate", Title: "Available subcommands:"})
	generateCmd.SetHelpCommandGroupID("generate")
	generateCmd.SetCompletionCommandGroupID("generate")
	rootCmd.AddCommand(generateCmd)
}

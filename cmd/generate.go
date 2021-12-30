/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Command to generate static files (sitemap, rss, menu)",
	Long: resources.GetAsciiArt() + `
Used to generate static files through its own subcommands`,
	Run: func(cmd *cobra.Command, args []string) {
		printer := utils.PrinterContent{
			Title: "generate command called",
		}
		// LOG TO STDOUT
		printer.SetContent("Run 'sveltin generate -h'")
		utils.PrettyPrinter(&printer).Print()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

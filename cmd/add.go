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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add content and metadata to an existing resource.",
	Long: `Command used to add content and metadata to an existing resources through its own subcommands.

Run 'sveltin add -h' for further details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Exit if running sveltin commands from a not valid directory.
		isValidProject()

		cfg.log.Important("Run 'sveltin add -h'")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

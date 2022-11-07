/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// CliVersion is the current sveltin cli version number.
	CliVersion string = "0.10.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sveltin",
	Long:  `This is Sveltin's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(CliVersion)
	},
}

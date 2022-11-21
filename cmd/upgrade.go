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

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade to the latest Sveltin version",
	Long: resources.GetASCIIArt() + `
Command used to upgrade to the latest Sveltin version.
`,
	ValidArgs:             []string{"project"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

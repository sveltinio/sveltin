/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var update bool

//=============================================================================

var prepareCmd = &cobra.Command{
	Use:     "prepare",
	Aliases: []string{"init"},
	Short:   "Get all the dependencies from the `package.json` file",
	Long: resources.GetAsciiArt() + `
Initialize the Sveltin project getting all depencencies from the package.json file.

It wraps (npm|pnpm|yarn) install.

By default, sveltin uses pnpm as package manager. You can choose npm or yarn simply passing it via the –-package-manager flag.
	`,
	Run: RunPrepareCmd,
}

func RunPrepareCmd(cmd *cobra.Command, args []string) {
	printer := utils.PrinterContent{
		Title: "Prepare Sveltin project",
	}
	switch update {
	case true:
		// LOG TO STDOUT
		printer.SetContent("* Updating dependencies")
		utils.PrettyPrinter(&printer).Print()
		err := helpers.RunPMCommand(packageManager, "update", "", nil, false)
		common.CheckIfError(err)
	default:
		// LOG TO STDOUT
		printer.SetContent("* Getting dependencies")
		utils.PrettyPrinter(&printer).Print()
		err := helpers.RunPMCommand(packageManager, "install", "", nil, false)
		common.CheckIfError(err)
	}
}

func prepareCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&update, "update", "u", false, "Update dependencies")
}

func init() {
	rootCmd.AddCommand(prepareCmd)
	prepareCmdFlags(prepareCmd)
}

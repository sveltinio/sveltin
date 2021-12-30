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

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"serve"},
	Short:   "Run the server",
	Long: resources.GetAsciiArt() + `
It wraps svelte-kit defined commands to run the server`,
	Run: RunServerCmd,
}

func RunServerCmd(cmd *cobra.Command, args []string) {
	printer := utils.PrinterContent{
		Title: "Running Vite server",
	}
	// LOG TO STDOUT
	printer.SetContent("")
	utils.PrettyPrinter(&printer).Print()

	err := helpers.RunPMCommand(packageManager, "dev", "", nil, false)
	common.CheckIfError(err)
}

func serverCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&packageManager, "package-manager", "p", "pnpm", "The name of the your preferred package manager.")
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmdFlags(serverCmd)
}

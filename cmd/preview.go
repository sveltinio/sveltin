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

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview the production version locally",
	Long: resources.GetAsciiArt() + `
After you've built your app with sveltin build (or svelte-kit build),
you can start the production version locally with sveltin preview.

It wraps sveltekit-preview command.`,
	Run: RunPreviewCmd,
}

func RunPreviewCmd(cmd *cobra.Command, args []string) {
	printer := utils.PrinterContent{
		Title: "Preview your Sveltin project",
	}
	printer.SetContent("")
	utils.PrettyPrinter(&printer).Print()

	err := helpers.RunPMCommand(packageManager, "preview", "", nil, false)
	common.CheckIfError(err)
}

func previewCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&packageManager, "package-manager", "p", "pnpm", "The name of the your preferred package manager.")
}

func init() {
	rootCmd.AddCommand(previewCmd)
	previewCmdFlags(previewCmd)
}

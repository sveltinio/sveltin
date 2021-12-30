/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a production version of your static website",
	Long: resources.GetAsciiArt() + `
Builds a production version of your static website.

It wraps sveltekit-build command.

Ensure to edit env.production and .sveltin.toml files to reflect
your production environment
`,
	Run: RunBuildCmd,
}

func RunBuildCmd(cmd *cobra.Command, args []string) {
	printer := utils.PrinterContent{
		Title: "Building Sveltin project",
	}

	os.Setenv("VITE_PUBLIC_BASE_PATH", siteConfig.BaseURL)
	err := helpers.RunPMCommand(packageManager, "build", "", nil, false)
	common.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent("")
	utils.PrettyPrinter(&printer).Print()
}

func buildCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&packageManager, "package-manager", "p", "pnpm", "The name of the your preferred package manager.")
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmdFlags(buildCmd)
}

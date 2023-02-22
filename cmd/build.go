/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Builds a production version of your static website",
	Long: resources.GetASCIIArt() + `
Command used to build a production version of your static website.

It wraps vite build command.

Ensure to edit env.production and .sveltin.toml files to reflect
your production environment.
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(0),
	Run:                   RunBuildCmd,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument."))
		return comps, cobra.ShellCompDirectiveDefault
	},
}

// RunBuildCmd is the actual work function.
func RunBuildCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	cfg.log.Plain(markup.H1("Building the Sveltin project"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	os.Setenv("VITE_PUBLIC_BASE_PATH", cfg.prodData.BaseURL)
	err = helpers.RunPMCommand(npmClient.Name, "build", "", nil, false)
	utils.ExitIfError(err)

	cfg.log.Success("Done\n")
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

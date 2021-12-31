/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package helpers

import (
	"github.com/sveltinio/sveltin/sveltinlib/shell"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

// Execute package manager (npm, pnpm, yarn) commands
func RunPMCommand(pmName string, pmCmd string, mode string, packages []string, silentMode bool) error {
	nodePM := shell.NewNodePackageManager()
	var err error
	switch pmCmd {
	case "install":
		err = nodePM.RunInstall(pmName, pmCmd, silentMode)
	case "update":
		err = nodePM.RunUpdate(pmName, pmCmd, silentMode)
	case "dev", "build", "preview":
		err = nodePM.RunSvelteKitCommand(pmName, pmCmd, silentMode)
	case "addPackages":
		err = nodePM.RunAddPackages(pmName, pmCmd, mode, packages, silentMode)
	default:
		err = sveltinerr.NewPackageManagerCommandNotValidError()
	}
	utils.CheckIfError(err)
	return nil
}

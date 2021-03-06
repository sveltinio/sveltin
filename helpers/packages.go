/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers ...
package helpers

import (
	"github.com/sveltinio/sveltin/pkg/shell"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

// RunPMCommand executes package manager (npm, pnpm, yarn) commands.
func RunPMCommand(pmName string, pmCmd string, mode string, packages []string, silentMode bool) error {
	nodePM := shell.NewNodePackageManager()
	var err error
	switch pmCmd {
	case "install":
		err = nodePM.RunInstall(pmName, pmCmd, silentMode)
	case "update":
		err = nodePM.RunUpdate(pmName, pmCmd, silentMode)
	case "dev", "build", "preview", "prepare":
		err = nodePM.RunSvelteKitCommand(pmName, pmCmd, silentMode)
	case "addPackages":
		err = nodePM.RunAddPackages(pmName, pmCmd, mode, packages, silentMode)
	default:
		err = sveltinerr.NewPackageManagerCommandNotValidError()
	}
	utils.ExitIfError(err)
	return nil
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package helpers

import (
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/shell"
	"github.com/sveltinio/sveltin/utils"
)

// RunPMCommand executes package manager (npm, pnpm, yarn) commands.
func RunPMCommand(pmName string, pmCmd string, mode string, packages []string, silentMode bool) error {
	nodePM := shell.NewNPMClient()
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
		err = sveltinerr.NewNPMClientCommandNotValidError()
	}
	utils.ExitIfError(err)
	return nil
}

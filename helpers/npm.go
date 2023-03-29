/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package helpers

import (
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/npmclient"
)

// RunNPMCommand executes package manager (npm, pnpm, yarn) commands.
func RunNPMCommand(pmName string, pmCmd string, mode string, packages []string) error {
	switch pmCmd {
	case npmclient.InstallCmd:
		return npmclient.RunInstall(pmName, pmCmd)
	case npmclient.UpdateCmd:
		return npmclient.RunUpdate(pmName, pmCmd)
	case npmclient.AddCmd:
		return npmclient.RunAddPackages(pmName, pmCmd, mode, packages)
	case npmclient.DevCmd, npmclient.BuildCmd, npmclient.PreviewCmd, npmclient.PrepareCmd:
		return npmclient.RunSvelteKitCommand(pmName, pmCmd)
	default:
		return sveltinerr.NewNPMClientCommandNotValidError()
	}
}

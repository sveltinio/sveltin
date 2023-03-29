/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package npmclient defines the package manager client, the parser, the writer and utility functions for the package.json file.
package npmclient

import (
	"fmt"
	"strings"

	"github.com/bitfield/script"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

const (
	AddCmd     = "add"
	BuildCmd   = "build"
	DevCmd     = "dev"
	InstallCmd = "install"
	PrepareCmd = "prepare"
	PreviewCmd = "preview"
	UpdateCmd  = "update"
	UpgradeCmd = "upgrade"
)

// RunInstall execute the relative npmClient install command.
func RunInstall(npmClientName string, npmCommand string) error {
	if npmClientName == "" || npmCommand == "" {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	setNodeForceColor_256()

	cmdString := strings.Join([]string{npmClientName, npmCommand}, " ")
	if _, err := script.Exec(cmdString).Stdout(); err != nil {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	return nil
}

// RunUpdate execute the relative npmClient update command.
func RunUpdate(npmClientName string, npmCommand string) error {
	if npmClientName == "" || npmCommand == "" {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	setNodeForceColor_256()

	var npmCmd string
	if npmClientName == Yarn {
		npmCmd = UpgradeCmd
	} else {
		npmCmd = npmCommand
	}

	cmdString := strings.Join([]string{npmClientName, npmCmd}, " ")
	if _, err := script.Exec(cmdString).Stdout(); err != nil {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	return nil
}

// RunAddPackages execute the relative npmClient install|add package command.
func RunAddPackages(npmClientName string, npmCommand string, mode string, packages []string) error {
	if npmClientName == "" || npmCommand == "" || mode == "" || packages == nil {
		return sveltinerr.NewExecSystemCommandError(npmClientName, "")
	}

	setNodeForceColor_256()

	var npmCmd string
	if npmClientName == Npm {
		npmCmd = fmt.Sprintf("%s %s", InstallCmd, mode)
	} else {
		npmCmd = fmt.Sprintf("%s %s", AddCmd, mode)
	}

	fmt.Printf("  * adding %s as %s dependency", strings.Join(packages, ", "), mode)

	cmdString := strings.Join([]string{npmClientName, npmCmd}, " ")
	cmdTplString := strings.Join([]string{cmdString, "{{.}}"}, " ")
	pipe := script.Slice(packages).ExecForEach(cmdTplString)
	if _, err := pipe.Stdout(); err != nil {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	fmt.Println("\n✔ Done")
	return nil
}

func RunSvelteKitCommand(npmClientName string, npmCommand string) (err error) {
	if npmClientName == "" || npmCommand == "" {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	setNodeForceColor_256()

	var npmCmd string
	if npmClientName == Pnpm {
		npmCmd = npmCommand
	} else {
		npmCmd = fmt.Sprintf("run %s", npmCommand)
	}

	cmdString := strings.Join([]string{npmClientName, npmCmd}, " ")
	if _, err := script.Exec(cmdString).Stdout(); err != nil {
		return sveltinerr.NewExecSystemCommandError(npmClientName, npmCommand)
	}

	return nil
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package shell

import (
	"context"
	"fmt"

	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

// NPMClient is a Shell implementation used to interact with a npnClient.
type NPMClient struct {
	shell Shell
}

// NewNPMClient returns a pointer to a NodePackageManager struct.
func NewNPMClient() *NPMClient {
	return &NPMClient{
		shell: &LocalShell{},
	}
}

// GetShell returns a Shell.
func (s *NPMClient) GetShell() Shell {
	return s.shell
}

// RunInstall execute the relative npmClient install command.
func (s *NPMClient) RunInstall(pmName string, operation string, silentMode bool) error {
	if pmName == "" || operation == "" {
		return sveltinerr.NewExecSystemCommandError(pmName, operation)
	}
	pmCmd := operation
	err := s.GetShell().Execute(pmName, pmCmd, silentMode)
	if err != nil {
		return sveltinerr.NewExecSystemCommandError(pmName, pmCmd)
	}

	return nil
}

// RunUpdate execute the relative npmClient update command.
func (s *NPMClient) RunUpdate(pmName string, operation string, silentMode bool) error {
	if pmName == "" || operation == "" {
		return sveltinerr.NewExecSystemCommandError(pmName, operation)
	}

	var pmCmd string
	if pmName == "yarn" {
		pmCmd = "upgrade"
	} else {
		pmCmd = operation
	}

	err := s.GetShell().Execute(pmName, pmCmd, silentMode)
	if err != nil {
		return sveltinerr.NewExecSystemCommandError(pmName, pmCmd)
	}

	return nil
}

// RunSvelteKitCommand execute the relative npmClient sveltekit script command as defined on the package.json file.
func (s *NPMClient) RunSvelteKitCommand(pmName string, operation string, silentMode bool) (err error) {
	if pmName == "" || operation == "" {
		return sveltinerr.NewExecSystemCommandError(pmName, operation)
	}

	var pmCmd string
	if pmName == "pnpm" {
		pmCmd = operation
	} else {
		pmCmd = fmt.Sprintf("run %s", operation)
	}

	err = s.GetShell().Execute(pmName, pmCmd, silentMode)
	if err != nil {
		return sveltinerr.NewExecSystemCommandError(pmName, pmCmd)
	}

	return nil
}

// RunAddPackages execute the relative npmClient install|add package command.
func (s *NPMClient) RunAddPackages(pmName string, operation string, mode string, packages []string, silentMode bool) error {
	if pmName == "" || operation == "" || mode == "" || packages == nil {
		return sveltinerr.NewExecSystemCommandError(pmName, "")
	}

	var pmCmd string
	if pmName == "npm" {
		pmCmd = fmt.Sprintf("install %s", mode)
	} else {
		pmCmd = fmt.Sprintf("add %s", mode)
	}

	for _, p := range packages {
		fmt.Printf("  * %s\n", p)
		output, err := s.GetShell().BackgroundExecute(context.Background(), pmName, pmCmd, p)
		if err != nil {
			return sveltinerr.NewExecSystemCommandErrorWithMsg(err)
		}

		if !silentMode {
			fmt.Println(string(output))
		}
	}
	fmt.Println("\n✔ Done")
	return nil
}

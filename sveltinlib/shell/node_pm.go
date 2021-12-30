/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package shell

import (
	"context"
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/common"
)

type NodePackageManager struct {
	shell Shell
}

func NewNodePackageManager() *NodePackageManager {
	return &NodePackageManager{
		shell: &LocalShell{},
	}
}

func (n *NodePackageManager) GetShell() Shell {
	return n.shell
}

func (n *NodePackageManager) RunInstall(pmName string, operation string, silentMode bool) (err error) {
	if pmName == "" || operation == "" {
		return common.NewExecSystemCommandError()
	}
	pmCmd := operation
	err = n.GetShell().Execute(pmName, pmCmd, silentMode)
	if err != nil {
		return common.NewExecSystemCommandError()
	}

	return nil
}

func (n *NodePackageManager) RunUpdate(pmName string, operation string, silentMode bool) (err error) {
	if pmName == "" || operation == "" {
		return common.NewExecSystemCommandError()
	}

	var pmCmd string
	if pmName == "yarn" {
		pmCmd = "upgrade"
	} else {
		pmCmd = operation
	}

	err = n.GetShell().Execute(pmName, pmCmd, silentMode)
	if err != nil {
		return common.NewExecSystemCommandError()
	}

	return nil
}

func (n *NodePackageManager) RunSvelteKitCommand(pmName string, operation string, silentMode bool) (err error) {
	if pmName == "" || operation == "" {
		return common.NewExecSystemCommandError()
	}

	var pmCmd string
	if pmName == "pnpm" {
		pmCmd = operation
	} else {
		pmCmd = fmt.Sprintf("run %s", operation)
	}

	err = n.GetShell().Execute(pmName, pmCmd, silentMode)
	if err != nil {
		return common.NewExecSystemCommandError()
	}

	return nil
}

func (n *NodePackageManager) RunAddPackages(pmName string, operation string, mode string, packages []string, silentMode bool) error {
	if pmName == "" || operation == "" || mode == "" || packages == nil {
		return common.NewExecSystemCommandError()
	}

	var pmCmd string
	if pmName == "npm" {
		pmCmd = fmt.Sprintf("install %s", mode)
	} else {
		pmCmd = fmt.Sprintf("add %s", mode)
	}

	for _, p := range packages {
		jww.FEEDBACK.Printf("  * %s\n", p)
		output, err := n.GetShell().BackgroundExecute(context.Background(), pmName, pmCmd, p)
		if err != nil {
			return common.NewExecSystemCommandErrorWithMsg(err)
		}

		if !silentMode {
			jww.FEEDBACK.Println(string(output))
		}
	}
	jww.FEEDBACK.Println("\n✔ Done")
	return nil
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package shell

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"

	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

// LocalShell is a Shell implementation.
type LocalShell struct {
}

// Execute runs command on the local system.
func (s *LocalShell) Execute(cmdName string, cmdOptions string, silentMode bool) error {
	var cmd *exec.Cmd

	args := strings.Split(cmdOptions, " ")
	cmd = exec.Command(cmdName, args[0:]...)
	cmd.Stderr = os.Stderr

	if !silentMode {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Start()
	if err != nil {
		return sveltinerr.NewExecSystemCommandErrorWithMsg(err)
	}

	return cmd.Wait()
}

// BackgroundExecute runs an action on the npm client in background.
func (s *LocalShell) BackgroundExecute(ctx context.Context, cmdName string, cmdOptions string, packageList string) ([]byte, error) {
	args := strings.Split(cmdOptions, " ")
	if len(args) != 2 {
		err := errors.New("invalid number of arguments")
		return nil, sveltinerr.NewExecSystemCommandErrorWithMsg(err)
	}
	wrapperCmd := exec.CommandContext(ctx, cmdName, args[0], args[1], packageList)
	output, error := wrapperCmd.CombinedOutput()
	return output, error
}

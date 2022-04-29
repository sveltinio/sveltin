/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package shell ...
package shell

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// LocalShell is a Shell implementation.
type LocalShell struct {
}

// Execute runs an action on the npm client.
func (s *LocalShell) Execute(cmdName string, cmdOptions string, silentMode bool) error {
	var cmd *exec.Cmd

	args := strings.Split(cmdOptions, " ")
	if len(args) < 1 || len(args) > 2 {
		return sveltinerr.NewNumOfArgsNotValidError()
	}

	switch len(args) {
	case 1:
		cmd = exec.Command(cmdName, args[0])
	case 2:
		cmd = exec.Command(cmdName, args[0], args[1])
	default:
		err := errors.New("invalid number of arguments")
		return sveltinerr.NewNumOfArgsNotValidErrorWithMessage(err)
	}

	if !silentMode {
		cmd.Stdout = os.Stdout
	}

	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// BackgroundExecute runs an action on the npm client in background.
func (s *LocalShell) BackgroundExecute(ctx context.Context, pmName string, pmCmd string, packageList string) ([]byte, error) {
	args := strings.Split(pmCmd, " ")
	if len(args) != 2 {
		err := errors.New("invalid number of arguments")
		return nil, sveltinerr.NewExecSystemCommandErrorWithMsg(err)
	}
	wrapperCmd := exec.CommandContext(ctx, pmName, args[0], args[1], packageList)
	output, error := wrapperCmd.CombinedOutput()
	return output, error
}

/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package common ...
package common

import (
	"errors"
	"strconv"

	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// Contains returns true if an element is in a slice.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// CheckMinMaxArgs returns an error if the number of args is not within the expected range.
func CheckMinMaxArgs(items []string, min int, max int) error {
	switch numOfArgs := len(items); {
	case numOfArgs < min:
		errA := errors.New(`This command expects at least ` + strconv.Itoa(min) + ` argument.
Please check the help: sveltin [command] -h`)
		return sveltinerr.NewNumOfArgsNotValidErrorWithMessage(errA)
	case (numOfArgs >= min) && (numOfArgs <= max):
		return nil
	case numOfArgs > max:
		errA := errors.New(`This command expects maximum ` + strconv.Itoa(max) + ` arguments.
Please check the help: sveltin [command] -h`)
		return sveltinerr.NewNumOfArgsNotValidErrorWithMessage(errA)
	default:
		errA := errors.New("")
		return sveltinerr.NewDefaultError(errA)
	}
}

// CheckMaxArgs returns an error if there are more than N args.
func CheckMaxArgs(items []string, max int) error {
	var errorMsg string
	if max == 0 {
		errorMsg = "This command expects no arguments. Please check the help: sveltin [command] -h"
	} else {
		errorMsg = `This command expects maximum ` + strconv.Itoa(max) + ` arguments.
Please check the help: sveltin [command] -h`
	}

	switch numOfArgs := len(items); {
	case (numOfArgs >= 1) && (numOfArgs <= max):
		return nil
	case numOfArgs > max:
		errA := errors.New(errorMsg)
		return sveltinerr.NewNumOfArgsNotValidErrorWithMessage(errA)
	default:
		errA := errors.New(errorMsg)
		return sveltinerr.NewDefaultError(errA)
	}
}

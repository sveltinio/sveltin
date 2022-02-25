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
	"os"

	"github.com/manifoldco/promptui"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// PromptGetInput is used to prompt an input the user and retrieve the value.
func PromptGetInput(pc config.PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			errA := errors.New(pc.ErrorMsg)
			return sveltinerr.NewDefaultError(errA)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		jww.CRITICAL.Fatalf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

// PromptGetSelect is used to prompt a list of available options to the user and retrieve the selection.
func PromptGetSelect(items []string, pc config.PromptContent) string {
	var result string
	var err error

	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("your selection is invalid")
		}
		return nil
	}

	prompt := promptui.SelectWithAdd{
		Label:    pc.Label,
		Items:    items,
		Validate: validate,
	}

	_, result, err = prompt.Run()

	if err != nil {
		jww.CRITICAL.Fatalf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

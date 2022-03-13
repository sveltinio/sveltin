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

// PromptConfirm is used to ask for a yes or no ([Y/N]) question.
func PromptConfirm(label string) string {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		jww.CRITICAL.Fatalf("Exit %v\n", err)
		os.Exit(1)
	}
	return result
}

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
func PromptGetSelect(pc config.PromptContent, items interface{}, withTemplates bool) string {
	prompt := promptui.Select{
		Label: pc.Label,
	}
	if withTemplates {
		prompt.Templates = &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\033[0;34m\u25B6 {{ .Name | cyan }}",
			Inactive: "\033[0;37m\u2734 {{ .Name | white }}",
			Selected: "\033[0;32m\u2714 {{ .Name | white | faint}}",
		}
	}

	switch v := items.(type) {
	case []string:
		elements := []string{}
		elements = append(elements, v...)
		prompt.Items = elements
		prompt.Size = len(elements)

		i, _, err := prompt.Run()
		if err != nil {
			jww.CRITICAL.Fatalf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		return elements[i]
	case []config.PromptObject:
		elements := []config.PromptObject{}
		elements = append(elements, v...)
		prompt.Items = elements
		prompt.Size = len(elements)

		i, _, err := prompt.Run()
		if err != nil {
			jww.CRITICAL.Fatalf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		return elements[i].ID
	default:
		return ""
	}
}

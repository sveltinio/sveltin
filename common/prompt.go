/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package common ...
package common

import (
	"errors"
	"os"

	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
)

/**
CREDITS TO:
- https://github.com/manifoldco/promptui/issues/49#issuecomment-573814976
- https://github.com/manifoldco/promptui/issues/49#issuecomment-1012640880
**/

type bellSkipper struct{}

// Write implements an io.WriterCloser over os.Stderr, but it skips the terminal
// bell character.
func (bs *bellSkipper) Write(b []byte) (int, error) {
	//const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == readline.CharBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

// Close implements an io.WriterCloser over os.Stderr.
func (bs *bellSkipper) Close() error {
	return os.Stderr.Close()
}

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
func PromptGetInput(pc config.PromptContent, validate func(string) error, defaultValue string) (string, error) {
	defaultInputValidator := func(input string) error {
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
	}
	if validate != nil {
		prompt.Validate = validate
	} else {
		prompt.Validate = defaultInputValidator
	}

	if defaultValue != "" {
		prompt.Default = defaultValue
	}

	result, err := prompt.Run()
	if err != nil {
		errA := errors.New(pc.ErrorMsg)
		return "", sveltinerr.NewDefaultError(errA)
	}
	return result, nil
}

// PromptGetSelect is used to prompt a list of available options to the user and retrieve the selection.
func PromptGetSelect(pc config.PromptContent, items interface{}, withTemplates bool) (string, error) {
	prompt := promptui.Select{
		Label:  pc.Label,
		Stdout: &bellSkipper{},
	}
	if withTemplates {
		prompt.Templates = &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\033[0;34m\u25B8 {{ .Name | cyan }}",         // ▸
			Inactive: "\033[0;37m\u25B8 {{ .Name | white }}",        // ▸
			Selected: "\033[0;32m\u2714 {{ .Name | white | faint}}", // ✔
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
			errA := errors.New(pc.ErrorMsg)
			return "", sveltinerr.NewDefaultError(errA)
		}
		return elements[i], nil
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
		return elements[i].ID, nil
	default:
		errA := errors.New(pc.ErrorMsg)
		return "", sveltinerr.NewDefaultError(errA)

	}
}

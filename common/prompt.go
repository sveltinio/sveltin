/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package common

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/config"
)

func PromptGetInput(pc config.PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			errA := errors.New(pc.ErrorMsg)
			return NewDefaultError(errA)
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

func PromptGetSelect(items []string, pc config.PromptContent) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.Label,
			Items: items,
		}

		index, result, err = prompt.Run()
		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		jww.CRITICAL.Fatalf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

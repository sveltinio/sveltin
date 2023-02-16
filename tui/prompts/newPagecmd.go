package prompts

import (
	"errors"

	"github.com/charmbracelet/bubbles/list"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/common"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

const (
	// Svelte set svelte as the language used to scaffold a new page
	Svelte string = "svelte"
	// Markdown set markdown as the language used to scaffold a new page
	Markdown string = "markdown"
)

//=============================================================================

// AskPageNameHandler if no flag passed, prompts the user to set the page name.
func AskPageNameHandler(inputs []string) (string, error) {
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		pageNamePromptContent := &input.Config{
			Placeholder: "What's the page name?",
			ErrorMsg:    "Please, provide a name for the page.",
		}
		name, err := input.Run(pageNamePromptContent)
		if err != nil {
			return "", err
		}
		return utils.ToSlug(name), nil
	case numOfArgs == 1:
		name = inputs[0]
		return utils.ToSlug(name), nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

// SelectPageLanguageHandler if not flag passed, prompts the user to select the page type when creating new page.
func SelectPageLanguageHandler(name string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: Svelte, Desc: "Svelte"},
		choose.Item{Name: Markdown, Desc: "Markdown (mdsvex)"},
	}

	switch nameLenght := len(name); {
	case nameLenght == 0:
		pagePromptContent := &choose.Config{
			Title:    "How do you want to compose your page?",
			ErrorMsg: "Please, provide a page language",
		}
		result, err := choose.Run(pagePromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := choose.GetItemsKeys(entries)
		if !common.Contains(valid, name) {
			return "", sveltinerr.NewPageTypeNotValidError()
		}
		return name, nil
	default:
		return "", sveltinerr.NewPageTypeNotValidError()
	}
}

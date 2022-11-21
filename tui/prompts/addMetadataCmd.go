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

// AskMetadataNameHandler if not value, prompts the user to set the metadata name.
func AskMetadataNameHandler(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		metadataNamePromptContent := &input.Config{
			Placeholder: "What's the metadata name?",
			ErrorMsg:    "Please, provide a name for the metadata.",
		}

		result, err := input.Run(metadataNamePromptContent)
		if err != nil {
			return "", err
		}

		return utils.ToSlug(result), nil
	case numOfArgs == 1:
		return utils.ToSlug(inputs[0]), nil
	default:
		err := errors.New("something went wrong: name not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

// SelectMetadataTypeHandler if not flag passed, prompts the user to select the metadata type.
func SelectMetadataTypeHandler(mdTypeFlag string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: "single", Desc: "(1:1) One-to-One"},
		choose.Item{Name: "list", Desc: "(1:m) One-to-Many"},
	}

	switch nameLenght := len(mdTypeFlag); {
	case nameLenght == 0:
		metadataTypePromptContent := &choose.Config{
			Title:    "Which relationship between your content and the metadata?",
			ErrorMsg: "Please, provide a metadata type.",
		}
		result, err := choose.Run(metadataTypePromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := choose.GetItemsKeys(entries)
		if !common.Contains(valid, mdTypeFlag) {
			return "", sveltinerr.NewMetadataTypeNotValidError()
		}
		return mdTypeFlag, nil
	default:
		return "", sveltinerr.NewMetadataTypeNotValidError()
	}
}

package prompts

import (
	"errors"

	"github.com/sveltinio/prompti/input"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

// AskResourceNameHandler if not value, prompts the user to set the resource name.
func AskResourceNameHandler(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		resourceNamePromptContent := &input.Config{
			Placeholder: "What's the resource name? (e.g. posts, portfolio ...)",
			ErrorMsg:    "Please, provide a name for the resource.",
		}
		result, err := input.Run(resourceNamePromptContent)
		if err != nil {
			return "", err
		}
		return utils.ToSlug(result), nil
	case numOfArgs == 1:
		return utils.ToSlug(inputs[0]), nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

package prompts

import (
	"errors"

	"github.com/sveltinio/prompti/input"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

// AskThemeName prompts the user to set the theme name.
func AskThemeName(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		themeNamePromptContent := &input.Config{
			Placeholder: "What's the theme name?",
			ErrorMsg:    "Please, provide a name for the theme.",
		}
		result, err := input.Run(themeNamePromptContent)
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

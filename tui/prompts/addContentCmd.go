package prompts

import (
	"errors"

	"github.com/sveltinio/prompti/input"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

const (
	// Blank represents the fontmatter-only template id used when generating the content file.
	Blank string = "blank"
	// Sample represents the sample-content template id used when generating the content file.
	Sample string = "sample"
)

//=============================================================================

// AskContentNameHandler if not value, prompts the user to set the content name.
func AskContentNameHandler(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		contentNamePromptContent := &input.Config{
			Placeholder: "What's the content title? (it will be the slug to the page)",
			ErrorMsg:    "Please, provide a title for the content.",
		}
		contentName, err := input.Run(contentNamePromptContent)
		if err != nil {
			return "", err
		}
		return utils.ToSlug(contentName), nil
	case numOfArgs == 1:
		return utils.ToSlug(inputs[0]), nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

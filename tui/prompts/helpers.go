package prompts

import (
	"github.com/spf13/afero"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

// SelectResourceHandler if not flag passed, prompts the user to select the resource from the available ones.
func SelectResourceHandler(fs afero.Fs, mdResourceFlag string, s *config.SveltinSettings) (string, error) {
	availableResources := helpers.GetAllResources(fs, s.GetContentPath())
	resourcesList := choose.ToListItem(availableResources)

	switch nameLenght := len(mdResourceFlag); {
	case nameLenght == 0:
		resourcePromptContent := &choose.Config{
			Title:    "Which existing resource?",
			ErrorMsg: "Please, provide an existing resource name.",
		}

		//result, err := common.PromptGetSelect(resourcePromptContent, availableResources, false)
		result, err := choose.Run(resourcePromptContent, resourcesList)
		if err != nil {
			return "", err
		}
		return utils.ToSlug(result), nil
	case nameLenght != 0:
		if !common.Contains(availableResources, mdResourceFlag) {
			return "", sveltinerr.NewResourceNotFoundError()
		}
		return utils.ToSlug(mdResourceFlag), nil
	default:
		return "", sveltinerr.NewResourceNotFoundError()
	}
}

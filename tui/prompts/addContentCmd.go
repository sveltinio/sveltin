package prompts

import (
	"errors"
	"path"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

const (
	// Blank represents the fontmatter-only template id used when generating the content file.
	Blank string = "blank"
	// Sample represents the sample-content template id used when generating the content file.
	Sample string = "sample"
)

//=============================================================================

// AskContentNameHandler if not value, prompts the user to set the content name.
func AskContentNameHandler(fs afero.Fs, inputs []string, isSample bool, s *config.SveltinSettings) (*tpltypes.ContentData, error) {
	contentType := Blank
	if isSample {
		contentType = Sample
	}

	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		contentNamePromptContent := &input.Config{
			Placeholder: "What's the content title? (it will be the slug to the page)",
			ErrorMsg:    "Please, provide a title for the content.",
		}
		contentName, err := input.Run(contentNamePromptContent)
		if err != nil {
			return nil, err
		}

		contentResource, err := selectFromResourceList(fs, s)
		if err != nil {
			return nil, err
		}

		return &tpltypes.ContentData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	case numOfArgs == 1:
		name := inputs[0]
		contentResource, contentName := path.Split(name)
		contentResource = strings.ReplaceAll(contentResource, "/", "")

		err := helpers.ResourceExists(fs, contentResource, s)
		if err != nil {
			return nil, err
		}

		return &tpltypes.ContentData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

func selectFromResourceList(fs afero.Fs, s *config.SveltinSettings) (string, error) {
	availableResources := helpers.GetAllResources(fs, s.GetContentPath())

	entries := choose.ToListItem(availableResources)
	resourcePromptContent := &choose.Config{
		Title:    "Which existing resource?",
		ErrorMsg: "Please, provide an existing resource name.",
	}
	result, err := choose.Run(resourcePromptContent, entries)
	if err != nil {
		return "", err
	}
	return result, nil
}

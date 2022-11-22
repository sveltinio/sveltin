package prompts

import (
	"errors"

	"github.com/charmbracelet/bubbles/list"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/common"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/tpltypes"
)

// names for the available CSS Lib options
const (
	Bootstrap   string = "bootstrap"
	Bulma       string = "bulma"
	Scss        string = "scss"
	TailwindCSS string = "tailwindcss"
	VanillaCSS  string = "vanillacss"
)

//=============================================================================

// AskProjectNameHandler if no value, prompts the user to set the project name.
func AskProjectNameHandler(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		projectNamePromptConfig := &input.Config{
			Message:     "What's your project name?",
			Placeholder: "Please, provide a name for your project",
			ErrorMsg:    "Project name is mandatory",
		}
		result, err := input.Run(projectNamePromptConfig)
		if err != nil {
			return "", err
		}
		return result, nil
	case numOfArgs == 1:
		return inputs[0], nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

// SelectCSSLibHandler if no flag passed, prompts the user to select the CSS lib to be used with the project.
func SelectCSSLibHandler(cssLibName string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: Bootstrap, Desc: "Bootstrap"},
		choose.Item{Name: Bulma, Desc: "Bulma"},
		choose.Item{Name: VanillaCSS, Desc: "Plain CSS"},
		choose.Item{Name: Scss, Desc: "Scss/Sass"},
		choose.Item{Name: TailwindCSS, Desc: "Tailwind CSS"},
	}

	switch nameLenght := len(cssLibName); {
	case nameLenght == 0:
		cssPromptContent := &choose.Config{
			Title:    "Which CSS lib?",
			ErrorMsg: "Please, select the CSS Lib.",
		}

		result, err := choose.Run(cssPromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := choose.GetItemsKeys(entries)
		if !common.Contains(valid, cssLibName) {
			return "", sveltinerr.NewOptionNotValidError(cssLibName, valid)
		}
		return cssLibName, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

// SelectThemeHandler if no flag passed, prompts a list of available themes (blank, sveltin)
func SelectThemeHandler(themeFlag string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: tpltypes.BlankTheme, Desc: "Create a new theme"},
		choose.Item{Name: tpltypes.SveltinTheme, Desc: "Sveltin default theme"},
	}
	switch themeFlagLenght := len(themeFlag); {
	case themeFlagLenght == 0:
		themePromptContent := &choose.Config{
			Title:    "Which theme template?",
			ErrorMsg: "Please, select the theme option.",
		}

		result, err := choose.Run(themePromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case themeFlagLenght != 0:
		valid := choose.GetItemsKeys(entries)
		if !common.Contains(valid, themeFlag) {
			return "", sveltinerr.NewOptionNotValidError(themeFlag, valid)
		}
		return themeFlag, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

// SelectNPMClientHandler if no flag passed, prompts a list of installed npm client and ask the user to select one.
func SelectNPMClientHandler(items []string, npmClientFlagValue string) (string, error) {
	if len(items) == 0 {
		err := errors.New("it seems there is no package manager installed on your machine. We cannot proceed now")
		return "", sveltinerr.NewNPMClientNotFoundError(err)
	}

	entries := choose.ToListItem(items)

	switch nameLenght := len(npmClientFlagValue); {
	case nameLenght == 0:
		if len(items) == 1 {
			return items[0], nil
		}
		pmPromptContent := &choose.Config{
			Title:    "Which package manager?",
			ErrorMsg: "Please, select the package manager.",
		}

		result, err := choose.Run(pmPromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		if !common.Contains(items, npmClientFlagValue) {
			return "", sveltinerr.NewOptionNotValidError(npmClientFlagValue, items)
		}
		return npmClientFlagValue, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

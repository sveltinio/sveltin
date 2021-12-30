/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"embed"
	"errors"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/css"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	withCSS string
)

const (
	VANILLACSS  string = "Vanilla CSS"
	TAILWINDCSS string = "TailwindCSS"
	BULMA       string = "Bulma"
	BOOTSTRAP   string = "Bootstrap"
)

//=============================================================================

var newThemeCmd = &cobra.Command{
	Use:   "theme [name]",
	Short: "Create a new theme structure",
	Long: resources.GetAsciiArt() + `
Create a new theme (skeleton) called [name] in "themes" folder`,

	Run: RunNewThemeCmd,
}

func RunNewThemeCmd(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "New theme folder structure is going to be created",
	}

	themeName, err := promptThemeName(args)
	common.CheckIfError(err)

	cssLibName, err := promptCSSLibName(withCSS)
	common.CheckIfError(err)

	// NEW FOLDER: themes/<theme_name>
	newThemeFolder := composer.NewFolder(themeName)

	// NEW FOLDER: themes/<theme_name>/components
	logger.AppendItem("Creating components folder")
	componentsFolder := composer.NewFolder(pathMaker.GetThemeComponentsFolder())
	newThemeFolder.Add(componentsFolder)

	// NEW FOLDER: themes/<theme_name>/partials
	logger.AppendItem("Creating partials folder")
	partialsFolder := composer.NewFolder(pathMaker.GetThemePartialsFolder())
	newThemeFolder.Add(partialsFolder)

	// ADD FILE themes/<theme_name>/theme.config.js
	logger.AppendItem("Adding theme.config.js file")
	configFile := &composer.File{
		Name:       conf.GetThemeConfigFilename(),
		TemplateId: "theme_config",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(configFile)

	// ADD FILE themes/<theme_name>/README.md
	logger.AppendItem("Adding README.md file")
	readMeFile := &composer.File{
		Name:       "README.md",
		TemplateId: "readme",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(readMeFile)

	// ADD FILE themes/<theme_name>/LICENSE
	logger.AppendItem("Adding LICENSE file")
	licenseFile := &composer.File{
		Name:       "LICENSE",
		TemplateId: "license",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(licenseFile)

	// SET FOLDER STRUCTURE
	themesFolder := fsManager.GetFolder(THEMES)
	themesFolder.Add(newThemeFolder)
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(themesFolder)

	// GENERATE STRUCTURE
	sfs := factory.NewThemeArtifact(&resources.SveltinFS, AppFs)
	err = projectFolder.Create(sfs)
	common.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()

	if cssLibName != VANILLACSS {
		logger.Reset()
		logger.AppendItem("Installing Dependencies")
		printer = utils.PrinterContent{
			Title: "Setup the CSS library",
		}

		// LOG TO STDOUT
		printer.SetContent(logger.Render())
		utils.PrettyPrinter(&printer).Print()
		err = setupCSSLib(&resources.SveltinFS, AppFs, cssLibName, &conf, packageManager)
		common.CheckIfError(err)
	}
}

func themeCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&withCSS, "css", "c", "", "The name of the CSS framework to use. Possible values: css, tailwindcss, bulma, bootstrap")
	cmd.Flags().StringVarP(&packageManager, "package-manager", "p", "pnpm", "The name of the your preferred package manager.")
}

func init() {
	newCmd.AddCommand(newThemeCmd)
	themeCmdFlags(newThemeCmd)
}

//=============================================================================

func promptThemeName(inputs []string) (string, error) {
	var name string
	switch themeArg := len(inputs); {
	case themeArg < 1:
		themeNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your sveltin theme.",
			Label:    "What's the name of your theme?",
		}
		name = common.PromptGetInput(themeNamePromptContent)
		return name, nil
	case themeArg == 1:
		name = inputs[0]
		return name, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", common.NewDefaultError(err)
	}
}

func promptCSSLibName(useFlag string) (string, error) {
	var css string
	switch nameLenght := len(useFlag); {
	case nameLenght == 0:
		cssPromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your sveltin theme.",
			Label:    "What's the name of your theme?",
		}
		css = common.PromptGetSelect([]string{VANILLACSS, TAILWINDCSS, BULMA, BOOTSTRAP}, cssPromptContent)
		return css, nil
	case nameLenght != 0:
		css = useFlag
		return css, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", common.NewDefaultError(err)
	}
}

//=============================================================================

func setupCSSLib(efs *embed.FS, fs afero.Fs, name string, conf *config.SveltinConfig, packageManager string) error {
	switch name {
	case VANILLACSS:
		vanillaCSS := &css.VanillaCSS{}
		c := css.CSSLib{
			ICSSLib: vanillaCSS,
		}
		return c.Setup(efs, fs, conf, packageManager)
	case TAILWINDCSS:
		tailwind := &css.TailwindCSS{}
		c := css.CSSLib{
			ICSSLib: tailwind,
		}
		return c.Setup(efs, fs, conf, packageManager)
	case BULMA:
		bulma := &css.Bulma{}
		c := css.CSSLib{
			ICSSLib: bulma,
		}
		return c.Setup(efs, fs, conf, packageManager)
	case BOOTSTRAP:
		boostrap := &css.Bootstrap{}
		c := css.CSSLib{
			ICSSLib: boostrap,
		}
		return c.Setup(efs, fs, conf, packageManager)
	default:
		return common.NewOptionNotValidError()
	}
}

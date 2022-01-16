/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"embed"
	"errors"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/css"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	withCSSLib    string
	withThemeName string
)

const (
	VANILLACSS  string = "vanillacss"
	TAILWINDCSS string = "tailwindcss"
	BULMA       string = "bulma"
	BOOTSTRAP   string = "bootstrap"
)

const (
	DEFAULTS  string = "defaults"
	EXTERNALS string = "externals"
	WEBSITE   string = "website"
	MENU      string = "menu"
	DOTENV    string = "dotenv"
)

//=============================================================================

var newCmd = &cobra.Command{
	Use:     "new <project>",
	Aliases: []string{"create"},
	Args:    cobra.RangeArgs(0, 3),
	Short:   "Command to create projects, resources, contents, pages, metadata, theme",
	Long: resources.GetAsciiArt() + `
This command creates customized resources like blog posts, a new theme for your website, new page etc. depending on the subcommand used with it.

Examples:

sveltin new blog
sveltin new blog --css tailwindcss
sveltin new blog --css vanillacss -t myTheme
sveltin new resource posts`,
	Run: NewCmdRun,
}

func NewCmdRun(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "A new Sveltin based project will be created",
	}

	projectName, err := promptProjectName(args)
	utils.CheckIfError(err)

	cssLibName, err := promptCSSLibName(withCSSLib)
	utils.CheckIfError(err)

	themeName := getThemeName(projectName)

	setupPackageManager()

	// Clone starter template github repository
	starterTemplate := appTemplatesMap[SVELTEKIT_STARTER]
	logger.AppendItem(`Cloning ` + starterTemplate.URL)
	utils.GitClone(&starterTemplate, pathMaker.GetProjectRoot(projectName))

	// GET FOLDER: <project_name>
	projectFolder := fsManager.GetFolder(projectName)

	// NEW FOLDER: config
	logger.AppendItem("Creating 'config' folder")
	configFolder := composer.NewFolder(CONFIG)
	projectFolder.Add(configFolder)

	// NEW FILE: config/<filename>.js
	logger.AppendItem("Adding config files")
	for _, elem := range []string{DEFAULTS, EXTERNALS, WEBSITE, MENU} {
		f := fsManager.NewConfigFile(projectName, elem, CLI_VERSION)
		configFolder.Add(f)
	}

	// NEW FOLDER: content
	logger.AppendItem("Creating 'content' folder")
	contentFolder := composer.NewFolder(CONTENT)
	projectFolder.Add(contentFolder)

	// NEW FOLDER: themes
	logger.AppendItem("Creating 'themes' folder")
	themesFolder := composer.NewFolder(THEMES)

	newThemeFolder := makeThemeStructure(themeName)
	themesFolder.Add(newThemeFolder)
	// ADD themes folder to the project
	projectFolder.Add(themesFolder)

	// NEW FILES: .env.development and env.production
	logger.AppendItem("Adding 'dotenv' files")
	for _, item := range []string{DOTENV_DEV, DOTENV_PROD} {
		f := fsManager.NewDotEnvFile(projectName, item)
		projectFolder.Add(f)
	}

	// SET FOLDER STRUCTURE
	rootFolder := fsManager.GetFolder(ROOT)
	rootFolder.Add(projectFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewProjectArtifact(&resources.SveltinFS, AppFs)
	err = rootFolder.Create(sfs)
	utils.CheckIfError(err)

	// SETUP THE CSS LIB
	logger.AppendItem("Setup the CSS Lib")
	err = setupCSSLib(&resources.SveltinFS, AppFs, cssLibName, &conf, projectName, themeName)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func newCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&npmClient, "npmClient", "n", "", "The name of the your preferred npm client.")
	cmd.Flags().StringVarP(&withThemeName, "theme", "t", "", "The name of the theme you are going to create")
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "", "The name of the CSS framework to use. Possible values: vanillacss, tailwindcss, bulma, bootstrap")
}

func init() {
	newCmdFlags(newCmd)
	rootCmd.AddCommand(newCmd)
}

//=============================================================================

func promptProjectName(inputs []string) (string, error) {
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		projectNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your website.",
			Label:    "What's the name of your website?",
		}
		name = common.PromptGetInput(projectNamePromptContent)
		return name, nil
	case numOfArgs == 1:
		name = inputs[0]
		return name, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

func promptCSSLibName(cssLibName string) (string, error) {
	var css string
	valid := []string{VANILLACSS, TAILWINDCSS, BULMA, BOOTSTRAP}
	switch nameLenght := len(cssLibName); {
	case nameLenght == 0:
		cssPromptContent := config.PromptContent{
			ErrorMsg: "Please, provide the CSS Lib name.",
			Label:    "Which CSS lib do you want to use?",
		}
		css = common.PromptGetSelect(valid, cssPromptContent)
		return css, nil
	case nameLenght != 0:
		if !common.Contains(valid, cssLibName) {
			return "", sveltinerr.NewOptionNotValidError()
		}
		css = cssLibName
		return css, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

func getThemeName(name string) string {
	if len(withThemeName) != 0 {
		return withThemeName
	} else {
		return strings.Join([]string{name, "theme"}, "_")
	}
}

func promptNPMClient(items []string) (string, error) {
	if len(items) == 0 {
		err := errors.New("it seems there is no package manager on your machine")
		return "", sveltinerr.NewDefaultError(err)
	}
	var client string
	switch nameLenght := len(npmClient); {
	case nameLenght == 0:
		if len(items) == 1 {
			client = items[0]
		} else {
			pmPromptContent := config.PromptContent{
				ErrorMsg: "Please, provide the name of the package manager.",
				Label:    "Which package manager do you want to use?",
			}
			_pm := common.PromptGetSelect(items, pmPromptContent)
			if common.Contains(items, _pm) {
				client = _pm
			} else {
				errN := errors.New("invalid selection. Valid options are " + strings.Join(items, ", "))
				return "", sveltinerr.NewDefaultError(errN)
			}
		}
		return client, nil
	case nameLenght != 0:
		if !common.Contains(items, npmClient) {
			return "", sveltinerr.NewOptionNotValidError()
		}
		client = npmClient
		return client, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func makeThemeStructure(themeName string) *composer.Folder {
	// NEW FOLDER: themes/<theme_name>
	newThemeFolder := composer.NewFolder(themeName)

	// NEW FOLDER: themes/<theme_name>/components
	componentsFolder := composer.NewFolder(pathMaker.GetThemeComponentsFolder())
	newThemeFolder.Add(componentsFolder)

	// NEW FOLDER: themes/<theme_name>/partials
	partialsFolder := composer.NewFolder(pathMaker.GetThemePartialsFolder())
	newThemeFolder.Add(partialsFolder)

	// ADD FILE themes/<theme_name>/theme.config.js
	configFile := &composer.File{
		Name:       conf.GetThemeConfigFilename(),
		TemplateId: "theme_config",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(configFile)

	// ADD FILE themes/<theme_name>/README.md
	readMeFile := &composer.File{
		Name:       "README.md",
		TemplateId: "readme",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(readMeFile)

	// ADD FILE themes/<theme_name>/LICENSE
	licenseFile := &composer.File{
		Name:       "LICENSE",
		TemplateId: "license",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(licenseFile)

	return newThemeFolder
}

/**
 * Read the settings file, if it does not exists and no -p flag,
 * prompt to select the package manager from the ones currently
 * installed on the machine and store its value as settings.
 */
func setupPackageManager() {
	if len(settings.GetNPMClient()) == 0 && len(npmClient) == 0 {
		selectedPackageManager, err := promptNPMClient(utils.GetAvailableNPMClientList())
		utils.CheckIfError(err)
		storeSelectedNPMClient(selectedPackageManager)
	}
}

func setupCSSLib(efs *embed.FS, fs afero.Fs, name string, conf *config.SveltinConfig, projectName string, themeName string) error {
	switch name {
	case VANILLACSS:
		vanillaCSS := &css.VanillaCSS{}
		c := css.CSSLib{
			ICSSLib: vanillaCSS,
		}
		return c.Setup(efs, fs, conf, projectName, themeName)
	case TAILWINDCSS:
		tailwind := &css.TailwindCSS{}
		c := css.CSSLib{
			ICSSLib: tailwind,
		}
		return c.Setup(efs, fs, conf, projectName, themeName)
	case BULMA:
		bulma := &css.Bulma{}
		c := css.CSSLib{
			ICSSLib: bulma,
		}
		return c.Setup(efs, fs, conf, projectName, themeName)
	case BOOTSTRAP:
		boostrap := &css.Bootstrap{}
		c := css.CSSLib{
			ICSSLib: boostrap,
		}
		return c.Setup(efs, fs, conf, projectName, themeName)
	default:
		return sveltinerr.NewOptionNotValidError()
	}
}

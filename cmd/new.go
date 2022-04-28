/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/css"
	"github.com/sveltinio/sveltin/sveltinlib/npmc"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	withStyle      string
	withCSSLib     string
	withThemeName  string
	withPortNumber string
)

// names for the available style options
const (
	StyleDefault string = "default"
	StyleNone    string = "none"
)

// names for the available CSS Lib options
const (
	VanillaCSS  string = "vanillacss"
	TailwindCSS string = "tailwindcss"
	Bulma       string = "bulma"
	Bootstrap   string = "bootstrap"
	Scss        string = "scss"
)

// names for config files
const (
	Defaults  string = "defaults"
	Externals string = "externals"
	Website   string = "website"
	Menu      string = "menu"
	DotEnv    string = "dotenv"
)

//=============================================================================

var newCmd = &cobra.Command{
	Use:     "new <project>",
	Aliases: []string{"create"},
	Args:    cobra.RangeArgs(0, 3),
	Short:   "Command to create projects, resources, contents, pages, metadata, theme",
	Long: resources.GetASCIIArt() + `
This command creates projects, resources (e.g. blog posts), pagee, content, metadata etc. depending on the subcommand used with it.

Examples:

sveltin new blog
sveltin new blog --css tailwindcss
sveltin new blog --css vanillacss -t myTheme
sveltin new resource posts`,
	Run: NewCmdRun,
}

// NewCmdRun is the actual work function.
func NewCmdRun(cmd *cobra.Command, args []string) {
	projectName, err := promptProjectName(args)
	utils.ExitIfError(err)

	projectStyles, err := promptProjectStyle(withStyle)
	utils.ExitIfError(err)

	cssLibName, err := promptCSSLibName(withCSSLib)
	utils.ExitIfError(err)

	themeName := getThemeName(projectName)

	npmClient := getSelectedNPMClient()
	npmClientName = npmClient.Name

	log.Plain(utils.Underline("A new Sveltin project will be created"))

	// Clone starter template github repository
	starterTemplate := appTemplatesMap[SvelteKitStarter]
	log.Info(fmt.Sprintf("Cloning the %s repos", starterTemplate.Name))
	err = utils.GitClone(starterTemplate.URL, pathMaker.GetProjectRoot(projectName))
	utils.ExitIfError(err)

	// GET FOLDER: <project_name>
	projectFolder := fsManager.GetFolder(projectName)

	log.Info("Creating the project folder structure")

	// NEW FOLDER: config
	configFolder := composer.NewFolder(Config)
	projectFolder.Add(configFolder)

	// NEW FILE: config/<filename>.js
	for _, elem := range []string{Defaults, Externals, Website, Menu} {
		f := fsManager.NewConfigFile(projectName, elem, CliVersion)
		configFolder.Add(f)
	}

	// NEW FOLDER: content
	contentFolder := composer.NewFolder(Content)
	projectFolder.Add(contentFolder)

	// GET FOLDER: src/routes folder
	routesFolder := fsManager.GetFolder(Routes)
	// NEW FILE: index.svelte
	indexFile := &composer.File{
		Name:       helpers.GetResourceRouteFilename(Index, &conf),
		TemplateID: Index,
		TemplateData: &config.TemplateData{
			ThemeName: themeName,
		},
	}
	// ADD src/routes folder to the project
	routesFolder.Add(indexFile)
	projectFolder.Add(routesFolder)

	// NEW FOLDER: themes
	themesFolder := composer.NewFolder(Themes)

	newThemeFolder := makeThemeStructure(themeName)
	themesFolder.Add(newThemeFolder)
	// ADD themes folder to the project
	projectFolder.Add(themesFolder)

	// NEW FILE: env.production
	dotEnvTplData := &config.TemplateData{
		Name:    DotEnvProd,
		BaseURL: fmt.Sprintf("http://%s.com", projectName),
	}
	f := fsManager.NewDotEnvFile(projectName, dotEnvTplData)
	projectFolder.Add(f)

	// SET FOLDER STRUCTURE
	rootFolder := fsManager.GetFolder(Root)
	rootFolder.Add(projectFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewProjectArtifact(&resources.SveltinFS, AppFs)
	err = rootFolder.Create(sfs)
	utils.ExitIfError(err)

	// SETUP THE CSS LIB
	log.Info("Setting up the CSS Lib")
	tplData := config.TemplateData{
		ProjectName: projectName,
		NPMClient:   npmClient.ToString(),
		PortNumber:  withPortNumber,
		ThemeName:   themeName,
	}
	err = setupCSSLib(&resources.SveltinFS, AppFs, cssLibName, isSveltinStyles(projectStyles), &conf, &tplData)
	utils.ExitIfError(err)

	log.Success("Done")

	// NEXT STEPS
	log.Plain(utils.Underline("Next Steps"))
	log.Plain(common.HelperTextNewProject(projectName))
}

func newCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&withStyle, "style", "s", "", "Default styles or unstyled. Possible values: default, none")
	cmd.Flags().StringVarP(&npmClientName, "npmClient", "n", "", "The name of your preferred npm client")
	cmd.Flags().StringVarP(&withThemeName, "theme", "t", "", "The name of the theme you are going to create")
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "", "The name of the CSS framework to use. Possible values: vanillacss, tailwindcss, bulma, bootstrap, scss")
	cmd.Flags().StringVarP(&withPortNumber, "port", "p", "3000", "The port to start the server on")
}

func init() {
	newCmdFlags(newCmd)
	rootCmd.AddCommand(newCmd)
}

//=============================================================================

func promptProjectName(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		projectNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your project.",
			Label:    "What's the name of your project?",
		}
		result, err := common.PromptGetInput(projectNamePromptContent, nil, "")
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

func promptProjectStyle(stylesName string) (string, error) {
	promptObjects := []config.PromptObject{
		{ID: StyleDefault, Name: "Sveltin default styles"},
		{ID: StyleNone, Name: "None"},
	}

	switch nameLenght := len(stylesName); {
	case nameLenght == 0:
		stylesPromptContent := config.PromptContent{
			ErrorMsg: "Please, provide the style name",
			Label:    "Which style for your Sveltin app?",
		}
		result, err := common.PromptGetSelect(stylesPromptContent, promptObjects, true)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := common.GetPromptObjectKeys(promptObjects)
		if !common.Contains(valid, stylesName) {
			return "", sveltinerr.NewOptionNotValidError(stylesName, valid)
		}
		return stylesName, nil
	default:
		err := fmt.Errorf("something went wrong: value not valid! You used: %s", stylesName)
		return "", sveltinerr.NewDefaultError(err)
	}
}

func promptCSSLibName(cssLibName string) (string, error) {
	promptObjects := []config.PromptObject{
		{ID: VanillaCSS, Name: "Plain CSS"},
		{ID: Scss, Name: "Scss/Sass"},
		{ID: TailwindCSS, Name: "Tailwind CSS"},
		{ID: Bulma, Name: "Bulma"},
		{ID: Bootstrap, Name: "Bootstrap"},
	}

	switch nameLenght := len(cssLibName); {
	case nameLenght == 0:
		cssPromptContent := config.PromptContent{
			ErrorMsg: "Please, provide the CSS Lib name.",
			Label:    "Which CSS lib do you want to use?",
		}
		result, err := common.PromptGetSelect(cssPromptContent, promptObjects, true)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := common.GetPromptObjectKeys(promptObjects)
		if !common.Contains(valid, cssLibName) {
			return "", sveltinerr.NewOptionNotValidError(cssLibName, valid)
		}
		return cssLibName, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

func getThemeName(name string) string {
	if len(withThemeName) != 0 {
		return withThemeName
	}
	return strings.Join([]string{name, "theme"}, "_")
}

func promptNPMClient(items []string) (string, error) {
	if len(items) == 0 {
		err := errors.New("it seems there is no package manager on your machine")
		return "", sveltinerr.NewDefaultError(err)
	}

	switch nameLenght := len(npmClientName); {
	case nameLenght == 0:
		if len(items) == 1 {
			return items[0], nil
		}
		pmPromptContent := config.PromptContent{
			ErrorMsg: "Please, provide the name of the package manager.",
			Label:    "Which package manager do you want to use?",
		}

		result, err := common.PromptGetSelect(pmPromptContent, items, false)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		if !common.Contains(items, npmClientName) {
			return "", sveltinerr.NewOptionNotValidError(npmClientName, items)
		}
		return npmClientName, nil
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
		TemplateID: "theme_config",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(configFile)

	// ADD FILE themes/<theme_name>/README.md
	readMeFile := &composer.File{
		Name:       utils.ToMDFile("readme", true),
		TemplateID: "readme",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(readMeFile)

	// ADD FILE themes/<theme_name>/LICENSE
	licenseFile := &composer.File{
		Name:       "LICENSE",
		TemplateID: "license",
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
func getSelectedNPMClient() npmc.NPMClient {
	installedNPMClients := utils.GetInstalledNPMClientList()
	npmClientNames := utils.GetNPMClientNames(installedNPMClients)
	client, err := promptNPMClient(npmClientNames)
	utils.ExitIfError(err)
	return utils.GetSelectedNPMClient(installedNPMClients, client)
}

func setupCSSLib(efs *embed.FS, fs afero.Fs, cssLibName string, isStyled bool, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	switch cssLibName {
	case VanillaCSS:
		vanillaCSS := css.NewVanillaCSS(isStyled, efs, fs, conf, tplData)
		return vanillaCSS.Setup()
	case Scss:
		scss := css.NewScss(isStyled, efs, fs, conf, tplData)
		return scss.Setup()
	case TailwindCSS:
		tailwind := css.NewTailwindCSS(isStyled, efs, fs, conf, tplData)
		return tailwind.Setup()
	case Bulma:
		bulma := css.NewBulma(isStyled, efs, fs, conf, tplData)
		return bulma.Setup()
	case Bootstrap:
		boostrap := css.NewBootstrap(isStyled, efs, fs, conf, tplData)
		return boostrap.Setup()
	default:
		return sveltinerr.NewOptionNotValidError(cssLibName, []string{"vanillacss", "tailwindcss", "bulma", "bootstrap", "scss"})
	}
}

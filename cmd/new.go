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
	withCSSLib     string
	withThemeName  string
	withPortNumber string
)

// names for the available CSS Lib options
const (
	VANILLACSS  string = "vanillacss"
	TAILWINDCSS string = "tailwindcss"
	BULMA       string = "bulma"
	BOOTSTRAP   string = "bootstrap"
	SCSS        string = "scss"
)

// names for config files
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

	cssLibName, err := promptCSSLibName(withCSSLib)
	utils.ExitIfError(err)

	themeName := getThemeName(projectName)

	npmClient := getSelectedNPMClient()
	npmClientName = npmClient.Name

	log.Plain(utils.Underline("A new Sveltin project will be created"))

	// Clone starter template github repository
	starterTemplate := appTemplatesMap[SVELTEKIT_STARTER]
	log.Info(fmt.Sprintf("Cloning the %s repos", starterTemplate.Name))
	err = utils.GitClone(&starterTemplate, pathMaker.GetProjectRoot(projectName))
	utils.ExitIfError(err)

	// GET FOLDER: <project_name>
	projectFolder := fsManager.GetFolder(projectName)

	log.Info("Creating the project folder structure")

	// NEW FOLDER: config
	configFolder := composer.NewFolder(CONFIG)
	projectFolder.Add(configFolder)

	// NEW FILE: config/<filename>.js
	for _, elem := range []string{DEFAULTS, EXTERNALS, WEBSITE, MENU} {
		f := fsManager.NewConfigFile(projectName, elem, CLI_VERSION)
		configFolder.Add(f)
	}

	// NEW FOLDER: content
	contentFolder := composer.NewFolder(CONTENT)
	projectFolder.Add(contentFolder)

	// GET FOLDER: src/routes folder
	routesFolder := fsManager.GetFolder(ROUTES)
	// NEW FILE: index.svelte
	indexFile := &composer.File{
		Name:       helpers.GetResourceRouteFilename(INDEX, &conf),
		TemplateId: INDEX,
		TemplateData: &config.TemplateData{
			ThemeName: themeName,
		},
	}
	// ADD src/routes folder to the project
	routesFolder.Add(indexFile)
	projectFolder.Add(routesFolder)

	// NEW FOLDER: themes
	themesFolder := composer.NewFolder(THEMES)

	newThemeFolder := makeThemeStructure(themeName)
	themesFolder.Add(newThemeFolder)
	// ADD themes folder to the project
	projectFolder.Add(themesFolder)

	// NEW FILE: env.production
	dotEnvTplData := &config.TemplateData{
		Name:    DOTENV_PROD,
		BaseURL: fmt.Sprintf("http://%s.com", projectName),
	}
	f := fsManager.NewDotEnvFile(projectName, dotEnvTplData)
	projectFolder.Add(f)

	// SET FOLDER STRUCTURE
	rootFolder := fsManager.GetFolder(ROOT)
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
	err = setupCSSLib(&resources.SveltinFS, AppFs, cssLibName, &conf, &tplData)
	utils.ExitIfError(err)

	log.Success("Done")

	// NEXT STEPS
	log.Plain(utils.Underline("Next Steps"))
	log.Plain(common.HelperTextNewProject(projectName))
}

func newCmdFlags(cmd *cobra.Command) {
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
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		projectNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your project.",
			Label:    "What's the name of your project?",
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
	promptObjects := []config.PromptObject{
		{Id: VANILLACSS, Name: "Plain CSS"},
		{Id: SCSS, Name: "Scss/Sass"},
		{Id: TAILWINDCSS, Name: "Tailwind CSS"},
		{Id: BULMA, Name: "Bulma"},
		{Id: BOOTSTRAP, Name: "Bootstrap"},
	}

	switch nameLenght := len(cssLibName); {
	case nameLenght == 0:
		cssPromptContent := config.PromptContent{
			ErrorMsg: "Please, provide the CSS Lib name.",
			Label:    "Which CSS lib do you want to use?",
		}
		return common.PromptGetSelect(cssPromptContent, promptObjects, true), nil
	case nameLenght != 0:
		valid := common.GetPromptObjectKeys(promptObjects)
		if !common.Contains(valid, cssLibName) {
			return "", sveltinerr.NewOptionNotValidError()
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
		} else {
			pmPromptContent := config.PromptContent{
				ErrorMsg: "Please, provide the name of the package manager.",
				Label:    "Which package manager do you want to use?",
			}
			return common.PromptGetSelect(pmPromptContent, items, false), nil
		}
	case nameLenght != 0:
		if !common.Contains(items, npmClientName) {
			errN := errors.New("invalid selection. Valid options are " + strings.Join(items, ", "))
			return "", sveltinerr.NewDefaultError(errN)
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
		TemplateId: "theme_config",
		TemplateData: &config.TemplateData{
			Name: themeName,
		},
	}
	newThemeFolder.Add(configFile)

	// ADD FILE themes/<theme_name>/README.md
	readMeFile := &composer.File{
		Name:       utils.ToMDFile("readme", true),
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
func getSelectedNPMClient() npmc.NPMClient {
	installedNPMClients := utils.GetInstalledNPMClientList()
	npmClientNames := utils.GetNPMClientNames(installedNPMClients)
	client, err := promptNPMClient(npmClientNames)
	utils.ExitIfError(err)
	return utils.GetSelectedNPMClient(installedNPMClients, client)
}

func setupCSSLib(efs *embed.FS, fs afero.Fs, cssLibName string, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	switch cssLibName {
	case VANILLACSS:
		vanillaCSS := &css.VanillaCSS{}
		c := css.CSSLib{
			ICSSLib: vanillaCSS,
		}
		return c.Setup(efs, fs, conf, tplData)
	case SCSS:
		scss := &css.Scss{}
		c := css.CSSLib{
			ICSSLib: scss,
		}
		return c.Setup(efs, fs, conf, tplData)
	case TAILWINDCSS:
		tailwind := &css.TailwindCSS{}
		c := css.CSSLib{
			ICSSLib: tailwind,
		}
		return c.Setup(efs, fs, conf, tplData)
	case BULMA:
		bulma := &css.Bulma{}
		c := css.CSSLib{
			ICSSLib: bulma,
		}
		return c.Setup(efs, fs, conf, tplData)
	case BOOTSTRAP:
		boostrap := &css.Bootstrap{}
		c := css.CSSLib{
			ICSSLib: boostrap,
		}
		return c.Setup(efs, fs, conf, tplData)
	default:
		return sveltinerr.NewOptionNotValidError()
	}
}

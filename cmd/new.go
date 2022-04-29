/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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
	"github.com/sveltinio/sveltin/sveltinlib/shell"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	withCSSLib     string
	withThemeName  string
	withPortNumber string
	withGit        bool
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
	Short:   "Command to create projects, resources, contents, pages, metadata and themes.",
	Long: resources.GetASCIIArt() + `
This command creates projects, resources (e.g. blog posts, recipes, ...), pages, content, metadata, themes etc. depending on the subcommand used with it.

Examples:

sveltin new blog
sveltin new blog --css tailwindcss
sveltin new blog --css vanillacss -t myTheme
sveltin new resource posts
sveltin new theme paper --css tailwindcss --npmClient pnpm`,
	Run: NewCmdRun,
}

// NewCmdRun is the actual work function.
func NewCmdRun(cmd *cobra.Command, args []string) {
	projectName, err := promptProjectName(args)
	utils.ExitIfError(err)

	cssLibName, err := promptCSSLibName(withCSSLib)
	utils.ExitIfError(err)

	themeSelection, err := promptThemeSelection(withThemeName)
	utils.ExitIfError(err)

	themeData, err := buildThemeData(themeSelection, withThemeName, projectName, cssLibName)
	utils.ExitIfError(err)

	npmClient := getSelectedNPMClient()
	npmClientName = npmClient.Name

	log.Plain(utils.Underline("A new Sveltin project will be created"))
	// Clone starter template github repository
	starterTemplate := appTemplatesMap[SvelteKitStarter]
	log.Info(fmt.Sprintf("Cloning the %s repos", starterTemplate.Name))
	gitClient := shell.NewGitClient()
	err = gitClient.RunGitClone(starterTemplate.URL, pathMaker.GetProjectRoot(projectName), true)
	//err = utils.GitClone(starterTemplate.URL, pathMaker.GetProjectRoot(projectName))
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
			Theme: themeData,
		},
	}
	// ADD src/routes folder to the project
	routesFolder.Add(indexFile)
	projectFolder.Add(routesFolder)

	// NEW FOLDER: themes
	themesFolder := composer.NewFolder(Themes)

	newThemeFolder := makeThemeFolderStructure(themeData)
	if newThemeFolder != nil {
		themesFolder.Add(newThemeFolder)
	}
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
		Theme:       themeData,
	}
	err = setupCSSLib(&resources.SveltinFS, AppFs, &conf, &tplData)
	utils.ExitIfError(err)

	// INITIALIZE GIT REPO
	if isInitGitRepo(withGit) {
		log.Info("Initializing empty Git repository")
		err = gitClient.RunInit(projectFolder.GetPath(), true)
		utils.ExitIfError(err)
	}

	log.Success("Done")

	// NEXT STEPS
	log.Plain(utils.Underline("Next Steps"))
	if themeData.ID != config.ExistingTheme {
		log.Plain(common.HelperTextNewProject(projectName))
	} else {
		log.Plain(common.HelperTextNewProjectWithExistingTheme(projectName))
	}

}

func newCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&npmClientName, "npmClient", "n", "", "The name of your preferred npm client")
	cmd.Flags().StringVarP(&withThemeName, "theme", "t", "", "The theme you are going to create or reuse")
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "", "The name of the CSS framework to use. Valid: vanillacss, tailwindcss, bulma, bootstrap, scss")
	cmd.Flags().StringVarP(&withPortNumber, "port", "p", "3000", "The port to start the server on")
	cmd.Flags().BoolVarP(&withGit, "git", "g", false, "Initialize an empty Git repository")
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

func promptThemeSelection(themeFlag string) (string, error) {
	switch themeFlagLenght := len(themeFlag); {
	case themeFlagLenght == 0:
		promptObjects := []config.PromptObject{
			{ID: config.BlankTheme, Name: "Create a new theme"},
			{ID: config.SveltinTheme, Name: "Sveltin default theme"},
			{ID: config.ExistingTheme, Name: "Use an existing theme"},
		}
		themePromptContent := config.PromptContent{
			ErrorMsg: "Please, select a theme option.",
			Label:    "Do you wish to create a new theme or using an existing one?",
		}
		result, err := common.PromptGetSelect(themePromptContent, promptObjects, true)
		if err != nil {
			return "", err
		}
		return result, nil
	case themeFlagLenght != 0:
		return themeFlag, nil
	default:
		err := errors.New("something went wrong: value not valid")
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
			Label:    "Which CSS lib do you want to use for your theme?",
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

func buildThemeData(input, themeFlagValue, projectName, cssLibName string) (*config.ThemeData, error) {
	switch input {
	case config.BlankTheme:
		defaultThemeName := strings.Join([]string{projectName, "theme"}, "_")
		newThemePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your theme.",
			Label:    "What's the name for your new theme?",
		}
		themeName, err := common.PromptGetInput(newThemePromptContent, nil, defaultThemeName)
		if err != nil {
			return nil, err
		}
		return &config.ThemeData{
			ID:     config.BlankTheme,
			IsNew:  true,
			Name:   themeName,
			CSSLib: cssLibName,
		}, nil
	case config.SveltinTheme:
		return &config.ThemeData{
			ID:     config.SveltinTheme,
			IsNew:  false,
			Name:   "sveltin_theme",
			CSSLib: cssLibName,
		}, nil
	case config.ExistingTheme:
		return &config.ThemeData{
			ID:     config.ExistingTheme,
			IsNew:  false,
			CSSLib: cssLibName,
		}, nil
	default:
		if utils.IsValidURL(input) {
			_, err := utils.NewGitHubURLParser(input)
			if err != nil {
				return nil, err
			}
			return &config.ThemeData{
				ID:     config.ExistingTheme,
				IsNew:  false,
				CSSLib: cssLibName,
			}, nil
		}

		return &config.ThemeData{
			ID:     config.BlankTheme,
			IsNew:  true,
			Name:   getNewThemeName(themeFlagValue, projectName),
			CSSLib: cssLibName,
		}, nil
	}
}

func getNewThemeName(value, projectName string) string {
	if len(value) != 0 {
		return value
	}
	return strings.Join([]string{projectName, "theme"}, "_")
}

func makeThemeFolderStructure(entry *config.ThemeData) *composer.Folder {
	if entry.IsNew || entry.ID == config.SveltinTheme {
		return createNewThemeLocalFolder(entry)
	}
	return nil
}

func createNewThemeLocalFolder(themeData *config.ThemeData) *composer.Folder {
	// NEW FOLDER: themes/<theme_name>
	newThemeFolder := composer.NewFolder(themeData.Name)

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
			Theme: themeData,
		},
	}
	newThemeFolder.Add(configFile)

	// ADD FILE themes/<theme_name>/README.md
	readMeFile := &composer.File{
		Name:       utils.ToMDFile("readme", true),
		TemplateID: "readme",
		TemplateData: &config.TemplateData{
			Name: themeData.Name,
		},
	}
	newThemeFolder.Add(readMeFile)

	// ADD FILE themes/<theme_name>/LICENSE
	licenseFile := &composer.File{
		Name:       "LICENSE",
		TemplateID: "license",
		TemplateData: &config.TemplateData{
			Name: themeData.Name,
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

func setupCSSLib(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	switch tplData.Theme.CSSLib {
	case VanillaCSS:
		vanillaCSS := css.NewVanillaCSS(efs, fs, conf, tplData)
		return vanillaCSS.Setup(true)
	case Scss:
		scss := css.NewScss(efs, fs, conf, tplData)
		return scss.Setup(true)
	case TailwindCSS:
		tailwind := css.NewTailwindCSS(efs, fs, conf, tplData)
		return tailwind.Setup(true)
	case Bulma:
		bulma := css.NewBulma(efs, fs, conf, tplData)
		return bulma.Setup(true)
	case Bootstrap:
		boostrap := css.NewBootstrap(efs, fs, conf, tplData)
		return boostrap.Setup(true)
	default:
		return sveltinerr.NewOptionNotValidError(tplData.Theme.CSSLib, []string{"vanillacss", "tailwindcss", "bulma", "bootstrap", "scss"})
	}
}

func isInitGitRepo(gitFlagValue bool) bool {
	return gitFlagValue
}

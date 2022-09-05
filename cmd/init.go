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

	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	"github.com/sveltinio/sveltin/internal/css"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/npmc"
	"github.com/sveltinio/sveltin/internal/shell"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init <project>",
	Aliases: []string{"create"},
	Args:    cobra.RangeArgs(0, 3),
	Short:   "Initialize a new sveltin project",
	Long: resources.GetASCIIArt() + `
Command to scaffold a new sveltin project.

Examples:

sveltin init blog
sveltin init blog --css tailwindcss
sveltin init blog --css vanillacss -t myTheme
sveltin init portfolio -c tailwindcss -t paper -n pnpm -p 3030 --git`,
	Run: InitCmdRun,
}

// InitCmdRun is the actual work function.
func InitCmdRun(cmd *cobra.Command, args []string) {
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

	cfg.log.Plain(markup.H1("Initializing a new Sveltin project"))

	// Clone starter template github repository
	starterTemplate := cfg.startersMap[SvelteKitStarter]
	cfg.log.Info(fmt.Sprintf("Cloning the %s repos", starterTemplate.Name))

	gitClient := shell.NewGitClient()
	err = gitClient.RunGitClone(starterTemplate.URL, cfg.pathMaker.GetProjectRoot(projectName), true)
	utils.ExitIfError(err)

	// GET FOLDER: <project_name>
	cfg.log.Info("Creating the project folder structure")

	// MAKE FOLDER STRUCTURE: config folder
	configFolder, err := makeProjectFolderStructure(ConfigFolder, "", nil)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: content folder
	contentFolder, err := makeProjectFolderStructure(ContentFolder, "", nil)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes folder
	routesFolder, err := makeProjectFolderStructure(RoutesFolder, "", themeData)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: themes/<theme_name> folder
	themeFolder, err := makeProjectFolderStructure(ThemesFolder, "", themeData)
	utils.ExitIfError(err)

	// NEW FILE: env.production
	dotEnvTplData := &config.TemplateData{
		Name:    DotEnvProdFile,
		BaseURL: fmt.Sprintf("http://%s.com", projectName),
	}
	f := cfg.fsManager.NewDotEnvFile(projectName, dotEnvTplData)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(projectName)
	projectFolder.Add(f)
	projectFolder.Add(configFolder)
	projectFolder.Add(contentFolder)
	projectFolder.Add(routesFolder)
	projectFolder.Add(themeFolder)

	rootFolder := cfg.fsManager.GetFolder(RootFolder)
	rootFolder.Add(projectFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewProjectArtifact(&resources.SveltinFS, cfg.fs)
	err = rootFolder.Create(sfs)
	utils.ExitIfError(err)

	// SETUP THE CSS LIB
	cfg.log.Info("Setting up the CSS Lib")
	tplData := config.TemplateData{
		ProjectName: projectName,
		NPMClient:   npmClient.ToString(),
		PortNumber:  withPortNumber,
		Theme:       themeData,
	}
	err = setupCSSLib(&resources.SveltinFS, cfg, &tplData)
	utils.ExitIfError(err)

	// INITIALIZE GIT REPO
	if isInitGitRepo(withGit) {
		cfg.log.Info("Initializing empty Git repository")
		err = gitClient.RunInit(projectFolder.GetPath(), true)
		utils.ExitIfError(err)
	}

	cfg.log.Success("Done")

	projectConfigSummary := &common.UserProjectConfig{
		ProjectName:   projectName,
		CSSLibName:    cssLibName,
		ThemeName:     themeSelection,
		NPMClientName: npmClient.Desc,
	}

	// NEXT STEPS
	if themeData.ID != config.ExistingTheme {
		common.PrintNextStepsHelperForNewProject(projectConfigSummary)
	} else {
		common.PrintNextStepsHelperForNewProjectWithExistingTheme(projectConfigSummary)
	}

}

func initCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&npmClientName, "npmClient", "n", "", "The name of your preferred npm client")
	cmd.Flags().StringVarP(&withThemeName, "theme", "t", "", "The theme you are going to create or reuse")
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "", "The name of the CSS framework to use. Valid: vanillacss, tailwindcss, bulma, bootstrap, scss")
	cmd.Flags().StringVarP(&withPortNumber, "port", "p", "5173", "The port to start the server on")
	cmd.Flags().BoolVarP(&withGit, "git", "g", false, "Initialize an empty Git repository")
}

func init() {
	initCmdFlags(initCmd)
	rootCmd.AddCommand(initCmd)
}

//=============================================================================

func promptProjectName(inputs []string) (string, error) {
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

func promptCSSLibName(cssLibName string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: VanillaCSS, Desc: "Plain CSS"},
		choose.Item{Name: Scss, Desc: "Scss/Sass"},
		choose.Item{Name: TailwindCSS, Desc: "Tailwind CSS"},
		choose.Item{Name: Bulma, Desc: "Bulma"},
		choose.Item{Name: Bootstrap, Desc: "Bootstrap"},
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

func promptThemeSelection(themeFlag string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: config.BlankTheme, Desc: "Create a new theme"},
		choose.Item{Name: config.SveltinTheme, Desc: "Sveltin default theme"},
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

func promptNPMClient(items []string) (string, error) {
	if len(items) == 0 {
		err := errors.New("it seems there is no package manager installed on your machine. We cannot proceed now")
		return "", sveltinerr.NewNPMClientNotFoundError(err)
	}

	entries := choose.ToListItem(items)

	switch nameLenght := len(npmClientName); {
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

func buildThemeData(themeSelection, themeFlagValue, projectName, cssLibName string) (*config.ThemeData, error) {
	switch themeSelection {
	case config.BlankTheme:
		defaultThemeName := strings.Join([]string{projectName, "theme"}, "_")
		newThemePromptContent := &input.Config{
			Initial:     defaultThemeName,
			Message:     "What's the your new theme name?",
			Placeholder: "Please, provide a name for your theme.",
		}
		themeName, err := input.Run(newThemePromptContent)
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
		if utils.IsValidURL(themeSelection) {
			_, err := utils.NewGitHubURLParser(themeSelection)
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

func setupCSSLib(efs *embed.FS, cfg appConfig, tplData *config.TemplateData) error {
	switch tplData.Theme.CSSLib {
	case VanillaCSS:
		vanillaCSS := css.NewVanillaCSS(efs, cfg.fs, cfg.sveltin, tplData)
		return vanillaCSS.Setup(true)
	case Scss:
		scss := css.NewScss(efs, cfg.fs, cfg.sveltin, tplData)
		return scss.Setup(true)
	case TailwindCSS:
		tailwind := css.NewTailwindCSS(efs, cfg.fs, cfg.sveltin, tplData)
		return tailwind.Setup(true)
	case Bulma:
		bulma := css.NewBulma(efs, cfg.fs, cfg.sveltin, tplData)
		return bulma.Setup(true)
	case Bootstrap:
		boostrap := css.NewBootstrap(efs, cfg.fs, cfg.sveltin, tplData)
		return boostrap.Setup(true)
	default:
		return sveltinerr.NewOptionNotValidError(tplData.Theme.CSSLib, []string{"vanillacss", "tailwindcss", "bulma", "bootstrap", "scss"})
	}
}

func isInitGitRepo(gitFlagValue bool) bool {
	return gitFlagValue
}

//=============================================================================

func makeProjectFolderStructure(folderName string, projectName string, themeData *config.ThemeData) (*composer.Folder, error) {
	switch folderName {
	case ConfigFolder:
		return createProjectConfigLocalFolder(projectName), nil
	case ContentFolder:
		return createProjectContentLocalFolder(), nil
	case RoutesFolder:
		return createProjectRoutesLocalFolder(themeData), nil
	case ThemesFolder:
		if themeData.IsNew || themeData.ID == config.SveltinTheme {
			return createProjectThemeLocalFolder(themeData), nil
		}
		return nil, nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)

	}
}

//=============================================================================

func createProjectConfigLocalFolder(projectName string) *composer.Folder {
	// NEW FOLDER: config
	configFolder := composer.NewFolder(ConfigFolder)

	// NEW FILE: config/<filename>.js
	for _, elem := range []string{Defaults, Externals, Website, Menu} {
		f := cfg.fsManager.NewConfigFile(projectName, elem, CliVersion)
		configFolder.Add(f)
	}
	return configFolder
}

func createProjectContentLocalFolder() *composer.Folder {
	// NEW FOLDER: content
	return composer.NewFolder(ContentFolder)
}

func createProjectRoutesLocalFolder(themeData *config.ThemeData) *composer.Folder {
	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FILE: index.svelte
	indexFile := &composer.File{
		Name:       helpers.GetResourceRouteFilename(IndexFile, cfg.sveltin),
		TemplateID: IndexFile,
		TemplateData: &config.TemplateData{
			Theme: themeData,
		},
	}
	routesFolder.Add(indexFile)
	return routesFolder
}

func createProjectThemeLocalFolder(themeData *config.ThemeData) *composer.Folder {
	// NEW FOLDER: themes
	themesFolder := composer.NewFolder(ThemesFolder)

	// NEW FOLDER: themes/<theme_name>
	newThemeFolder := composer.NewFolder(themeData.Name)

	// NEW FOLDER: themes/<theme_name>/components
	componentsFolder := composer.NewFolder(cfg.pathMaker.GetThemeComponentsFolder())
	newThemeFolder.Add(componentsFolder)

	// NEW FOLDER: themes/<theme_name>/partials
	partialsFolder := composer.NewFolder(cfg.pathMaker.GetThemePartialsFolder())
	newThemeFolder.Add(partialsFolder)

	// ADD FILE themes/<theme_name>/theme.config.js
	configFile := &composer.File{
		Name:       cfg.sveltin.GetThemeConfigFilename(),
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
	themesFolder.Add(newThemeFolder)

	return themesFolder
}

//=============================================================================

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"embed"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	"github.com/sveltinio/sveltin/internal/css"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/npmc"
	"github.com/sveltinio/sveltin/internal/shell"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/feedbacks"
	"github.com/sveltinio/sveltin/tui/prompts"
	"github.com/sveltinio/sveltin/utils"
	logger "github.com/sveltinio/yinlog"
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
	Bootstrap   string = "bootstrap"
	Bulma       string = "bulma"
	Scss        string = "scss"
	TailwindCSS string = "tailwindcss"
	VanillaCSS  string = "vanillacss"
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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init [name]",
	Aliases: []string{"create"},
	Short:   "Initialize a new Sveltin project",
	Long: resources.GetASCIIArt() + `
Command used to initialize/scaffold a new sveltin project.

Examples:

sveltin init blog
sveltin init blog --css tailwindcss
sveltin init blog --css vanillacss -t myTheme
sveltin init portfolio -c tailwindcss -t paper -n pnpm -p 3030 --git
`,
	DisableFlagsInUseLine: true,
	Run:                   InitCmdRun,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, activehelps.Hint("You must choose a name for the project"))
		} else {
			comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any more arguments but accepts flags."))
		}
		return comps, cobra.ShellCompDirectiveDefault
	},
}

// InitCmdRun is the actual work function.
func InitCmdRun(cmd *cobra.Command, args []string) {
	projectName, err := prompts.AskProjectNameHandler(args)
	utils.ExitIfError(err)

	cssLibName, err := prompts.SelectCSSLibHandler(withCSSLib)
	utils.ExitIfError(err)

	themeSelection, err := prompts.SelectThemeHandler(withThemeName)
	utils.ExitIfError(err)
	themeData, err := buildThemeData(themeSelection, withThemeName, projectName, cssLibName)
	utils.ExitIfError(err)

	npmClient := getSelectedNPMClient(npmClientName, cfg.log)
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
		Name: DotEnvProdFile,
		Vite: &tpltypes.ViteData{
			BaseURL: fmt.Sprintf("http://%s.com", projectName),
		},
	}
	envFile := cfg.fsManager.NewDotEnvFile(projectName, dotEnvTplData)

	// NEW FILE: .sveltin.json
	sveltinConfigTplData := newSveltinJsonTplData(projectName, themeData)
	sveltinJSONConfigFile := cfg.fsManager.NewJSONConfigFile(sveltinConfigTplData)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(projectName)
	projectFolder.Add(configFolder)
	projectFolder.Add(contentFolder)
	projectFolder.Add(routesFolder)
	projectFolder.Add(themeFolder)
	projectFolder.Add(envFile)
	projectFolder.Add(sveltinJSONConfigFile)

	rootFolder := cfg.fsManager.GetFolder(RootFolder)
	rootFolder.Add(projectFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewProjectArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err = rootFolder.Create(sfs)
	utils.ExitIfError(err)

	// COPY FILE: sveltin.d.ts
	saveTo := path.Join(cfg.pathMaker.GetProjectRoot(projectName), cfg.pathMaker.GetSrcFolder())
	err = cfg.fsManager.CopyFileFromEmbed(&resources.SveltinStaticFS, cfg.fs, resources.SveltinFilesFS, SveltinDTSFileId, saveTo)
	utils.ExitIfError(err)

	// COPY FILE: mdsvex.config.js
	saveTo = cfg.pathMaker.GetProjectRoot(projectName)
	err = cfg.fsManager.CopyFileFromEmbed(&resources.SveltinStaticFS, cfg.fs, resources.SveltinFilesFS, MDsveXFileId, saveTo)
	utils.ExitIfError(err)

	// SETUP THE CSS LIB
	cfg.log.Info("Setting up the CSS Lib")
	tplData := config.TemplateData{
		ProjectName: projectName,
		NPMClient: &tpltypes.NPMClientData{
			Name:    npmClient.Name,
			Version: npmClient.Version,
			Info:    npmClient.ToString(),
		},
		Vite: &tpltypes.ViteData{
			Port: withPortNumber,
		},
		Theme: themeData,
	}
	err = setupCSSLib(&resources.SveltinTemplatesFS, cfg, &tplData)
	utils.ExitIfError(err)

	// INITIALIZE GIT REPO
	if isInitGitRepo(withGit) {
		cfg.log.Info("Initializing empty Git repository")
		err = gitClient.RunInit(projectFolder.GetPath(), true)
		utils.ExitIfError(err)
	}

	cfg.log.Success("Done\n")

	projectConfigSummary := config.NewProjectConfig(projectName, cssLibName, themeSelection, npmClient.Desc)
	// NEXT STEPS
	if themeData.ID != tpltypes.ExistingTheme {
		feedbacks.ShowNewProjectNextStepsHelpMessage(projectConfigSummary)
	} else {
		feedbacks.ShowNewProjectWithExistingThemeNextStepsHelpMessage(projectConfigSummary)
	}

}

func initCmdFlags(cmd *cobra.Command) {
	// npmClient flag
	cmd.Flags().StringVarP(&npmClientName, "npmClient", "n", "", "The name of your preferred npm client")
	err := cmd.RegisterFlagCompletionFunc("npmClient", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		installedNPMClients := utils.GetInstalledNPMClientList()
		npmcNames := utils.GetNPMClientNames(installedNPMClients)
		return npmcNames, cobra.ShellCompDirectiveDefault
	})
	utils.ExitIfError(err)
	// theme flag
	cmd.Flags().StringVarP(&withThemeName, "theme", "t", "", "The theme you are going to create or reuse")
	err = cmd.RegisterFlagCompletionFunc("theme", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{tpltypes.BlankTheme, tpltypes.SveltinTheme}, cobra.ShellCompDirectiveDefault
	})
	utils.ExitIfError(err)
	// css flag
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "", "The CSS lib to use. Valid: vanillacss, tailwindcss, bulma, bootstrap, scss")
	err = cmd.RegisterFlagCompletionFunc("css", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return css.AvailableCSSLib, cobra.ShellCompDirectiveDefault
	})
	utils.ExitIfError(err)
	// others
	cmd.Flags().StringVarP(&withPortNumber, "port", "p", "5173", "The port to start the server on")
	cmd.Flags().BoolVarP(&withGit, "git", "g", false, "Initialize an empty Git repository")
}

func init() {
	initCmdFlags(initCmd)
	rootCmd.AddCommand(initCmd)
}

//=============================================================================

func buildThemeData(themeSelection, themeFlagValue, projectName, cssLibName string) (*tpltypes.ThemeData, error) {
	switch themeSelection {
	case tpltypes.BlankTheme:
		defaultThemeName := strings.Join([]string{projectName, "theme"}, "_")
		newThemePromptContent := &input.Config{
			Initial:     defaultThemeName,
			Message:     "What's your theme name?",
			Placeholder: "Please, provide a name for your theme.",
		}
		themeName, err := input.Run(newThemePromptContent)
		if err != nil {
			return nil, err
		}
		return &tpltypes.ThemeData{
			ID:     tpltypes.BlankTheme,
			IsNew:  true,
			Name:   themeName,
			CSSLib: cssLibName,
		}, nil
	case tpltypes.SveltinTheme:
		return &tpltypes.ThemeData{
			ID:     tpltypes.SveltinTheme,
			IsNew:  false,
			Name:   "sveltin_theme",
			CSSLib: cssLibName,
		}, nil
	case tpltypes.ExistingTheme:
		return &tpltypes.ThemeData{
			ID:     tpltypes.ExistingTheme,
			IsNew:  false,
			CSSLib: cssLibName,
		}, nil
	default:
		if utils.IsValidURL(themeSelection) {
			_, err := utils.NewGitHubURLParser(themeSelection)
			if err != nil {
				return nil, err
			}
			return &tpltypes.ThemeData{
				ID:     tpltypes.ExistingTheme,
				IsNew:  false,
				CSSLib: cssLibName,
			}, nil
		}

		return &tpltypes.ThemeData{
			ID:     tpltypes.BlankTheme,
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
func getSelectedNPMClient(npmcFlag string, logger *logger.Logger) npmc.NPMClient {
	installedNPMClients := utils.GetInstalledNPMClientList()
	npmcNames := utils.GetNPMClientNames(installedNPMClients)
	client, err := prompts.SelectNPMClientHandler(npmcNames, npmcFlag, logger)
	utils.ExitIfError(err)
	return utils.GetSelectedNPMClient(installedNPMClients, client)
}

func setupCSSLib(efs *embed.FS, cfg appConfig, tplData *config.TemplateData) error {
	switch tplData.Theme.CSSLib {
	case VanillaCSS:
		vanillaCSS := css.NewVanillaCSS(efs, cfg.fs, cfg.settings, tplData)
		return vanillaCSS.Setup(true)
	case Scss:
		scss := css.NewScss(efs, cfg.fs, cfg.settings, tplData)
		return scss.Setup(true)
	case TailwindCSS:
		tailwind := css.NewTailwindCSS(efs, cfg.fs, cfg.settings, tplData)
		return tailwind.Setup(true)
	case Bulma:
		bulma := css.NewBulma(efs, cfg.fs, cfg.settings, tplData)
		return bulma.Setup(true)
	case Bootstrap:
		boostrap := css.NewBootstrap(efs, cfg.fs, cfg.settings, tplData)
		return boostrap.Setup(true)
	default:
		return sveltinerr.NewOptionNotValidError(tplData.Theme.CSSLib, []string{"vanillacss", "tailwindcss", "bulma", "bootstrap", "scss"})
	}
}

func isInitGitRepo(gitFlagValue bool) bool {
	return gitFlagValue
}

//=============================================================================

func makeProjectFolderStructure(folderName string, projectName string, themeData *tpltypes.ThemeData) (*composer.Folder, error) {
	switch folderName {
	case ConfigFolder:
		return createProjectConfigLocalFolder(projectName), nil
	case ContentFolder:
		return createProjectContentLocalFolder(), nil
	case RoutesFolder:
		return createProjectRoutesLocalFolder(themeData), nil
	case ThemesFolder:
		if themeData.IsNew || themeData.ID == tpltypes.SveltinTheme {
			return createProjectThemeLocalFolder(themeData), nil
		}
		return nil, nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

func newSveltinJsonTplData(projectName string, themeData *tpltypes.ThemeData) *config.TemplateData {
	return &config.TemplateData{
		Name: ProjectSettingsFile,
		ProjectSettings: &tpltypes.ProjectSettings{
			Name:    projectName,
			BaseURL: fmt.Sprintf("http://%s.com", projectName),
			Sitemap: tpltypes.SitemapData{
				ChangeFreq: "monthly",
				Priority:   0.5,
			},
			Theme: *themeData,
			Sveltin: tpltypes.SveltinCLIData{
				Version: CliVersion,
			},
		},
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

func createProjectRoutesLocalFolder(themeData *tpltypes.ThemeData) *composer.Folder {
	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FILE: index.svelte
	indexFile := &composer.File{
		Name:       helpers.GetResourceRouteFilename(IndexFileId, cfg.settings),
		TemplateID: IndexFileId,
		TemplateData: &config.TemplateData{
			Theme: themeData,
		},
	}
	routesFolder.Add(indexFile)
	return routesFolder
}

func createProjectThemeLocalFolder(themeData *tpltypes.ThemeData) *composer.Folder {
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
		Name:       cfg.settings.GetThemeConfigFilename(),
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
			Theme: themeData,
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
}

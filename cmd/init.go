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
	"github.com/sveltinio/sveltin/internal/gitclient"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/notifier"
	"github.com/sveltinio/sveltin/internal/npmc"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/feedbacks"
	"github.com/sveltinio/sveltin/tui/prompts"
	"github.com/sveltinio/sveltin/utils"
	logger "github.com/sveltinio/yinlog"
)

// names for the available style options.
const (
	StyleDefault string = "default"
	StyleNone    string = "none"
)

// names for config files.
const (
	Defaults  string = "defaults"
	Externals string = "externals"
	Website   string = "website"
	Menu      string = "menu"
	DotEnv    string = "dotenv"
)

var (
	// How to use the command.
	initCmdExample = `sveltin init blog --css tailwindcss
sveltin init blog --css unocss -n pnpm -p 3030 --git`
	// Short description shown in the 'help' output.
	initCmdShortMsg = "Initialize a new Sveltin project"
	// Long message shown in the 'help <this-command>' output.
	initCmdLongMsg = utils.MakeCmdLongMsg("Command used to initialize/scaffold a new sveltin project.")
)

// Bind command flags.
var (
	withCSSLib     string
	withThemeName  string
	withPortNumber string
	withGit        bool
)

//=============================================================================

var initCmd = &cobra.Command{
	Use:               "init [project_name]",
	Aliases:           []string{"create"},
	Example:           initCmdExample,
	Short:             initCmdShortMsg,
	Long:              initCmdLongMsg,
	ValidArgsFunction: initCmdValidArgs,
	PreRun:            initCmdPreRunHook,
	Run:               InitCmdRun,
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

	cfg.log.Plain(markup.H1("Setup a new Sveltin project"))

	// Clone starter template github repository
	starterTemplate := cfg.startersMap[SvelteKitStarter]
	cfg.log.Info(fmt.Sprintf("Getting the %s", starterTemplate.Name))

	err = gitclient.RunClone(starterTemplate.URL, CliVersion, cfg.pathMaker.GetProjectRoot(projectName))
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
			BaseURL: fmt.Sprintf("https://%s", projectName),
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
		cfg.log.Info("Initializing an empty Git repository")
		err = gitclient.RunInit(projectFolder.GetPath())
		//err = gitClient.RunInit(projectFolder.GetPath(), true)
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

// Command initialization.
func init() {
	initCmdFlags(initCmd)
	rootCmd.AddCommand(initCmd)
}

//=============================================================================

// Assign flags to the command.
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
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "",
		fmt.Sprintf("The CSS lib to use. Valid: %s, %s, %s, %s, %s, %s",
			css.Bootstrap, css.Bulma, css.Scss, css.TailwindCSS, css.UnoCSS, css.VanillaCSS))
	err = cmd.RegisterFlagCompletionFunc("css", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return css.AvailableCSSLib, cobra.ShellCompDirectiveDefault
	})
	utils.ExitIfError(err)

	// others
	cmd.Flags().StringVarP(&withPortNumber, "port", "p", "5173", "The port to start the server on")
	cmd.Flags().BoolVarP(&withGit, "git", "g", false, "Initialize an empty Git repository")
}

// Adding Active Help messages enhancing shell completions.
func initCmdValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	if len(args) == 0 {
		comps = cobra.AppendActiveHelp(comps, activehelps.Hint("You must choose a name for the project"))
	} else {
		comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any more arguments but accepts flags."))
	}
	return comps, cobra.ShellCompDirectiveDefault
}

// Run before the main Run function of init command to check and alert about newer version.
func initCmdPreRunHook(cmd *cobra.Command, args []string) {
	handleReleaseNotifier(notifier.InitCmd)
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
	case css.Bootstrap:
		boostrap := css.NewBootstrap(efs, cfg.fs, cfg.settings, tplData)
		return boostrap.Setup(true)
	case css.Bulma:
		bulma := css.NewBulma(efs, cfg.fs, cfg.settings, tplData)
		return bulma.Setup(true)
	case css.Scss:
		scss := css.NewScss(efs, cfg.fs, cfg.settings, tplData)
		return scss.Setup(true)
	case css.TailwindCSS:
		tailwind := css.NewTailwindCSS(efs, cfg.fs, cfg.settings, tplData)
		return tailwind.Setup(true)
	case css.UnoCSS:
		unocss := css.NewUnoCSS(efs, cfg.fs, cfg.settings, tplData)
		return unocss.Setup(true)
	case css.VanillaCSS:
		vanillaCSS := css.NewVanillaCSS(efs, cfg.fs, cfg.settings, tplData)
		return vanillaCSS.Setup(true)
	default:
		return sveltinerr.NewOptionNotValidError(
			tplData.Theme.CSSLib,
			[]string{
				css.Bootstrap,
				css.Bulma,
				css.Scss,
				css.TailwindCSS,
				css.UnoCSS,
				css.VanillaCSS,
			})
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
			BaseURL: fmt.Sprintf("https://%s", projectName),
			Sitemap: tpltypes.SitemapData{
				ChangeFreq: "monthly",
				Priority:   0.5,
			},
			Theme: *themeData,
			Sveltin: tpltypes.SveltinCLIData{
				Version:      CliVersion,
				CheckUpdates: utils.NewTrue(),
				LastCheck:    utils.TodayISO(),
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

	// NEW FILE: src/routes/{+page.svelte, +page.server.ts}
	for _, item := range []string{IndexPageFileId, IndexPageLoadFileId} {
		f := &composer.File{
			Name:       helpers.GetRouteFilename(item, cfg.settings),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Theme: themeData,
			},
		}
		routesFolder.Add(f)
	}

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

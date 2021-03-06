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
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/pkg/composer"
	"github.com/sveltinio/sveltin/pkg/css"
	"github.com/sveltinio/sveltin/pkg/shell"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var newThemeCmd = &cobra.Command{
	Use:     "theme <name>",
	Aliases: []string{"t"},
	Short:   "Command to create a new theme",
	Long: resources.GetASCIIArt() + `
This command help you creating new themes for projects, so that can be shared with others and reused.

Examples:

sveltin new theme paper
sveltin new theme paper --css tailwindcss
`,
	Run: NewThemeCmdRun,
}

// NewThemeCmdRun is the actual work function.
func NewThemeCmdRun(cmd *cobra.Command, args []string) {
	// Exit if running the command from an existing sveltin project folder.
	isValidForThemeMaker()

	themeName, err := promptThemeName(args)
	utils.ExitIfError(err)
	cfg.log.Info(themeName)

	projectName := themeName + "_project"

	cssLibName, err := promptCSSLibName(withCSSLib)
	utils.ExitIfError(err)
	cfg.log.Info(cssLibName)

	npmClient := getSelectedNPMClient()
	npmClientName = npmClient.Name

	cfg.log.Plain(utils.Underline("A Starter project will be created"))

	// Clone starter template github repository
	themeStarterTemplate := cfg.startersMap[ThemeStarter]
	cfg.log.Info(fmt.Sprintf("Cloning the %s repos", themeStarterTemplate.Name))

	gitClient := shell.NewGitClient()
	err = gitClient.RunGitClone(themeStarterTemplate.URL, cfg.pathMaker.GetProjectRoot(projectName), true)
	// TO BE REMOVED: err = utils.GitClone(themeStarterTemplate.URL, pathMaker.GetProjectRoot(projectName))
	utils.ExitIfError(err)

	// NEW FILE: config/defaults.js
	f := cfg.fsManager.NewConfigFile(projectName, Defaults, CliVersion)
	// NEW FOLDER: config
	configFolder := composer.NewFolder(ConfigFolder)
	configFolder.Add(f)

	// MAKE FOLDER STRUCTURE: themes/<theme_name> folder
	themeData := &config.ThemeData{
		ID:     config.BlankTheme,
		IsNew:  true,
		Name:   themeName,
		CSSLib: cssLibName,
	}
	themeFolder, err := makeProjectFolderStructure(ThemesFolder, "", themeData)
	utils.ExitIfError(err)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(projectName)
	projectFolder.Add(configFolder)
	projectFolder.Add(themeFolder)

	rootFolder := cfg.fsManager.GetFolder(RootFolder)
	rootFolder.Add(projectFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewThemeArtifact(&resources.SveltinFS, cfg.fs)
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
	err = setupThemeCSSLib(&resources.SveltinFS, cfg, &tplData)
	utils.ExitIfError(err)

	cfg.log.Success("Done")

	// NEXT STEPS
	cfg.log.Plain(utils.Underline("Next Steps"))
	cfg.log.Plain(common.HelperTextNewTheme(projectName))
}

func newThemeCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&withCSSLib, "css", "c", "", "The name of the CSS framework to use. Possible values: vanillacss, tailwindcss, bulma, bootstrap, scss")
	cmd.Flags().StringVarP(&npmClientName, "npmClient", "n", "", "The name of your preferred npm client")
}

func init() {
	newThemeCmdFlags(newThemeCmd)
	newCmd.AddCommand(newThemeCmd)
}

//=============================================================================

// isValidForThemeMaker returns error if find the package.json file within the current folder.
func isValidForThemeMaker() {
	pwd, _ := os.Getwd()
	pathToPkgJSON := filepath.Join(pwd, "package.json")
	exists, _ := afero.Exists(cfg.fs, pathToPkgJSON)
	if exists {
		err := sveltinerr.NewNotEmptyProjectError(pathToPkgJSON)
		jww.FATAL.Fatalf("\x1b[31;1m✘ %s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
}

func promptThemeName(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		themeNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for the theme.",
			Label:    "What's the theme name?",
		}
		result, err := common.PromptGetInput(themeNamePromptContent, nil, "")
		if err != nil {
			return "", err
		}
		return utils.ToSlug(result), nil
	case numOfArgs == 1:
		return utils.ToSlug(inputs[0]), nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

func setupThemeCSSLib(efs *embed.FS, cfg appConfig, tplData *config.TemplateData) error {
	switch tplData.Theme.CSSLib {
	case VanillaCSS:
		vanillaCSS := css.NewVanillaCSS(efs, cfg.fs, cfg.sveltin, tplData)
		return vanillaCSS.Setup(false)
	case Scss:
		scss := css.NewScss(efs, cfg.fs, cfg.sveltin, tplData)
		return scss.Setup(false)
	case TailwindCSS:
		tailwind := css.NewTailwindCSS(efs, cfg.fs, cfg.sveltin, tplData)
		return tailwind.Setup(false)
	case Bulma:
		bulma := css.NewBulma(efs, cfg.fs, cfg.sveltin, tplData)
		return bulma.Setup(false)
	case Bootstrap:
		boostrap := css.NewBootstrap(efs, cfg.fs, cfg.sveltin, tplData)
		return boostrap.Setup(false)
	default:
		return sveltinerr.NewOptionNotValidError(tplData.Theme.CSSLib, []string{"vanillacss", "tailwindcss", "bulma", "bootstrap", "scss"})
	}
}

//=============================================================================

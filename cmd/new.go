/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

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
	Short:   "Command to create projects, resources, contents, pages, metadata, theme",
	Long: resources.GetAsciiArt() + `
This command creates customized resources like blog posts, a new theme for your website, new page etc. depending on the subcommand used with it.

Examples:

sveltin new blog
sveltin new resource posts
sveltin new theme basic`,
	Run: NewCmdRun,
}

func NewCmdRun(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "A new Sveltin based project will be created",
	}

	err := common.CheckMinMaxArgs(args, 0, 3)
	utils.CheckIfError(err)

	projectName, err := promptProjectName(args)
	utils.CheckIfError(err)

	setPackageManager()

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

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func newCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&packageManager, "package-manager", "p", "", "The name of the your preferred package manager.")
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

func promptPackageManager(items []string) (string, error) {
	var pm string
	switch nameLenght := len(packageManager); {
	case nameLenght == 0:
		if len(items) == 1 {
			pm = items[0]
		} else {
			pmPromptContent := config.PromptContent{
				ErrorMsg: "Please, provide the name of the package manager.",
				Label:    "Which package manager do you want to use?",
			}
			_pm := common.PromptGetSelect(items, pmPromptContent)
			if common.Contains(items, _pm) {
				pm = _pm
			} else {
				errN := errors.New("invalid selection. Valid options are " + strings.Join(items, ", "))
				return "", sveltinerr.NewDefaultError(errN)
			}
		}
		return pm, nil
	case nameLenght != 0:
		if !common.Contains(items, packageManager) {
			return "", sveltinerr.NewOptionNotValidError()
		}
		pm = packageManager
		return pm, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

/**
 * Read the settings file, if it does not exists and no -p flag,
 * prompt to select the package manager from the ones currently
 * installed on the machine and store its value as settings.
 */
func setPackageManager() {
	if len(settings.GetPackageManager()) == 0 && len(packageManager) == 0 {
		selectedPackageManager, err := promptPackageManager(utils.GetAvailablePackageMangerList())
		utils.CheckIfError(err)
		storeSelectedPackageManager(selectedPackageManager)
	}
}

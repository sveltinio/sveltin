/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
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
	common.CheckIfError(err)

	projectName, err := promptProjectName(args)
	common.CheckIfError(err)

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
	common.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func init() {
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
		return "", common.NewDefaultError(err)
	}
}

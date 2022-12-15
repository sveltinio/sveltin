/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package feedbacks

import (
	"fmt"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/utils"
	logger "github.com/sveltinio/yinlog"
)

// ShowNewProjectNextStepsHelpMessage prints an help message as next steps for a project creation.
func ShowNewProjectNextStepsHelpMessage(uc *config.ProjectConfig) {
	steps := []string{
		fmt.Sprintf("cd %s", uc.ProjectName),
		fmt.Sprintf("sveltin install %s", markup.Faint("(or npm run install, pnpm install, ...)")),
		fmt.Sprintf("sveltin server %s", markup.Faint(" (or npm run dev, pnpm dev, ...)")),
	}
	nextStepsMsg := markup.NewOLWithTitle("Next Steps", steps)

	fmt.Printf("%s%s\n", nextStepsMsg, devServerInfoMessage())
}

// ShowNewProjectWithExistingThemeNextStepsHelpMessage pprints an help message as next steps for a project creation with an existing theme'.
func ShowNewProjectWithExistingThemeNextStepsHelpMessage(uc *config.ProjectConfig) {
	steps := []string{
		fmt.Sprintf("cd %s", uc.ProjectName),
		fmt.Sprintf("sveltin install %s", markup.Faint("(or npm run install, pnpm install, ...)")),
		"git init",
		"git submodule add <github_repu_url_for_the_theme> themes/<theme_name>",
		"Follow the instructions on the README file from the theme creator",
		fmt.Sprintf("sveltin server %s", markup.Faint(" (or npm run dev, pnpm dev, ...)")),
	}
	nextStepsMsg := markup.NewOLWithTitle("Next Steps", steps)

	fmt.Printf("%s%s\n", nextStepsMsg, devServerInfoMessage())
}

// ShowNewResourceHelpMessage prints an help message string for 'resource creation'.
func ShowNewResourceHelpMessage(name string) {
	exampleString := fmt.Sprintf("sveltin new content %s/getting-started", name)
	entries := []string{
		markup.P("Start by adding content to it, e.g."),
		markup.BR,
		markup.Code(exampleString),
	}

	fmt.Println(markup.Section("Resource ready to be used.", entries))
}

// ShowNewMetadataHelpMessage prints an help message string for 'metadata creation'.
func ShowNewMetadataHelpMessage(metadataInfo *tpltypes.MetadataData) {
	var exampleString string
	if metadataInfo.Type == "single" {
		exampleString = fmt.Sprintf("%s: your_value", utils.ToSnakeCase(metadataInfo.Name))
	} else {
		exampleString = utils.ToSnakeCase(metadataInfo.Name) + `:
  - value 1
  - value 2`
	}

	entries := []string{
		markup.P("Ensure your markdown frontmatter includes it, e.g."),
		markup.BR,
		markup.Code(exampleString),
	}

	fmt.Println(markup.Section("Metadata ready to be used.", entries))
}

// ShowNewThemeHelpMessage returns an help message string for 'theme creation'.
func ShowNewThemeHelpMessage(pc *config.ProjectConfig) {
	steps := []string{
		fmt.Sprintf("cd %s", pc.ProjectName),
		fmt.Sprintf("sveltin install %s", markup.Faint("(or npm run install, pnpm install, ...)")),
		"Create your theme components and partials",
		fmt.Sprintf("sveltin server %s", markup.Faint(" (or npm run dev, pnpm dev, ...)")),
	}
	nextStepsMsg := markup.NewOLWithTitle("Next Steps", steps)

	fmt.Printf("%s\n%s\n", nextStepsMsg, themingInfoMessage())
}

// ShowDryRunMessage prints a message box for commands supporting the 'dry-run mode'.
func ShowDryRunMessage() {
	fmt.Println(markup.Bordered(markup.Centered(fmt.Sprintf("%s\n\n%s", markup.Underline("DRY-RUN MODE"), "Nothing will really happen! Just simulating the process."))))
}

// ShowDeployCommandWarningMessages display a set of useful information for the deploy over FTP process.
func ShowDeployCommandWarningMessages(isBackup bool) {
	listLogger := logger.NewListLogger()
	listLogger.Logger.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    true,
		Icons:     true,
	})

	listLogger.Title("Be aware! The deploy command will perform the following actions")
	if isBackup {
		listLogger.Append(logger.WarningLevel, "Create a backup of the existing content on the remote folder")
	}
	listLogger.Append(logger.WarningLevel, "Delete existing content except what specified with --exclude or --withExcludeFile flags")
	listLogger.Append(logger.WarningLevel, "Upload content to the remote folder")
	listLogger.Render()
}

// ShowUpgradeCommandMessage display a set of useful information when running the upgrade command.
func ShowUpgradeCommandMessage() {
	listLogger := logger.NewListLogger()
	listLogger.Logger.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    true,
		Icons:     true,
	})

	infoText := `
 Sveltin will try to migrate as much as it can to make your project ready.

 Something cannot be easily handled. The main goal is to update sveltin's files.
 If there is something purely related to SvelteKit we do not cover by migrations,
 you will see error from SvelteKit by running the server as usual.
`
	listLogger.Title("\n[INFO] Migrate Project Command")
	listLogger.Append(logger.DefaultLevel, infoText)

	listLogger.Append(logger.ImportantLevel, "Ensure to commit your changes to keep track of what the command applies")
	listLogger.Render()
}

func devServerInfoMessage() string {
	return markup.Section("To stop the dev server, hit Ctrl-C",
		[]string{
			markup.Faint("Visit the Quick Start guide") +
				markup.Divider +
				markup.A("https://docs.sveltin.io/quick-start")})
}

func themingInfoMessage() string {
	return markup.Section("To stop the dev server, hit Ctrl-C",
		[]string{
			markup.Faint("Visit the Theme guide") +
				markup.Divider +
				markup.A("https://docs.sveltin.io/theming"),
		})
}

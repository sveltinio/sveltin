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
	exampleString := fmt.Sprintf("sveltin add content getting-started --to %s", name)
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
		exampleString = fmt.Sprintf(`%[1]v:
  - value 1
  - value 2

or %[1]v: ['value_1', 'value_2']`, utils.ToSnakeCase(metadataInfo.Name))
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

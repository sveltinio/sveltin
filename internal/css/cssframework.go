/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package css ...
package css

import (
	"embed"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/resources"
)

// CSSLib identifies the CSS lib to be used.
type CSSLib struct {
	Name    string
	EFS     *embed.FS
	FS      afero.Fs
	Config  *config.SveltinConfig
	TplData *config.TemplateData
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *CSSLib) Setup(isNewProject bool) error {
	// When creating a new Theme (sveltin new theme)
	if !isNewProject {
		return makeTheme(cssLib)
	}

	// When creating a fresh new Project (sveltin new <project_name>)
	switch cssLib.TplData.Theme.ID {
	case config.BlankTheme:
		return makeUnstyled(cssLib)
	case config.SveltinTheme:
		return makeSveltinStyled(cssLib)
	case config.ExistingTheme:
		return makeTheme(cssLib)
	default:
		return sveltinerr.NewOptionNotValidError(cssLib.TplData.Theme.Name, config.AvailableThemes[:])
	}
}

func makeSveltinStyled(cssLib *CSSLib) error {
	embeddedResources, err := makeResourcesMapForStyled(cssLib)
	if err != nil {
		return err
	}
	// Copying the package.json config file
	sourceFile := embeddedResources[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content := template.Run(cssLib.EFS)
	saveAs := filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = embeddedResources[SvelteConfigFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying vite.config.ts file
	sourceFile = embeddedResources[ViteConfigFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "vite.config.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	err = copyAdditionalConfigFiles(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = embeddedResources[AppHTMLFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	err = copyStylesheets(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying +layout.svelte. file
	sourceFile = embeddedResources[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +layout.ts. file
	sourceFile = embeddedResources[LayoutTSFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +error.svelte. file
	sourceFile = embeddedResources[ErrorFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+error.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying Hero.svelte component
	sourceFile = embeddedResources[HeroFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = embeddedResources[FooterFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "partials", "Footer.svelte")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	return nil
}

func makeUnstyled(cssLib *CSSLib) error {
	embeddedResources, err := makeResourcesMapForUnstyled(cssLib)
	if err != nil {
		return err
	}

	// Copying the package.json config file
	sourceFile := embeddedResources[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content := template.Run(cssLib.EFS)
	saveAs := filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = embeddedResources[SvelteConfigFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying vite.config.ts file
	sourceFile = embeddedResources[ViteConfigFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "vite.config.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	err = copyAdditionalConfigFiles(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = embeddedResources[AppHTMLFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	err = copyStylesheets(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying +layout.svelte. file
	sourceFile = embeddedResources[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +layout.ts. file
	sourceFile = embeddedResources[LayoutTSFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +error.svelte. file
	sourceFile = embeddedResources[ErrorFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+error.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying Hero.svelte component
	sourceFile = embeddedResources[HeroFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	return nil
}

func makeTheme(cssLib *CSSLib) error {
	embeddedResources, err := makeResourcesMapForUnstyled(cssLib)
	if err != nil {
		return err
	}
	// Copying the package.json config file
	sourceFile := embeddedResources[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content := template.Run(cssLib.EFS)
	saveAs := filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = embeddedResources[SvelteConfigFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying vite.config.ts file
	sourceFile = embeddedResources[ViteConfigFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "vite.config.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	err = copyAdditionalConfigFiles(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = embeddedResources[AppHTMLFileID]
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	err = copyStylesheets(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying +layout.svelte. file
	sourceFile = embeddedResources[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +layout.ts. file
	sourceFile = embeddedResources[LayoutTSFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +error.svelte. file
	sourceFile = embeddedResources[ErrorFileID]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+error.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	return nil
}

func makeResourcesMapForStyled(cssLib *CSSLib) (map[string]string, error) {
	switch cssLib.Name {
	case Bootstrap:
		return resources.BootstrapSveltinThemeFS, nil
	case Bulma:
		return resources.BulmaSveltinThemeFS, nil
	case Scss:
		return resources.SCSSSveltinThemeFS, nil
	case TailwindCSS:
		return resources.TailwindSveltinThemeFS, nil
	case VanillaCSS:
		return resources.VanillaSveltinThemeFS, nil
	default:
		return nil, sveltinerr.NewOptionNotValidError(cssLib.Name, AvailableCSSLib)
	}
}

func makeResourcesMapForUnstyled(cssLib *CSSLib) (map[string]string, error) {
	switch cssLib.Name {
	case Bootstrap:
		return resources.BootstrapBlankThemeFS, nil
	case Bulma:
		return resources.BulmaBlankThemeFS, nil
	case Scss:
		return resources.SCSSBlankThemeFS, nil
	case TailwindCSS:
		return resources.TailwindBlankThemeFS, nil
	case VanillaCSS:
		return resources.VanillaBlankThemeFS, nil
	default:
		return nil, sveltinerr.NewOptionNotValidError(cssLib.Name, AvailableCSSLib)
	}
}

func copyAdditionalConfigFiles(embeddedResources map[string]string, cssLib *CSSLib) error {
	if cssLib.Name == TailwindCSS {
		// Copying tailwindcss config file
		sourceFile := embeddedResources[TailwindConfigFileID]
		saveAs := filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "tailwind.config.cjs")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
		// Copying postcss config file
		sourceFile = embeddedResources[PostCSSFileID]
		saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "postcss.config.cjs")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	}
	return nil
}

func copyStylesheets(embeddedResources map[string]string, cssLib *CSSLib) error {
	if cssLib.Name == TailwindCSS || cssLib.Name == VanillaCSS {
		// Copying app.css file
		sourceFile := embeddedResources[AppCSSFileID]
		saveAs := filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.css")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	} else {
		// Copying app.scss file
		sourceFile := embeddedResources[AppCSSFileID]
		saveAs := filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.scss")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
		// Copying variables.scss file
		sourceFile = embeddedResources[VariablesFileID]
		saveAs = filepath.Join(cssLib.Config.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "_variables.scss")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	}
	return nil
}

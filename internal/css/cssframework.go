/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package css implements the Template Method design pattern to allow the project setup based on the selected CSSLib.
package css

import (
	"embed"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
)

// CSSLib identifies the CSS lib to be used.
type CSSLib struct {
	Name     string
	EFS      *embed.FS
	FS       afero.Fs
	Settings *config.SveltinSettings
	TplData  *config.TemplateData
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *CSSLib) Setup(isNewProject bool) error {
	// When creating a new Theme (sveltin new theme)
	if !isNewProject {
		return makeTheme(cssLib)
	}

	// When creating a fresh new Project (sveltin new <project_name>)
	switch cssLib.TplData.Theme.ID {
	case tpltypes.BlankTheme:
		return makeUnstyled(cssLib)
	case tpltypes.SveltinTheme:
		return makeSveltinStyled(cssLib)
	case tpltypes.ExistingTheme:
		return makeTheme(cssLib)
	default:
		return sveltinerr.NewOptionNotValidError(cssLib.TplData.Theme.Name, tpltypes.AvailableThemes[:])
	}
}

func makeSveltinStyled(cssLib *CSSLib) error {
	embeddedResources, err := makeResourcesMapForStyled(cssLib)
	if err != nil {
		return err
	}
	// Copying the package.json config file
	sourceFile := embeddedResources[PackageJSONFileId]
	template := helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content := template.Run(cssLib.EFS)
	saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = embeddedResources[SvelteConfigFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying msdvex.config.js file
	sourceFile = embeddedResources[MDsveXFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "mdsvex.config.js")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying vite.config.ts file
	sourceFile = embeddedResources[ViteConfigFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "vite.config.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	err = copyAdditionalConfigFiles(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = embeddedResources[AppHTMLFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	err = copyStylesheets(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying +layout.svelte. file
	sourceFile = embeddedResources[LayoutFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +layout.ts. file
	sourceFile = embeddedResources[LayoutTSFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +error.svelte. file
	sourceFile = embeddedResources[ErrorFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+error.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying md-layout.svelte. file
	sourceFile = embeddedResources[MDLayoutFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "components", "md-layout.svelte")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	// Copying Hero.svelte component
	sourceFile = embeddedResources[HeroFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = embeddedResources[FooterFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "partials", "Footer.svelte")
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
	sourceFile := embeddedResources[PackageJSONFileId]
	template := helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content := template.Run(cssLib.EFS)
	saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = embeddedResources[SvelteConfigFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	// Copying msdvex.config.js file
	sourceFile = embeddedResources[MDsveXFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "mdsvex.config.js")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying vite.config.ts file
	sourceFile = embeddedResources[ViteConfigFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "vite.config.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	err = copyAdditionalConfigFiles(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = embeddedResources[AppHTMLFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	err = copyStylesheets(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying +layout.svelte. file
	sourceFile = embeddedResources[LayoutFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +layout.ts. file
	sourceFile = embeddedResources[LayoutTSFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +error.svelte. file
	sourceFile = embeddedResources[ErrorFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+error.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying md-layout.svelte. file
	sourceFile = embeddedResources[MDLayoutFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "components", "md-layout.svelte")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	// Copying Hero.svelte component
	sourceFile = embeddedResources[HeroFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "themes", cssLib.TplData.Theme.Name, "partials", "Hero.svelte")
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
	sourceFile := embeddedResources[PackageJSONFileId]
	template := helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content := template.Run(cssLib.EFS)
	saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = embeddedResources[SvelteConfigFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying vite.config.ts file
	sourceFile = embeddedResources[ViteConfigFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "vite.config.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	err = copyAdditionalConfigFiles(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = embeddedResources[AppHTMLFileId]
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
		return err
	}

	err = copyStylesheets(embeddedResources, cssLib)
	if err != nil {
		return err
	}

	// Copying +layout.svelte. file
	sourceFile = embeddedResources[LayoutFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +layout.ts. file
	sourceFile = embeddedResources[LayoutTSFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+layout.ts")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	// Copying +error.svelte. file
	sourceFile = embeddedResources[ErrorFileId]
	template = helpers.BuildTemplate(sourceFile, nil, cssLib.TplData)
	content = template.Run(cssLib.EFS)
	saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "routes", "+error.svelte")
	if err := helpers.WriteContentToDisk(cssLib.FS, saveAs, content); err != nil {
		return err
	}

	return nil
}

func makeResourcesMapForStyled(cssLib *CSSLib) (map[string]string, error) {
	switch cssLib.Name {
	case Bootstrap:
		return resources.BootstrapSveltinThemeFilesMap, nil
	case Bulma:
		return resources.BulmaSveltinThemeFilesMap, nil
	case Scss:
		return resources.SassSveltinThemeFilesMap, nil
	case TailwindCSS:
		return resources.TailwindSveltinThemeFilesMap, nil
	case VanillaCSS:
		return resources.VanillaSveltinThemeFilesMap, nil
	default:
		return nil, sveltinerr.NewOptionNotValidError(cssLib.Name, AvailableCSSLib)
	}
}

func makeResourcesMapForUnstyled(cssLib *CSSLib) (map[string]string, error) {
	switch cssLib.Name {
	case Bootstrap:
		return resources.BootstrapBlankThemeFilesMap, nil
	case Bulma:
		return resources.BulmaBlankThemeFilesMap, nil
	case Scss:
		return resources.SassBlankThemeFilesMap, nil
	case TailwindCSS:
		return resources.TailwindBlankThemeFilesMap, nil
	case VanillaCSS:
		return resources.VanillaBlankThemeFilesMap, nil
	default:
		return nil, sveltinerr.NewOptionNotValidError(cssLib.Name, AvailableCSSLib)
	}
}

func copyAdditionalConfigFiles(embeddedResources map[string]string, cssLib *CSSLib) error {
	if cssLib.Name == TailwindCSS {
		// Copying tailwindcss config file
		sourceFile := embeddedResources[TailwindConfigFileId]
		saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "tailwind.config.cjs")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
		// Copying postcss config file
		sourceFile = embeddedResources[PostCSSFileId]
		saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "postcss.config.cjs")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	}
	return nil
}

func copyStylesheets(embeddedResources map[string]string, cssLib *CSSLib) error {

	switch cssLib.Name {
	case Bootstrap, Bulma, Scss, VanillaCSS:
		// Copying tw-preflight.css file
		sourceFile := embeddedResources[ResetCSSFileId]
		saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "tw-preflight.css")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	default:
		break
	}

	switch cssLib.Name {
	case TailwindCSS, VanillaCSS:
		// Copying app.css file
		sourceFile := embeddedResources[AppCSSFileId]
		saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.css")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	case Bootstrap, Bulma, Scss:
		// Copying app.scss file
		sourceFile := embeddedResources[AppCSSFileId]
		saveAs := filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "app.scss")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
		// Copying variables.scss file
		sourceFile = embeddedResources[VariablesFileId]
		saveAs = filepath.Join(cssLib.Settings.GetProjectRoot(), cssLib.TplData.ProjectName, "src", "_variables.scss")
		if err := common.MoveFile(cssLib.EFS, cssLib.FS, sourceFile, saveAs, false); err != nil {
			return err
		}
	}

	return nil
}

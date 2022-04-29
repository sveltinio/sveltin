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
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// Scss identifies the CSS lib to be used.
type Scss struct {
	EFS    *embed.FS
	FS     afero.Fs
	Config *config.SveltinConfig
	Data   *config.TemplateData
}

// NewScss returns a pointer to a Scss struct.
func NewScss(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) *Scss {
	return &Scss{
		EFS:    efs,
		FS:     fs,
		Config: conf,
		Data:   tplData,
	}
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *Scss) Setup(isNewProject bool) error {
	// When creating a new Theme (sveltin new theme)
	if !isNewProject {
		return cssLib.makeTheme(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
	}

	// When creating a fresh new Project (sveltin new <project_name>)
	switch cssLib.Data.Theme.ID {
	case config.BlankTheme:
		return cssLib.makeUnstyled(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
	case config.SveltinTheme:
		return cssLib.makeSveltinStyled(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
	case config.ExistingTheme:
		return cssLib.makeTheme(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
	default:
		return sveltinerr.NewOptionNotValidError(cssLib.Data.Theme.Name, config.AvailableThemes[:])
	}
}

func (cssLib *Scss) makeSveltinStyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	scssFS := common.UnionMap(resources.SveltinSCSSLibFS, resources.SveltinSCSSLibStyledFS)

	// Copying the package.json config file
	sourceFile := scssFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = scssFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = scssFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = scssFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = scssFS[VariablesFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "_variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = scssFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = scssFS[HeroFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = scssFS[FooterFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Footer.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *Scss) makeUnstyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	scssFS := common.UnionMap(resources.SveltinSCSSLibFS, resources.SveltinSCSSLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := scssFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = scssFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = scssFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = scssFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = scssFS[VariablesFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "_variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = scssFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = scssFS[HeroFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *Scss) makeTheme(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	scssFS := common.UnionMap(resources.SveltinSCSSLibFS, resources.SveltinSCSSLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := scssFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = scssFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = scssFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = scssFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = scssFS[VariablesFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "_variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = scssFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	return nil
}

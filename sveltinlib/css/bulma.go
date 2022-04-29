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

// Bulma identifies the CSS lib to be used.
type Bulma struct {
	EFS    *embed.FS
	FS     afero.Fs
	Config *config.SveltinConfig
	Data   *config.TemplateData
}

// NewBulma returns a pointer to a Bulma struct.
func NewBulma(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) *Bulma {
	return &Bulma{
		EFS:    efs,
		FS:     fs,
		Config: conf,
		Data:   tplData,
	}
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *Bulma) Setup(isNewProject bool) error {
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

func (cssLib *Bulma) makeSveltinStyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	bulmaFS := common.UnionMap(resources.SveltinBulmaLibFS, resources.SveltinBulmaLibStyledFS)

	// Copying the package.json config file
	sourceFile := bulmaFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs) //helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = bulmaFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = bulmaFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = bulmaFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = bulmaFS[VariablesFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "_variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = bulmaFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = bulmaFS[HeroFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = bulmaFS[FooterFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Footer.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *Bulma) makeUnstyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	bulmaFS := common.UnionMap(resources.SveltinBulmaLibFS, resources.SveltinBulmaLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := bulmaFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs) //helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = bulmaFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = bulmaFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = bulmaFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = bulmaFS[VariablesFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "_variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = bulmaFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = bulmaFS[HeroFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *Bulma) makeTheme(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	bulmaFS := common.UnionMap(resources.SveltinBulmaLibFS, resources.SveltinBulmaLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := bulmaFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs) //helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = bulmaFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = bulmaFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = bulmaFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = bulmaFS[VariablesFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "_variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = bulmaFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	return nil
}

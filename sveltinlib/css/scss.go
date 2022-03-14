/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
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
)

// Scss identifies the CSS lib to be used.
type Scss struct {
	EFS      *embed.FS
	FS       afero.Fs
	Config   *config.SveltinConfig
	Data     *config.TemplateData
	IsStyled bool
}

// NewScss returns a pointer to a Bootstrap struct.
func NewScss(isStyled bool, efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) *Scss {
	return &Scss{
		EFS:      efs,
		FS:       fs,
		Config:   conf,
		Data:     tplData,
		IsStyled: isStyled,
	}
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *Scss) Setup() error {
	if cssLib.IsStyled {
		return cssLib.makeStyled(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
	}
	return cssLib.makeUnstyled(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
}

func (cssLib *Scss) makeStyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
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
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "variables.scss")
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
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = scssFS[FooterFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Footer.svelte")
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
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "variables.scss")
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
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

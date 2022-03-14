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

// Bulma identifies the CSS lib to be used.
type Bulma struct {
	EFS      *embed.FS
	FS       afero.Fs
	Config   *config.SveltinConfig
	Data     *config.TemplateData
	IsStyled bool
}

// NewBulma returns a pointer to a Bootstrap struct.
func NewBulma(IsStyled bool, efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) *Bulma {
	return &Bulma{
		EFS:      efs,
		FS:       fs,
		Config:   conf,
		Data:     tplData,
		IsStyled: IsStyled,
	}
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *Bulma) Setup() error {
	if cssLib.IsStyled {
		return cssLib.makeStyled(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
	}
	return cssLib.makeUnstyled(cssLib.EFS, cssLib.FS, cssLib.Config, cssLib.Data)
}

func (cssLib *Bulma) makeStyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	bulmaFS := common.UnionMap(resources.SveltinBootstrabLibFS, resources.SveltinBootstrabLibStyledFS)

	// Copying the package.json config file
	sourceFile := bulmaFS["package_json"]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs) //helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	// Copying svelte.config.js file
	sourceFile = bulmaFS["svelte_config"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = bulmaFS["app_html"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = bulmaFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = bulmaFS["variables_scss"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	// Copying __layout.svelte. file
	sourceFile = bulmaFS["layout"]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	// Copying Hero.svelte component
	sourceFile = bulmaFS["hero"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = bulmaFS["footer"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Footer.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *Bulma) makeUnstyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	bulmaFS := common.UnionMap(resources.SveltinBulmaLibFS, resources.SveltinBulmaLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := bulmaFS["package_json"]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs) //helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	// Copying svelte.config.js file
	sourceFile = bulmaFS["svelte_config"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	// Copying app.html file
	sourceFile = bulmaFS["app_html"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = bulmaFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = bulmaFS["variables_scss"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "variables.scss")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	// Copying __layout.svelte. file
	sourceFile = bulmaFS["layout"]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	// Copying Hero.svelte component
	sourceFile = bulmaFS["hero"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

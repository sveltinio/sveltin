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
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/resources"
)

// Bulma identifies the CSS lib to be used.
type Bulma struct {
	CSSLib
}

func (f *Bulma) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	// Copying the package.json config file
	sourceFile := resources.SveltinBulmaCSSThemeFS["package_json"]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs) //helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = resources.SveltinBulmaCSSThemeFS["layout"]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = resources.SveltinBulmaCSSThemeFS["app_html"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = resources.SveltinBulmaCSSThemeFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = resources.SveltinBulmaCSSThemeFS["variables_scss"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "variables.scss")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = resources.SveltinBulmaCSSThemeFS["svelte_config"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = resources.SveltinBulmaCSSThemeFS["hero"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = resources.SveltinBulmaCSSThemeFS["footer"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Footer.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	return nil
}

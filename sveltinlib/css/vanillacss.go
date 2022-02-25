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

// VanillaCSS identifies the CSS lib to be used.
type VanillaCSS struct {
	CSSLib
}

func (f *VanillaCSS) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	// Copying the package.json config file
	sourceFile := resources.SveltinVanillaCSSThemeFS["package_json"]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = resources.SveltinVanillaCSSThemeFS["layout"]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = resources.SveltinVanillaCSSThemeFS["app_html"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.css file
	sourceFile = resources.SveltinVanillaCSSThemeFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.css")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = resources.SveltinVanillaCSSThemeFS["svelte_config"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = resources.SveltinVanillaCSSThemeFS["hero"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = resources.SveltinVanillaCSSThemeFS["footer"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Footer.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	return nil
}

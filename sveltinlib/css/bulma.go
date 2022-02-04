/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package css

import (
	"embed"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/resources"
)

type Bulma struct {
	CSSLib
}

func (f *Bulma) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, projectName string, themeName string) error {
	// Copying the package.json config file
	sourceFile := resources.SveltinBulmaCSSThemeFS["package_json"]
	tplConfig := helpers.NewTplConfig(sourceFile, nil, config.TemplateData{
		ProjectName: projectName,
	})
	content := helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), projectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = resources.SveltinBulmaCSSThemeFS["layout"]
	tplConfig = helpers.NewTplConfig(sourceFile, nil, config.TemplateData{
		Name: themeName,
	})
	content = helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = resources.SveltinBulmaCSSThemeFS["app_html"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "src", "app.html")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = resources.SveltinBulmaCSSThemeFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "src", "app.scss")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = resources.SveltinBulmaCSSThemeFS["variables_scss"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "src", "variables.scss")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = resources.SveltinBulmaCSSThemeFS["svelte_config"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "svelte.config.js")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = resources.SveltinBulmaCSSThemeFS["hero"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "themes", themeName, "partials", "Hero.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = resources.SveltinBulmaCSSThemeFS["footer"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "themes", themeName, "partials", "Footer.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	return nil
}

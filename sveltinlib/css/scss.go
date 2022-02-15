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

type Scss struct {
	CSSLib
}

func (f *Scss) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	// Copying the package.json config file
	sourceFile := resources.SveltinSCSSThemeFS["package_json"]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = resources.SveltinSCSSThemeFS["layout"]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = resources.SveltinSCSSThemeFS["app_html"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.scss file
	sourceFile = resources.SveltinSCSSThemeFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.scss")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying variables.scss file
	sourceFile = resources.SveltinSCSSThemeFS["variables_scss"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "variables.scss")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = resources.SveltinSCSSThemeFS["svelte_config"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Hero.svelte component
	sourceFile = resources.SveltinSCSSThemeFS["hero"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Hero.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = resources.SveltinSCSSThemeFS["footer"]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.ThemeName, "partials", "Footer.svelte")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	return nil
}

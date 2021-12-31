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

type VanillaCSS struct {
	CSSLib
}

func (f *VanillaCSS) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, projectName string) error {
	// Copying the package.json config file
	sourceFile := resources.SveltinVanillaCSSThemeFS["package_json"]
	tplConfig := helpers.NewTplConfig(sourceFile, nil, config.TemplateData{
		ProjectName: projectName,
	})
	content := helpers.ExecSveltinTpl(efs, tplConfig)
	saveAs := filepath.Join(conf.GetProjectRoot(), projectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying app.css file
	sourceFile = resources.SveltinVanillaCSSThemeFS["app_css"]
	saveAs = filepath.Join(conf.GetProjectRoot(), projectName, "src", "app.css")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	return nil
}

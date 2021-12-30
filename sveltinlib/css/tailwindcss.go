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

type TailwindCSS struct {
	CSSLib
}

func (f *TailwindCSS) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig) error {
	// Copying tailwindcss config file
	sourceFile := resources.SveltinThemeFS["tailwind_css_config"]
	saveAs := filepath.Join(conf.GetProjectRoot(), "tailwind.config.cjs")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying postcss config file
	sourceFile = resources.SveltinThemeFS["postcss"]
	saveAs = filepath.Join(conf.GetProjectRoot(), "postcss.config.cjs")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying app.css file
	sourceFile = resources.SveltinThemeFS["tailwind_css_file"]
	saveAs = filepath.Join(conf.GetSrcPath(), "app.css")
	if err := f.copyConfigFiles(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (f *TailwindCSS) runCommand(pm string) error {
	packages := []string{"postcss@latest", "postcss-load-config@latest", "autoprefixer@latest", "cssnano@latest", "tailwindcss@latest"}
	tailwindPlugins := []string{"@tailwindcss/typography@latest", "@tailwindcss/line-clamp@latest", "@tailwindcss/aspect-ratio@latest"}
	// Installing packages
	if err := helpers.RunPMCommand(pm, "addPackages", "-D", append(packages, tailwindPlugins...), true); err != nil {
		return err
	}
	return nil
}

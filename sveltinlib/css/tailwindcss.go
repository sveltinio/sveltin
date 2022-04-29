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
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// TailwindCSS identifies the CSS lib to be used.
type TailwindCSS struct {
	EFS    *embed.FS
	FS     afero.Fs
	Config *config.SveltinConfig
	Data   *config.TemplateData
}

// NewTailwindCSS returns a pointer to a TailwindCSS struct.
func NewTailwindCSS(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) *TailwindCSS {
	return &TailwindCSS{
		EFS:    efs,
		FS:     fs,
		Config: conf,
		Data:   tplData,
	}
}

// Setup is responsible to create the files to setup the CSS Lib.
func (cssLib *TailwindCSS) Setup(isNewProject bool) error {
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

func (cssLib *TailwindCSS) makeSveltinStyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	tailwindFS := common.UnionMap(resources.SveltinTailwindLibFS, resources.SveltinTailwindLibStyledFS)

	// Copying the package.json config file
	sourceFile := tailwindFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	// Copying svelte.config.js file
	sourceFile = tailwindFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	// Copying tailwindcss config file
	sourceFile = tailwindFS[TailwindConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "tailwind.config.cjs")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying postcss config file
	sourceFile = tailwindFS[PostCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "postcss.config.cjs")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = tailwindFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.css file
	sourceFile = tailwindFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.css")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = tailwindFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copyng Hero.svelte component
	sourceFile = tailwindFS[HeroFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying Footer.svelte component
	sourceFile = tailwindFS[FooterFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Footer.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *TailwindCSS) makeUnstyled(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	tailwindFS := common.UnionMap(resources.SveltinTailwindLibFS, resources.SveltinTailwindLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := tailwindFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copying svelte.config.js file
	sourceFile = tailwindFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying tailwindcss config file
	sourceFile = tailwindFS[TailwindConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "tailwind.config.cjs")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying postcss config file
	sourceFile = tailwindFS[PostCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "postcss.config.cjs")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = tailwindFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.css file
	sourceFile = tailwindFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.css")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = tailwindFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}
	// Copyng Hero.svelte component
	sourceFile = tailwindFS[HeroFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "themes", tplData.Theme.Name, "partials", "Hero.svelte")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	return nil
}

func (cssLib *TailwindCSS) makeTheme(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	tailwindFS := common.UnionMap(resources.SveltinTailwindLibFS, resources.SveltinTailwindLibUnstyledFS)

	// Copying the package.json config file
	sourceFile := tailwindFS[PackageJSONFileID]
	template := helpers.BuildTemplate(sourceFile, nil, tplData)
	content := template.Run(efs)
	saveAs := filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "package.json")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	// Copying svelte.config.js file
	sourceFile = tailwindFS[SvelteConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "svelte.config.js")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}

	// Copying tailwindcss config file
	sourceFile = tailwindFS[TailwindConfigFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "tailwind.config.cjs")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying postcss config file
	sourceFile = tailwindFS[PostCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "postcss.config.cjs")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	// Copying app.html file
	sourceFile = tailwindFS[AppHTMLFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.html")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying app.css file
	sourceFile = tailwindFS[AppCSSFileID]
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "app.css")
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, true); err != nil {
		return err
	}
	// Copying __layout.svelte. file
	sourceFile = tailwindFS[LayoutFileID]
	template = helpers.BuildTemplate(sourceFile, nil, tplData)
	content = template.Run(efs)
	saveAs = filepath.Join(conf.GetProjectRoot(), tplData.ProjectName, "src", "routes", "__layout.svelte")
	if err := helpers.WriteContentToDisk(fs, saveAs, content); err != nil {
		return err
	}

	return nil
}

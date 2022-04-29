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

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
)

// ILib ...
type ILib interface {
	Setup(bool) error
	makeSveltinStyled(*embed.FS, afero.Fs, *config.SveltinConfig, *config.TemplateData) error
	makeUnstyled(*embed.FS, afero.Fs, *config.SveltinConfig, *config.TemplateData) error
	makeTheme(*embed.FS, afero.Fs, *config.SveltinConfig, *config.TemplateData) error
}

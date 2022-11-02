/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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

// NewBootstrap returns a pointer to a CSSLib struct for Bootstrap.
func NewBootstrap(efs *embed.FS, fs afero.Fs, settings *config.SveltinSettings, tplData *config.TemplateData) *CSSLib {
	return &CSSLib{
		Name:     Bootstrap,
		EFS:      efs,
		FS:       fs,
		Settings: settings,
		TplData:  tplData,
	}
}

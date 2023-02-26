/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package css

import (
	"embed"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
)

// NewUnoCSS returns a pointer to a CSSLib struct for TUnoCSS.
func NewUnoCSS(efs *embed.FS, fs afero.Fs, settings *config.SveltinSettings, tplData *config.TemplateData) *CSSLib {
	return &CSSLib{
		Name:     UnoCSS,
		EFS:      efs,
		FS:       fs,
		Settings: settings,
		TplData:  tplData,
	}
}

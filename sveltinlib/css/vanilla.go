/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package css

import (
	"embed"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
)

type VanillaCSS struct {
	CSSLib
}

func (f *VanillaCSS) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig) error {
	return nil
}

func (f *VanillaCSS) runCommand(pm string) error {
	return nil
}

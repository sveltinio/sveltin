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
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

type Bootstrap struct {
	CSSLib
}

func (f *Bootstrap) init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, projectName string, themeName string) error {
	return sveltinerr.NewNotImplementYetError()
}

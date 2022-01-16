/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package css

import (
	"embed"
	"errors"
	"time"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

type ICSSLib interface {
	init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, projectName string, themeName string) error
}

type CSSLib struct {
	ICSSLib ICSSLib
}

func (t *CSSLib) copyConfigFiles(efs *embed.FS, fs afero.Fs, sourceFile string, saveTo string, backup bool) error {
	if backup {
		destFileExists, err := common.FileExists(fs, saveTo)
		if destFileExists && err == nil {
			fs.Rename(saveTo, saveTo+`_backup_`+time.Now().Format("2006-01-02_15:04:05"))
			fs.Remove(saveTo)
		}
	}
	err := common.CopyFileFromEmbeddedFS(efs, fs, sourceFile, saveTo)
	if err != nil {
		nErr := errors.New("something went wrong running copyConfigFiles: " + err.Error())
		return sveltinerr.NewDefaultError(nErr)
	}
	return nil
}

func (c *CSSLib) Setup(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, projectName string, themeName string) error {
	if err := c.ICSSLib.init(efs, fs, conf, projectName, themeName); err != nil {
		return err
	}
	return nil
}

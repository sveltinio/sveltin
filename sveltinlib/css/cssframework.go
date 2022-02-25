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
	"errors"
	"time"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// ICSSLib defines the methods to be implemented by each CSSLib.
type ICSSLib interface {
	init(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error
}

// CSSLib ...
type CSSLib struct {
	ICSSLib ICSSLib
}

func (c *CSSLib) copyConfigFiles(efs *embed.FS, fs afero.Fs, sourceFile string, saveTo string, backup bool) error {
	if backup {
		destFileExists, err := common.FileExists(fs, saveTo)
		if destFileExists && err == nil {
			if err := fs.Rename(saveTo, saveTo+`_backup_`+time.Now().Format("2006-01-02_15:04:05")); err != nil {
				return err
			}
			if err := fs.Remove(saveTo); err != nil {
				return err
			}
		}
	}
	err := common.CopyFileFromEmbeddedFS(efs, fs, sourceFile, saveTo)
	if err != nil {
		nErr := errors.New("something went wrong running copyConfigFiles: " + err.Error())
		return sveltinerr.NewDefaultError(nErr)
	}
	return nil
}

// Setup call the relative init method on CSSLib
func (c *CSSLib) Setup(efs *embed.FS, fs afero.Fs, conf *config.SveltinConfig, tplData *config.TemplateData) error {
	if err := c.ICSSLib.init(efs, fs, conf, tplData); err != nil {
		return err
	}
	return nil
}

/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package factory

import (
	"embed"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
)

type Artifact struct {
	efs       *embed.FS
	fs        afero.Fs
	builder   string
	resources map[string]string
	data      *config.TemplateData
}

func (sf *Artifact) GetEFS() *embed.FS {
	return sf.efs
}

func (sf *Artifact) GetFS() afero.Fs {
	return sf.fs
}

func (sf *Artifact) GetBuilder() string {
	return sf.builder
}

func (sf *Artifact) GetTemplateData() *config.TemplateData {
	return sf.data
}

func (sf *Artifact) GetResources() map[string]string {
	return sf.resources
}

func (sf *Artifact) CreateFolder(x ...string) error {
	return common.MkDir(sf.fs, filepath.Join(x...))
}

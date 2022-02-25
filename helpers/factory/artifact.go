/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package factory ...
package factory

import (
	"embed"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
)

// Artifact is the struct with all is needed to define a sveltin's artifact.
type Artifact struct {
	efs       *embed.FS
	fs        afero.Fs
	builder   string
	resources map[string]string
	data      *config.TemplateData
}

// GetEFS returns a pointer to the embedded file system used by the Artifect.
func (sf *Artifact) GetEFS() *embed.FS {
	return sf.efs
}

// GetFS returns the afero.Fs implementation used by the Artifact.
func (sf *Artifact) GetFS() afero.Fs {
	return sf.fs
}

// GetBuilder returns the builder name used by the Artifact as string.
func (sf *Artifact) GetBuilder() string {
	return sf.builder
}

// GetTemplateData returns a pointer to the TemplateData struct used by the Artifact.
func (sf *Artifact) GetTemplateData() *config.TemplateData {
	return sf.data
}

// GetResources returns the map representing the resources for the sveltin project.
func (sf *Artifact) GetResources() map[string]string {
	return sf.resources
}

// CreateFolder wraps Mkdir to create a folders structure on the file system.
func (sf *Artifact) CreateFolder(x ...string) error {
	return common.MkDir(sf.fs, filepath.Join(x...))
}

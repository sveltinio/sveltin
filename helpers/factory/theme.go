/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package factory ...
package factory

import (
	"embed"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/resources"
)

// NewThemeArtifact create an Artifact struct.
func NewThemeArtifact(efs *embed.FS, fs afero.Fs) *Artifact {
	return &Artifact{
		efs:       efs,
		fs:        fs,
		builder:   "theme",
		resources: resources.ProjectFilesMap,
	}
}

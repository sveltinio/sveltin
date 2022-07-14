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
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/resources"
)

// NewResourceArtifact create an Artifact struct.
func NewResourceArtifact(efs *embed.FS, fs afero.Fs) *Artifact {
	return &Artifact{
		efs:       efs,
		fs:        fs,
		builder:   "resource",
		resources: common.UnionMap(resources.SveltinResourceFS, common.UnionMap(resources.SveltinAPIFS, resources.SveltinMatchersFS)),
	}
}

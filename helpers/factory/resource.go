/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package factory

import (
	"embed"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/resources"
)

func NewResourceArtifact(efs *embed.FS, fs afero.Fs) *Artifact {
	return &Artifact{
		efs:       efs,
		fs:        fs,
		builder:   "resource",
		resources: resources.SveltinResourceFS,
	}
}

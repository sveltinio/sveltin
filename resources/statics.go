/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package resources

import "embed"

// SveltinStaticFS is the name for the embedded assets used by Sveltin.
//
//go:embed statics/*
var SveltinStaticFS embed.FS

// SveltinFilesFS is a map for entries in files folder.
var SveltinFilesFS = EmbeddedFSEntry{
	"mdsvex":       "statics/files/mdsvex.config.js",
	"sveltin_d_ts": "statics/files/sveltin.d.ts",
}

// SveltinImagesFS is a map for entries in images folder.
var SveltinImagesFS = EmbeddedFSEntry{
	"dummy": "statics/images/dummy.jpeg",
}

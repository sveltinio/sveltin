/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

import "github.com/sveltinio/sveltin/internal/tpltypes"

// TemplateData is the struct representing all the data to be passed to a template file.
type TemplateData struct {
	ProjectName string
	Name        string
	Config      *SveltinConfig
	NPMClient   *tpltypes.NPMClientData
	Vite        *tpltypes.ViteData
	Page        *tpltypes.PageData
	Resource    *tpltypes.ResourceData
	Content     *tpltypes.ContentData
	Metadata    *tpltypes.MetadataData
	Menu        *tpltypes.MenuData
	NoPage      *tpltypes.NoPageData
	Theme       *tpltypes.ThemeData
	Misc        *tpltypes.MiscFileData
}

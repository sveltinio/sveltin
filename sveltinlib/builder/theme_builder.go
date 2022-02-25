/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder ...
package builder

import (
	"errors"
	"text/template"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

type themeContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewThemeContentBuilder create a themeContentBuilder struct.
func NewThemeContentBuilder() *themeContentBuilder {
	return &themeContentBuilder{}
}

func (b *themeContentBuilder) setContentType() {
	b.ContentType = "theme"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *themeContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *themeContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case "readme":
		b.PathToTplFile = b.EmbeddedResources["readme"]
		return nil
	case "license":
		b.PathToTplFile = b.EmbeddedResources["license"]
		return nil
	case "theme_config":
		b.PathToTplFile = b.EmbeddedResources["theme_config"]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateId set the id for the template to be used.
func (b *themeContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template.
func (b *themeContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *themeContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"CurrentYear": func() string {
			return utils.CurrentYear()
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *themeContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder ...
package builder

import (
	"strings"
	"text/template"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/utils"
)

type menuContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewMenuContentBuilder create a menuContentBuilder struct.
func NewMenuContentBuilder() *menuContentBuilder {
	return &menuContentBuilder{}
}

func (b *menuContentBuilder) setContentType() {
	b.ContentType = "menu"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *menuContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *menuContentBuilder) setPathToTplFile() error {
	b.PathToTplFile = b.EmbeddedResources[b.TemplateId]
	return nil
}

// SetTemplateId set the id for the template to be used.
func (b *menuContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template.
func (b *menuContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *menuContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize":  strings.Title,
		"StringsJoin": strings.Join,
		"ToURL": func(txt string) string {
			return utils.ToURL(txt)
		},
		"PlusOne": func(x int) int {
			return utils.PlusOne(x)
		},
		"Sum": func(x int, y int) int {
			return utils.Sum(x, y)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *menuContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

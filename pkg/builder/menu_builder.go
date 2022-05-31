/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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

// MenuContentBuilder represents the builder for the menu.
type MenuContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewMenuContentBuilder create a menuContentBuilder struct.
func NewMenuContentBuilder() *MenuContentBuilder {
	return &MenuContentBuilder{}
}

func (b *MenuContentBuilder) setContentType() {
	b.ContentType = "menu"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *MenuContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *MenuContentBuilder) setPathToTplFile() error {
	b.PathToTplFile = b.EmbeddedResources[b.TemplateID]
	return nil
}

// SetTemplateID set the id for the template to be used.
func (b *MenuContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *MenuContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *MenuContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": func(txt string) string {
			return utils.ToTitle(txt)
		},
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
func (b *MenuContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

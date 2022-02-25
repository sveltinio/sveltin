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
	"strings"
	"text/template"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

type nopContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewNoPageContentBuilder create a nopContentBuilder struct.
func NewNoPageContentBuilder() *nopContentBuilder {
	return &nopContentBuilder{}
}

func (b *nopContentBuilder) setContentType() {
	b.ContentType = "nopage"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *nopContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *nopContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case "rss":
		b.PathToTplFile = b.EmbeddedResources["rss_static"]
		return nil
	case "sitemap":
		b.PathToTplFile = b.EmbeddedResources["sitemap_static"]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateId set the id for the template to be used.
func (b *nopContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template.
func (b *nopContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *nopContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize":  strings.Title,
		"StringsJoin": strings.Join,
		"Trimmed": func(txt string) string {
			return utils.Trimmed(txt)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *nopContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

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

// NoPContentBuilder represents the builder for the no-page artefacts (sitemap and rss).
type NoPContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewNoPageContentBuilder create a NoPContentBuilder struct.
func NewNoPageContentBuilder() *NoPContentBuilder {
	return &NoPContentBuilder{}
}

func (b *NoPContentBuilder) setContentType() {
	b.ContentType = "nopage"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *NoPContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *NoPContentBuilder) setPathToTplFile() error {
	switch b.TemplateID {
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

// SetTemplateID set the id for the template to be used.
func (b *NoPContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *NoPContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *NoPContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": func(txt string) string {
			return utils.ToTitle(txt)
		},
		"StringsJoin": strings.Join,
		"Trimmed": func(txt string) string {
			return utils.Trimmed(txt)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *NoPContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

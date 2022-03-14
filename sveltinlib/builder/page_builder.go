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

// PublicPageContentBuilder represents the builder for the public page artefact.
type PublicPageContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewPageContentBuilder create a PublicPageContentBuilder struct.
func NewPageContentBuilder() *PublicPageContentBuilder {
	return &PublicPageContentBuilder{}
}

func (b *PublicPageContentBuilder) setContentType() {
	b.ContentType = "page"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *PublicPageContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *PublicPageContentBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case Svelte:
		b.PathToTplFile = b.EmbeddedResources[Svelte]
		return nil
	case Markdown:
		b.PathToTplFile = b.EmbeddedResources[Markdown]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *PublicPageContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *PublicPageContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *PublicPageContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": strings.Title,
		"ToTitle": func(text string) string {
			return utils.ToTitle(text)
		},
		"ToVariableName": func(text string) string {
			return utils.ToVariableName(text)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *PublicPageContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

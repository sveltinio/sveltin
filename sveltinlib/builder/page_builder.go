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

const (
	// SVELTE set svelte as the language used to scaffold a new page
	SVELTE string = "svelte"
	// MARKDOWN set markdown as the language used to scaffold a new page
	MARKDOWN string = "markdown"
)

type publicPageContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewPageContentBuilder create a publicPageContentBuilder struct.
func NewPageContentBuilder() *publicPageContentBuilder {
	return &publicPageContentBuilder{}
}

func (b *publicPageContentBuilder) setContentType() {
	b.ContentType = "page"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *publicPageContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *publicPageContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case SVELTE:
		b.PathToTplFile = b.EmbeddedResources[SVELTE]
		return nil
	case MARKDOWN:
		b.PathToTplFile = b.EmbeddedResources[MARKDOWN]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateId set the id for the template to be used.
func (b *publicPageContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template.
func (b *publicPageContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *publicPageContentBuilder) setFuncs() {
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
func (b *publicPageContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

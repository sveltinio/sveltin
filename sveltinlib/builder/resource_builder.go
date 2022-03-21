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

// ResourceContentBuilder represents the builder for the resource artefact.
type ResourceContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewResourceContentBuilder create a ResourceContentBuilder struct.
func NewResourceContentBuilder() *ResourceContentBuilder {
	return &ResourceContentBuilder{}
}

func (b *ResourceContentBuilder) setContentType() {
	b.ContentType = "resource"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *ResourceContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *ResourceContentBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case API:
		b.PathToTplFile = b.EmbeddedResources[API]
		return nil
	case Index:
		b.PathToTplFile = b.EmbeddedResources[Index]
		return nil
	case IndexEndpoint:
		b.PathToTplFile = b.EmbeddedResources[IndexEndpoint]
		return nil
	case Slug:
		b.PathToTplFile = b.EmbeddedResources[Slug]
		return nil
	case SlugEndpoint:
		b.PathToTplFile = b.EmbeddedResources[SlugEndpoint]
		return nil
	case Lib:
		b.PathToTplFile = b.EmbeddedResources[Lib]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *ResourceContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *ResourceContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *ResourceContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": func(txt string) string {
			return utils.ToTitle(txt)
		},
		"ToVariableName": func(txt string) string {
			return utils.ToVariableName(txt)
		},
		"ToLibFile": func(txt string) string {
			return utils.ToLibFile(txt)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *ResourceContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

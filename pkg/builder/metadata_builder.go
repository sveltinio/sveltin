/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

// MetadataContentBuilder represents the builder for the metadata artefact.
type MetadataContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewMetadataContentBuilder create a metadataContentBuilder struct.
func NewMetadataContentBuilder() *MetadataContentBuilder {
	return &MetadataContentBuilder{}
}

func (b *MetadataContentBuilder) setContentType() {
	b.ContentType = "metadata"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *MetadataContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *MetadataContentBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case API:
		if b.TemplateData.Type == "single" {
			b.PathToTplFile = b.EmbeddedResources[APISingle]
		} else if b.TemplateData.Type == "list" {
			b.PathToTplFile = b.EmbeddedResources[APIList]
		}
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
		if b.TemplateData.Type == "single" {
			b.PathToTplFile = b.EmbeddedResources[LibSingle]
		} else if b.TemplateData.Type == "list" {
			b.PathToTplFile = b.EmbeddedResources[LibList]
		}
		return nil
	default:
		errN := errors.New("FileNotFound on the EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *MetadataContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template
func (b *MetadataContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *MetadataContentBuilder) setFuncs() {
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
		"ToSnakeCase": func(txt string) string {
			return utils.ToSnakeCase(txt)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *MetadataContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

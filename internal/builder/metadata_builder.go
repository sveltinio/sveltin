/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package builder

import (
	"errors"
	"text/template"

	"github.com/sveltinio/sveltin/config"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
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
	case ApiMetadataIndex:
		b.PathToTplFile = b.EmbeddedResources[ApiMetadataIndex]
		return nil
	case ApiFolder:
		if b.TemplateData.Metadata.Type == "single" {
			b.PathToTplFile = b.EmbeddedResources[ApiMetadataSingle]
		} else if b.TemplateData.Metadata.Type == "list" {
			b.PathToTplFile = b.EmbeddedResources[ApiMetadataList]
		}
		return nil
	case GenericMatcher:
		b.PathToTplFile = b.EmbeddedResources[GenericMatcher]
		return nil
	case Index:
		if b.TemplateData.ProjectSettings.Theme.Style == Blank {
			b.PathToTplFile = b.EmbeddedResources[IndexThemeBlank]
		} else {
			b.PathToTplFile = b.EmbeddedResources[IndexThemeSveltin]
		}
		return nil
	case IndexEndpoint:
		b.PathToTplFile = b.EmbeddedResources[IndexEndpoint]
		return nil
	case Slug:
		if b.TemplateData.ProjectSettings.Theme.Style == Blank {
			b.PathToTplFile = b.EmbeddedResources[SlugThemeBlank]
		} else {
			b.PathToTplFile = b.EmbeddedResources[SlugThemeSveltin]
		}
		return nil
	case SlugEndpoint:
		b.PathToTplFile = b.EmbeddedResources[SlugEndpoint]
		return nil
	case Lib:
		if b.TemplateData.Metadata.Type == "single" {
			b.PathToTplFile = b.EmbeddedResources[LibSingle]
		} else if b.TemplateData.Metadata.Type == "list" {
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

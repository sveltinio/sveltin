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
	"github.com/sveltinio/sveltin/internal/tpltypes"
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
	case ApiIndexFile:
		b.PathToTplFile = b.EmbeddedResources[ApiIndexFile]
		return nil
	case ApiSlugFile:
		b.PathToTplFile = b.EmbeddedResources[ApiSlugFile]
		return nil
	case StringMatcher:
		b.PathToTplFile = b.EmbeddedResources[StringMatcher]
		return nil
	case GenericMatcher:
		b.PathToTplFile = b.EmbeddedResources[GenericMatcher]
		return nil
	case Index:
		if b.TemplateData.ProjectSettings.Theme.Style == tpltypes.Blank {
			b.PathToTplFile = b.EmbeddedResources[IndexThemeBlank]
		} else {
			b.PathToTplFile = b.EmbeddedResources[IndexThemeSveltin]
		}
		return nil
	case IndexPageLoad:
		b.PathToTplFile = b.EmbeddedResources[IndexPageLoad]
		return nil
	case Slug:
		if b.TemplateData.ProjectSettings.Theme.Style == tpltypes.Blank {
			b.PathToTplFile = b.EmbeddedResources[SlugThemeBlank]
		} else {
			b.PathToTplFile = b.EmbeddedResources[SlugThemeSveltin]
		}
		return nil
	case SlugPageLoad:
		b.PathToTplFile = b.EmbeddedResources[SlugPageLoad]
		return nil
	case SlugLayout:
		b.PathToTplFile = b.EmbeddedResources[SlugLayout]
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
		"ReplaceIfNested": func(txt string) string {
			return utils.ReplaceIfNested(txt)
		},
		"ToLibFile": func(txt string) string {
			return utils.ToLibFile(txt)
		},
		"ToSlug": func(txt string) string {
			return utils.ToSlug(txt)
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

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

const (
	// API is a string for the 'api' folder.
	API string = "api"
	// INDEX is a string for the 'index' file.
	INDEX string = "index"
	// SLUG is a string for the 'slug' file.
	SLUG string = "slug"
	// LIB is a string for the 'lib' folder.
	LIB string = "lib"
)

type resourceContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewResourceContentBuilder create a resourceContentBuilder struct.
func NewResourceContentBuilder() *resourceContentBuilder {
	return &resourceContentBuilder{}
}

func (b *resourceContentBuilder) setContentType() {
	b.ContentType = "resource"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *resourceContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *resourceContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case API:
		b.PathToTplFile = b.EmbeddedResources[API]
		return nil
	case INDEX:
		b.PathToTplFile = b.EmbeddedResources[INDEX]
		return nil
	case SLUG:
		b.PathToTplFile = b.EmbeddedResources[SLUG]
		return nil
	case LIB:
		b.PathToTplFile = b.EmbeddedResources[LIB]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateId set the id for the template to be used.
func (b *resourceContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template.
func (b *resourceContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *resourceContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"ToVariableName": func(txt string) string {
			return utils.ToVariableName(txt)
		},
		"ToLibFile": func(txt string) string {
			return utils.ToLibFile(txt)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *resourceContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

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

	"github.com/gosimple/slug"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

const (
	// BLANK represents the fontmatter-only template id used when generating the content file.
	BLANK string = "blank"
	// SAMPLE represents the sample-content template id used when generating the content file.
	SAMPLE string = "sample"
)

type resContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewResContentBuilder create a resContentBuilder struct.
func NewResContentBuilder() *resContentBuilder {
	return &resContentBuilder{}
}

func (b *resContentBuilder) setContentType() {
	b.ContentType = "resContent"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *resContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *resContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case BLANK:
		b.PathToTplFile = b.EmbeddedResources[b.TemplateId]
		return nil
	case SAMPLE:
		b.PathToTplFile = b.EmbeddedResources[b.TemplateId]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateId set the id for the template to be used.
func (b *resContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template.
func (b *resContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *resContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"ToSlug": slug.Make,
		"ToTitle": func(txt string) string {
			return utils.ToTitle(txt)
		},
		"Today": func() string {
			return utils.Today()
		},
		"ToVariableName": func(txt string) string {
			return utils.ToVariableName(txt)
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *resContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

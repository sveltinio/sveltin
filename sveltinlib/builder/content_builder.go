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

// ResContentBuilder represents the builder for the content artefact.
type ResContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewResContentBuilder create a ResContentBuilder struct.
func NewResContentBuilder() *ResContentBuilder {
	return &ResContentBuilder{}
}

func (b *ResContentBuilder) setContentType() {
	b.ContentType = "resContent"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *ResContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *ResContentBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case Blank:
		b.PathToTplFile = b.EmbeddedResources[b.TemplateID]
		return nil
	case Sample:
		b.PathToTplFile = b.EmbeddedResources[b.TemplateID]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *ResContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *ResContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *ResContentBuilder) setFuncs() {
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
func (b *ResContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

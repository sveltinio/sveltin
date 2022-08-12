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
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

// ThemeBuilder represents the builder for the project.
type ThemeBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewThemeBuilder create a ThemeBuilder struct.
func NewThemeBuilder() *ThemeBuilder {
	return &ThemeBuilder{}
}

func (b *ThemeBuilder) setContentType() {
	b.ContentType = "theme"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *ThemeBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *ThemeBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case Defaults:
		b.PathToTplFile = b.EmbeddedResources[Defaults]
		return nil
	case Readme:
		b.PathToTplFile = b.EmbeddedResources[Readme]
		return nil
	case License:
		b.PathToTplFile = b.EmbeddedResources[License]
		return nil
	case ThemeConfig:
		b.PathToTplFile = b.EmbeddedResources[ThemeConfig]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *ThemeBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *ThemeBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *ThemeBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"CurrentYear": func() string {
			return utils.CurrentYear()
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *ThemeBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

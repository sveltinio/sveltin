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

// ThemeContentBuilder represents the builder for the theme artefact.
type ThemeContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewThemeContentBuilder create a ThemeContentBuilder struct.
func NewThemeContentBuilder() *ThemeContentBuilder {
	return &ThemeContentBuilder{}
}

func (b *ThemeContentBuilder) setContentType() {
	b.ContentType = "theme"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *ThemeContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *ThemeContentBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case "readme":
		b.PathToTplFile = b.EmbeddedResources["readme"]
		return nil
	case "license":
		b.PathToTplFile = b.EmbeddedResources["license"]
		return nil
	case "theme_config":
		b.PathToTplFile = b.EmbeddedResources["theme_config"]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *ThemeContentBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *ThemeContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *ThemeContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"CurrentYear": func() string {
			return utils.CurrentYear()
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *ThemeContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

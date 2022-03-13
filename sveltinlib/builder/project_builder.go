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

// ProjectBuilder represents the builder for the project.
type ProjectBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateID        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewProjectBuilder create a ProjectBuilder struct.
func NewProjectBuilder() *ProjectBuilder {
	return &ProjectBuilder{}
}

func (b *ProjectBuilder) setContentType() {
	b.ContentType = "project"
}

// SetEmbeddedResources set the map to relative embed FS.
func (b *ProjectBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *ProjectBuilder) setPathToTplFile() error {
	switch b.TemplateID {
	case Defaults:
		b.PathToTplFile = b.EmbeddedResources[Defaults]
		return nil
	case Externals:
		b.PathToTplFile = b.EmbeddedResources[Externals]
		return nil
	case Website:
		b.PathToTplFile = b.EmbeddedResources[Website]
		return nil
	case Menu:
		b.PathToTplFile = b.EmbeddedResources[InitMenu]
		return nil
	case DotEnv:
		b.PathToTplFile = b.EmbeddedResources[DotEnv]
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
	case IndexPage:
		b.PathToTplFile = b.EmbeddedResources[IndexPage]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateID set the id for the template to be used.
func (b *ProjectBuilder) SetTemplateID(id string) {
	b.TemplateID = id
}

// SetTemplateData set the data used by the template.
func (b *ProjectBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *ProjectBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"CurrentYear": func() string {
			return utils.CurrentYear()
		},
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *ProjectBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateID:    b.TemplateID,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

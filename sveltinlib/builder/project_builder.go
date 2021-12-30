/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package builder

import (
	"errors"
	"text/template"

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/utils"
)

const (
	DEFAULTS  string = "defaults"
	EXTERNALS string = "externals"
	WEBSITE   string = "website"
	MENU      string = "menu"
	INIT_MENU string = "init_menu"
	DOTENV    string = "dotenv"
)

type projectBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

func NewProjectBuilder() *projectBuilder {
	return &projectBuilder{}
}

func (b *projectBuilder) setContentType() {
	b.ContentType = "project"
}

func (b *projectBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *projectBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case DEFAULTS:
		b.PathToTplFile = b.EmbeddedResources[DEFAULTS]
		return nil
	case EXTERNALS:
		b.PathToTplFile = b.EmbeddedResources[EXTERNALS]
		return nil
	case WEBSITE:
		b.PathToTplFile = b.EmbeddedResources[WEBSITE]
		return nil
	case MENU:
		b.PathToTplFile = b.EmbeddedResources[INIT_MENU]
		return nil
	case DOTENV:
		b.PathToTplFile = b.EmbeddedResources[DOTENV]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return common.NewDefaultError(errN)
	}
}

func (b *projectBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

func (b *projectBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *projectBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"CurrentYear": func() string {
			return utils.CurrentYear()
		},
	}
}

func (b *projectBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

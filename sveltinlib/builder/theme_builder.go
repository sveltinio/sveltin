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
	README  string = "readme"
	LICENSE string = "license"
)

type themeContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

func NewThemeContentBuilder() *themeContentBuilder {
	return &themeContentBuilder{}
}

func (b *themeContentBuilder) setContentType() {
	b.ContentType = "theme"
}

func (b *themeContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *themeContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
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
		return common.NewDefaultError(errN)
	}
}

func (b *themeContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

func (b *themeContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *themeContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"CurrentYear": func() string {
			return utils.CurrentYear()
		},
	}
}

func (b *themeContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

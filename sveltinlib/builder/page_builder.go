/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package builder

import (
	"errors"
	"strings"
	"text/template"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

const (
	SVELTE   string = "svelte"
	MARKDOWN string = "markdown"
)

type publicPageContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

func NewPageContentBuilder() *publicPageContentBuilder {
	return &publicPageContentBuilder{}
}

func (b *publicPageContentBuilder) setContentType() {
	b.ContentType = "page"
}

func (b *publicPageContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *publicPageContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case SVELTE:
		b.PathToTplFile = b.EmbeddedResources[SVELTE]
		return nil
	case MARKDOWN:
		b.PathToTplFile = b.EmbeddedResources[MARKDOWN]
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return sveltinerr.NewDefaultError(errN)
	}
}

func (b *publicPageContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

func (b *publicPageContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *publicPageContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": strings.Title,
	}
}

func (b *publicPageContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

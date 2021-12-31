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
	API   string = "api"
	INDEX string = "index"
	SLUG  string = "slug"
	LIB   string = "lib"
)

type resourceContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

func NewResourceContentBuilder() *resourceContentBuilder {
	return &resourceContentBuilder{}
}

func (b *resourceContentBuilder) setContentType() {
	b.ContentType = "resource"
}

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

func (b *resourceContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

func (b *resourceContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *resourceContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": strings.Title,
	}
}

func (b *resourceContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

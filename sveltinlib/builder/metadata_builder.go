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

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
)

const (
	API_SINGLE string = "api_single"
	API_LIST   string = "api_list"
	LIB_SINGLE string = "lib_single"
	LIB_LIST   string = "lib_list"
)

type metadataContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

func NewMetadataContentBuilder() *metadataContentBuilder {
	return &metadataContentBuilder{}
}

func (b *metadataContentBuilder) setContentType() {
	b.ContentType = "metadata"
}

func (b *metadataContentBuilder) SetEmbeddedResources(res map[string]string) {
	b.EmbeddedResources = res
}

func (b *metadataContentBuilder) setPathToTplFile() error {
	switch b.TemplateId {
	case API:
		if b.TemplateData.Type == "single" {
			b.PathToTplFile = b.EmbeddedResources[API_SINGLE]
		} else if b.TemplateData.Type == "list" {
			b.PathToTplFile = b.EmbeddedResources[API_LIST]
		}
		return nil
	case INDEX:
		b.PathToTplFile = b.EmbeddedResources[INDEX]
		return nil
	case SLUG:
		b.PathToTplFile = b.EmbeddedResources[SLUG]
		return nil
	case LIB:
		if b.TemplateData.Type == "single" {
			b.PathToTplFile = b.EmbeddedResources[LIB_SINGLE]
		} else if b.TemplateData.Type == "list" {
			b.PathToTplFile = b.EmbeddedResources[LIB_LIST]
		}
		return nil
	default:
		errN := errors.New("FileNotFound on EmbeddedFS")
		return common.NewDefaultError(errN)
	}
}

func (b *metadataContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

func (b *metadataContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *metadataContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": strings.Title,
	}
}

func (b *metadataContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

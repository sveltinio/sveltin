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
	"strings"
	"text/template"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

const (
	// API_SINGLE is a string representing the api template file
	// to be used when creating a metadata of type 'single'.
	API_SINGLE string = "api_single"
	// API_LIST is a string representing the api template file
	// to be used when creating a metadata of type 'list'.
	API_LIST string = "api_list"
	// LIB_SINGLE is a string representing the lib template file
	// to be used when creating a metadata of type 'single'.
	LIB_SINGLE string = "lib_single"
	// LIB_LIST is a string representing the libe template file
	//  to be used when creating a metadata of type 'single'.
	LIB_LIST string = "lib_list"
)

type metadataContentBuilder struct {
	ContentType       string
	EmbeddedResources map[string]string
	PathToTplFile     string
	TemplateId        string
	TemplateData      *config.TemplateData
	Funcs             template.FuncMap
}

// NewMetadataContentBuilder create a metadataContentBuilder struct.
func NewMetadataContentBuilder() *metadataContentBuilder {
	return &metadataContentBuilder{}
}

func (b *metadataContentBuilder) setContentType() {
	b.ContentType = "metadata"
}

// SetEmbeddedResources set the map to relative embed FS.
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
		return sveltinerr.NewDefaultError(errN)
	}
}

// SetTemplateId set the id for the template to be used.
func (b *metadataContentBuilder) SetTemplateId(id string) {
	b.TemplateId = id
}

// SetTemplateData set the data used by the template
func (b *metadataContentBuilder) SetTemplateData(artifactData *config.TemplateData) {
	b.TemplateData = artifactData
}

func (b *metadataContentBuilder) setFuncs() {
	b.Funcs = template.FuncMap{
		"Capitalize": strings.Title,
	}
}

// GetContent returns the full Content config needed by the Builder.
func (b *metadataContentBuilder) GetContent() Content {
	return Content{
		ContentType:   b.ContentType,
		PathToTplFile: b.PathToTplFile,
		TemplateId:    b.TemplateId,
		TemplateData:  b.TemplateData,
		Funcs:         b.Funcs,
	}
}

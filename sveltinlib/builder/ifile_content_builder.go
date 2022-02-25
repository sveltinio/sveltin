/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder ...
package builder

import "github.com/sveltinio/sveltin/config"

type iFileContentBuilder interface {
	setContentType()
	SetEmbeddedResources(res map[string]string)
	SetTemplateId(string)
	SetTemplateData(artifact *config.TemplateData)
	setPathToTplFile() error
	setFuncs()
	GetContent() Content
}

// GetContentBuilder returns an concrete implementation for iFileContentBuilder.
func GetContentBuilder(contentType string) iFileContentBuilder {
	switch contentType {
	case "project":
		return NewProjectBuilder()
	case "page":
		return NewPageContentBuilder()
	case "resource":
		return NewResourceContentBuilder()
	case "metadata":
		return NewMetadataContentBuilder()
	case "resContent":
		return NewResContentBuilder()
	case "menu":
		return NewMenuContentBuilder()
	case "nopage":
		return NewNoPageContentBuilder()
	case "theme":
		return NewThemeContentBuilder()
	default:
		return nil
	}
}

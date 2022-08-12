/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder ...
package builder

import "github.com/sveltinio/sveltin/config"

// IFileContentBuilder declares building steps that are common to all types of builders.
type IFileContentBuilder interface {
	setContentType()
	SetEmbeddedResources(res map[string]string)
	SetTemplateID(string)
	SetTemplateData(artifact *config.TemplateData)
	setPathToTplFile() error
	setFuncs()
	GetContent() Content
}

// GetContentBuilder returns an concrete implementation for iFileContentBuilder.
func GetContentBuilder(contentType string) IFileContentBuilder {
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
		return NewThemeBuilder()
	default:
		return nil
	}
}

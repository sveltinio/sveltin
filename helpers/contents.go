/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers ...
package helpers

import (
	"bytes"
	"embed"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/pkg/builder"
)

// IsValidFileForContent checks is the provided FileInfo has valid
// extension (.svelte, .svx, .mdx) to be used as content file.
func IsValidFileForContent(f fs.FileInfo) bool {
	acceptedExt := []string{".svelte", ".svx", ".mdx"}

	filePrefix := strings.HasPrefix(f.Name(), "__")
	fileExt := filepath.Ext(f.Name())
	if !filePrefix && common.Contains(acceptedExt, fileExt) {
		return true
	}
	return false
}

// PrepareContent returns a builder.Content struct used by the builder director.
func PrepareContent(name string, resources map[string]string, templateID string, data *config.TemplateData) builder.Content {
	contentBuilder := builder.GetContentBuilder(name)
	contentBuilder.SetEmbeddedResources(resources)
	contentBuilder.SetTemplateID(templateID)
	contentBuilder.SetTemplateData(data)

	director := builder.NewDirector(contentBuilder)
	return director.GetContent()
}

// MakeFileContent executes the template file with all its data and functions and returns the content file as []byte
func MakeFileContent(efs *embed.FS, content builder.Content) []byte {
	template := BuildTemplate(content.PathToTplFile, content.Funcs, content.TemplateData)
	return template.Run(efs)
}

// WriteContentToDisk saves content file to the file system.
func WriteContentToDisk(fs afero.Fs, saveAs string, fileContent []byte) error {
	err := common.WriteToDisk(fs, saveAs, bytes.NewReader(fileContent))
	if err != nil {
		return err
	}
	return nil
}

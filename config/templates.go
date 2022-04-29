/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

import (
	"bytes"
	"embed"
	"path/filepath"
	"text/template"

	jww "github.com/spf13/jwalterweatherman"
)

// TplConfig is the struct representing all is needed by a template file
// (path to the template, functions map and template data).
type TplConfig struct {
	PathToTplFile string
	Funcs         template.FuncMap
	Data          TemplateData
}

// Run executes the templates and return the content as []byte.
func (tplConfig *TplConfig) Run(embedFS *embed.FS) []byte {
	pathToTplFile := tplConfig.PathToTplFile
	tplFilename := filepath.Base(tplConfig.PathToTplFile)
	funcMap := tplConfig.Funcs

	tmpl := template.Must(template.New(tplFilename).Funcs(funcMap).ParseFS(embedFS, pathToTplFile))
	var writer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&writer, tplFilename, tplConfig.Data); err != nil {
		jww.FATAL.Fatalln(err.Error())
	}
	return writer.Bytes()
}

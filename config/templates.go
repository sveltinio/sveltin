/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

import (
	"bytes"
	"embed"
	"path/filepath"
	"text/template"

	jww "github.com/spf13/jwalterweatherman"
)

type TplConfig struct {
	PathToTplFile string
	Funcs         template.FuncMap
	Data          TemplateData
}

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

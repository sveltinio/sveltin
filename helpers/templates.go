/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package helpers

import (
	"bytes"
	"embed"
	"path/filepath"
	template "text/template"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/config"
)

func NewTplConfig(tplPath string, funcs template.FuncMap, data *config.TemplateData) *config.TplConfig {
	c := new(config.TplConfig)
	c.PathToTplFile = tplPath
	c.Funcs = funcs
	c.Data = *data
	return c
}

func ExecSveltinTpl(embedFS *embed.FS, tplConfig *config.TplConfig) []byte {
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

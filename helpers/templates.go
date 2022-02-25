/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers ...
package helpers

import (
	template "text/template"

	"github.com/sveltinio/sveltin/config"
)

// BuildTemplate creates TplConfig struct with all is needed for a golang template to be executed
func BuildTemplate(tplPath string, funcs template.FuncMap, data *config.TemplateData) *config.TplConfig {
	c := new(config.TplConfig)
	c.PathToTplFile = tplPath
	c.Funcs = funcs
	c.Data = *data
	return c
}

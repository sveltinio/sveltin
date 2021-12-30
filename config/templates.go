/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

import (
	"text/template"
)

type TplConfig struct {
	PathToTplFile string
	Funcs         template.FuncMap
	Data          TemplateData
}

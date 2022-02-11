/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

type TemplateData struct {
	ProjectName string
	NPMClient   string
	Name        string
	Resource    string
	Type        string
	Config      *SveltinConfig
	Menu        *MenuConfig
	NoPage      *NoPage
	Misc        string
}

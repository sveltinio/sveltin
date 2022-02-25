/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// TemplateData is the struct representing all the data to be passed to a template file.
type TemplateData struct {
	ProjectName string
	NPMClient   string
	BaseURL     string
	PortNumber  string
	Name        string
	Resource    string
	Type        string
	Config      *SveltinConfig
	Menu        *MenuConfig
	NoPage      *NoPage
	ThemeName   string
	Misc        string
}

/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

type Pages struct {
	Content string `mapstructure:"content"`
	Index   string `mapstructure:"index"`
	Slug    string `mapstructure:"slug"`
}

type NoPage struct {
	Config *ProjectConfig
	Items  *NoPageItems
}

type NoPageItems struct {
	Resources []string
	Content   map[string][]string
	Metadata  map[string][]string
	Pages     []string
}

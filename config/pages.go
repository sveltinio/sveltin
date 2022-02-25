/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// Pages is the struct representing a public page and its content for a sveltin project.
type Pages struct {
	Content string `mapstructure:"content"`
	Index   string `mapstructure:"index"`
	Slug    string `mapstructure:"slug"`
}

// NoPage is the struct representing a no-public page (sitemap and rss) for a sveltin project.
type NoPage struct {
	Config *ProjectConfig
	Items  *NoPageItems
}

// NoPageItems is the struct representing an item
// of no-public page (sitemap and rss) for a sveltin project.
type NoPageItems struct {
	Resources []string
	Content   map[string][]string
	Metadata  map[string][]string
	Pages     []string
}

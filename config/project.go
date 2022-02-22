/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

type ProjectConfig struct {
	BaseURL           string `mapstructure:"VITE_PUBLIC_BASE_PATH"`
	SitemapChangeFreq string `mapstructure:"sitemapChangeFreq"`
	SitemapPriority   string `mapstructure:"sitemapPriority"`
}

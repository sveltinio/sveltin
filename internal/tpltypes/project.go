/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package tpltypes defines structs used to define data shared with template files.
package tpltypes

// EnvProductionData is the struct used to map the env.production file props.
type EnvProductionData struct {
	BaseURL              string `mapstructure:"VITE_PUBLIC_BASE_PATH"`
	SvelteKitBuildFolder string `mapstructure:"SVELTEKIT_BUILD_FOLDER"`
	FTPHost              string `mapstructure:"FTP_HOST"`
	FTPPort              int    `mapstructure:"FTP_PORT"`
	FTPUser              string `mapstructure:"FTP_USER"`
	FTPPassword          string `mapstructure:"FTP_PASSWORD"`
	FTPServerFolder      string `mapstructure:"FTP_SERVER_FOLDER"`
	FTPDialTimeout       int    `mapstructure:"FTP_DIAL_TIMEOUT"`
	FTPEPSVMode          bool   `mapstructure:"FTP_EPSV"`
}

// ProjectSettings is the struct used to map the sveltin.json file props.
type ProjectSettings struct {
	Name    string         `mapstructure:"name" json:"name" validate:"required"`
	BaseURL string         `mapstructure:"baseurl" json:"baseurl" validate:"required,url"`
	Theme   ThemeData      `mapstructure:"theme" json:"theme" validate:"required"`
	Sitemap SitemapData    `mapstructure:"sitemap" json:"sitemap" validate:"required"`
	Sveltin SveltinCLIData `mapstructure:"sveltin" json:"sveltin" validate:"required"`
}

// SveltinCLIData is the struct used to map the sveltin cli props.
type SveltinCLIData struct {
	Version string `mapstructure:"version" json:"version" validate:"required,semver"`
}

// SitemapData is the struct used to map the sitemap props.
type SitemapData struct {
	ChangeFreq string  `mapstructure:"changeFreq" json:"changeFreq" validate:"required,oneof='always' 'hourly' 'daily' 'weekly' 'monthly' 'yearly' 'never'"`
	Priority   float32 `mapstructure:"priority" json:"priority" validate:"required,numeric"`
}

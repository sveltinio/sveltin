/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package tpltypes ...
package tpltypes

// EnvProductionData is the struct used to map the env.production file props.
type EnvProductionData struct {
	BaseURL              string `mapstructure:"VITE_PUBLIC_BASE_PATH"`
	SitemapChangeFreq    string `mapstructure:"sitemapChangeFreq"`
	SitemapPriority      string `mapstructure:"sitemapPriority"`
	SvelteKitBuildFolder string `mapstructure:"SVELTEKIT_BUILD_FOLDER"`
	FTPHost              string `mapstructure:"FTP_HOST"`
	FTPPort              int    `mapstructure:"FTP_PORT"`
	FTPUser              string `mapstructure:"FTP_USER"`
	FTPPassword          string `mapstructure:"FTP_PASSWORD"`
	FTPServerFolder      string `mapstructure:"FTP_SERVER_FOLDER"`
	FTPDialTimeout       int    `mapstructure:"FTP_DIAL_TIMEOUT"`
	FTPEPSVMode          bool   `mapstructure:"FTP_EPSV"`
}

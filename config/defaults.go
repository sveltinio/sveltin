/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

import (
	"os"
	"path/filepath"
)

// IConfig is the interface defining the methods to be implemented.
type IConfig interface {
	GetProjectRoot() string
	GetBuildPath() string
	GetConfigPath() string
	GetContentPath() string
	GetStaticPath() string
	GetSrcPath() string
	GetAPIPath() string
	GetAPIVersion() string
	GetAPIFilename() string
	GetPublicAPIFilename() string
	GetLibPath() string
	GetRoutesPath() string
	GetThemesPath() string
	GetThemeConfigFilename() string
	GetThemeComponentsPath() string
	GetThemePartialsPath() string
	GetIndexPageFilename() string
	GetSlugPageFilename() string
	GetContentPageFilename() string
}

// SveltinConfig is the struct used the map the YAML file.
type SveltinConfig struct {
	Pages Pages `mapstructure:"pages"`
	Paths Paths `mapstructure:"paths"`
	API   API   `mapstructure:"api"`
	Theme Theme `mapstructure:"theme"`
}

// GetProjectRoot returns a string representing the current working directory.
func (c *SveltinConfig) GetProjectRoot() string {
	pwd, _ := os.Getwd()
	return pwd
}

// GetBuildPath returns a string representing the path to the 'build' folder
// relative to the current working directory.
func (c *SveltinConfig) GetBuildPath() string {
	return filepath.Join(c.GetProjectRoot(), c.Paths.Build)
}

// GetConfigPath returns a string representing the path to the 'config' folder
// relative to the current working directory.
func (c *SveltinConfig) GetConfigPath() string {
	return c.Paths.Config
}

// GetContentPath returns a string representing the path to the 'content' folder
// relative to the current working directory.
func (c *SveltinConfig) GetContentPath() string {
	return c.Paths.Content
}

// GetStaticPath returns a string representing the path to the 'static' folder
// relative to the current working directory.
func (c *SveltinConfig) GetStaticPath() string {
	return c.Paths.Static
}

// GetSrcPath returns a string representing the path to the 'src' folder
// relative to the current working directory.
func (c *SveltinConfig) GetSrcPath() string {
	return c.Paths.Src
}

// GetLibPath returns a string representing the path to the 'src/lib' folder
// relative to the current working directory.
func (c *SveltinConfig) GetLibPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Lib)
}

// GetRoutesPath returns a string representing the path to the 'src/routes' folder
// relative to the current working directory.
func (c *SveltinConfig) GetRoutesPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Routes)
}

// GetAPIPath returns a string representing the path to the 'src/routes/api' folder
// relative to the current working directory.
func (c *SveltinConfig) GetAPIPath() string {
	return filepath.Join(c.GetRoutesPath(), c.Paths.API)
}

// GetAPIVersion returns a string representing the path to the 'src/routes/api/v1'
// folder relative to the current working directory.
func (c *SveltinConfig) GetAPIVersion() string {
	return c.API.Version
}

// GetAPIFilename returns a string representing the path to the 'src/routes/api/v1/version.ts'
// folder relative to the current working directory.
func (c *SveltinConfig) GetAPIFilename() string {
	return c.API.Resource.Filename
}

// GetPublicAPIFilename returns a string representing the path to the 'src/routes/api/v1/published.ts'
// file relative to the current working directory.
func (c *SveltinConfig) GetPublicAPIFilename() string {
	return c.API.Resource.Public
}

// GetMetadataAPIFilename returns a string representing the path to the 'src/routes/api/v1/<metadata>'
// folder relative to the current working directory.
func (c *SveltinConfig) GetMetadataAPIFilename(mdName string) string {
	return mdName + ".json.ts"
}

// GetThemesPath returns a string representing the path to the 'themes' folder
// relative to the current working directory.
func (c *SveltinConfig) GetThemesPath() string {
	return filepath.Join(c.GetProjectRoot(), c.Paths.Themes)
}

// GetThemeConfigFilename returns a string representing the path to the 'themes/theme.config.js'
// file relative to the current working directory.
func (c *SveltinConfig) GetThemeConfigFilename() string {
	return c.Theme.File
}

// GetThemeComponentsPath returns a string representing the path to the 'themes/<theme>/components'
// folder relative to the current working directory.
func (c *SveltinConfig) GetThemeComponentsPath() string {
	return c.Theme.Components
}

// GetThemePartialsPath returns a string representing the path to the 'themes/<theme>/partials'
// folder relative to the current working directory.
func (c *SveltinConfig) GetThemePartialsPath() string {
	return c.Theme.Partials
}

// GetIndexPageFilename returns a string representing the path to the 'index'
// file relative to the current working directory.
func (c *SveltinConfig) GetIndexPageFilename() string {
	return c.Pages.Index
}

// GetIndexEndpointFilename returns a string representing the path to the 'index'
// file relative to the current working directory.
func (c *SveltinConfig) GetIndexEndpointFilename() string {
	return c.Pages.IndexEndpoint
}

// GetSlugPageFilename returns a string representing the path to the 'slug'
// file relative to the current working directory.
func (c *SveltinConfig) GetSlugPageFilename() string {
	return c.Pages.Slug
}

// GetSlugEndpointFilename returns a string representing the path to the 'slug'
// file relative to the current working directory.
func (c *SveltinConfig) GetSlugEndpointFilename() string {
	return c.Pages.SlugEndpoint
}

// GetContentPageFilename returns a string representing the path to the 'content' folder
// relative to the current working directory.
func (c *SveltinConfig) GetContentPageFilename() string {
	return c.Pages.Content
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config contains structs and interfaces used to map sveltin artifact to configurations.
package config

import (
	"os"
	"path/filepath"

	"github.com/sveltinio/sveltin/internal/tpltypes"
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

// SveltinSettings is the struct used the map the YAML file.
type SveltinSettings struct {
	Pages Pages          `mapstructure:"pages"`
	Paths Paths          `mapstructure:"paths"`
	API   API            `mapstructure:"api"`
	Theme tpltypes.Theme `mapstructure:"theme"`
}

// GetProjectRoot returns a string representing the current working directory.
func (c *SveltinSettings) GetProjectRoot() string {
	pwd, _ := os.Getwd()
	return pwd
}

// GetBuildPath returns a string representing the path to the 'build' folder
// relative to the current working directory.
func (c *SveltinSettings) GetBuildPath() string {
	return filepath.Join(c.GetProjectRoot(), c.Paths.Build)
}

// GetConfigPath returns a string representing the path to the 'config' folder
// relative to the current working directory.
func (c *SveltinSettings) GetConfigPath() string {
	return c.Paths.Config
}

// GetContentPath returns a string representing the path to the 'content' folder
// relative to the current working directory.
func (c *SveltinSettings) GetContentPath() string {
	return c.Paths.Content
}

// GetStaticPath returns a string representing the path to the 'static' folder
// relative to the current working directory.
func (c *SveltinSettings) GetStaticPath() string {
	return c.Paths.Static
}

// GetSrcPath returns a string representing the path to the 'src' folder
// relative to the current working directory.
func (c *SveltinSettings) GetSrcPath() string {
	return c.Paths.Src
}

// GetLibPath returns a string representing the path to the 'src/lib' folder
// relative to the current working directory.
func (c *SveltinSettings) GetLibPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Lib)
}

// GetParamsPath returns a string representing the path to the 'src/params' folder
// relative to the current working directory.
func (c *SveltinSettings) GetParamsPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Params)
}

// GetRoutesPath returns a string representing the path to the 'src/routes' folder
// relative to the current working directory.
func (c *SveltinSettings) GetRoutesPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Routes)
}

// GetAPIPath returns a string representing the path to the 'src/routes/api' folder
// relative to the current working directory.
func (c *SveltinSettings) GetAPIPath() string {
	return filepath.Join(c.GetRoutesPath(), c.Paths.API)
}

// GetAPIVersion returns a string representing the path to the 'src/routes/api/v1'
// folder relative to the current working directory.
func (c *SveltinSettings) GetAPIVersion() string {
	return c.API.Version
}

// GetAPIFilename returns a string representing the path to the 'src/routes/api/v1/<resource>/index.ts'
// folder relative to the current working directory.
func (c *SveltinSettings) GetAPIFilename() string {
	return c.API.Filename
}

// GetThemesPath returns a string representing the path to the 'themes' folder
// relative to the current working directory.
func (c *SveltinSettings) GetThemesPath() string {
	return c.Paths.Themes
}

// GetThemeConfigFilename returns a string representing the path to the 'themes/theme.config.js'
// file relative to the current working directory.
func (c *SveltinSettings) GetThemeConfigFilename() string {
	return c.Theme.File
}

// GetThemeComponentsPath returns a string representing the path to the 'themes/<theme>/components'
// folder relative to the current working directory.
func (c *SveltinSettings) GetThemeComponentsPath() string {
	return c.Theme.Components
}

// GetThemePartialsPath returns a string representing the path to the 'themes/<theme>/partials'
// folder relative to the current working directory.
func (c *SveltinSettings) GetThemePartialsPath() string {
	return c.Theme.Partials
}

// GetIndexPageFilename returns '+page.svelte' file.
func (c *SveltinSettings) GetIndexPageFilename() string {
	return c.Pages.Index
}

// GetIndexEndpointFilename returns '+page.server.ts' file.
func (c *SveltinSettings) GetIndexEndpointFilename() string {
	return c.Pages.IndexEndpoint
}

// GetSlugPageFilename returns '+page.svelte' filename for the slug.
func (c *SveltinSettings) GetSlugPageFilename() string {
	return c.Pages.Slug
}

// GetSlugEndpointFilename returns '+page.ts' as filename for the slug.
func (c *SveltinSettings) GetSlugEndpointFilename() string {
	return c.Pages.SlugEndpoint
}

// GetSlugLayoutFilename returns '+layout.svelte' as filename for the slug.
func (c *SveltinSettings) GetSlugLayoutFilename() string {
	return c.Pages.SlugLayout
}

// GetContentPageFilename returns a string representing the path to the 'content' folder
// relative to the current working directory.
func (c *SveltinSettings) GetContentPageFilename() string {
	return c.Pages.Content
}

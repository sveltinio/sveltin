/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

import (
	"os"
	"path/filepath"
)

type SettingItem struct {
	PackageManager string `yaml:"packageManager"`
}

type SveltinSettings struct {
	Item SettingItem `yaml:"Settings"`
}

func (s *SveltinSettings) GetPackageManager() string {
	return s.Item.PackageManager
}

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

type SveltinConfig struct {
	Pages Pages `mapstructure:"pages"`
	Paths Paths `mapstructure:"paths"`
	API   API   `mapstructure:"api"`
	Theme Theme `mapstructure:"theme"`
}

func (c *SveltinConfig) GetProjectRoot() string {
	pwd, _ := os.Getwd()
	return pwd
}

func (c *SveltinConfig) GetBuildPath() string {
	return filepath.Join(c.GetProjectRoot(), c.Paths.Build)
}

func (c *SveltinConfig) GetConfigPath() string {
	return c.Paths.Config
}

func (c *SveltinConfig) GetContentPath() string {
	return c.Paths.Content
}

func (c *SveltinConfig) GetStaticPath() string {
	return c.Paths.Static
}

func (c *SveltinConfig) GetSrcPath() string {
	return c.Paths.Src
}

func (c *SveltinConfig) GetLibPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Lib)
}

func (c *SveltinConfig) GetRoutesPath() string {
	return filepath.Join(c.GetSrcPath(), c.Paths.Routes)
}

func (c *SveltinConfig) GetAPIPath() string {
	return filepath.Join(c.GetRoutesPath(), c.Paths.API)
}

func (c *SveltinConfig) GetAPIVersion() string {
	return c.API.Version
}

func (c *SveltinConfig) GetAPIFilename() string {
	return c.API.Resource.Filename
}

func (c *SveltinConfig) GetPublicAPIFilename() string {
	return c.API.Resource.Public
}

func (c *SveltinConfig) GetMetadataAPIFilename() string {
	return c.API.Metadata.Filename
}

func (c *SveltinConfig) GetPublicMetadataAPIFilename() string {
	return c.API.Metadata.Public
}

func (c *SveltinConfig) GetThemesPath() string {
	return filepath.Join(c.GetProjectRoot(), c.Paths.Themes)
}

func (c *SveltinConfig) GetThemeConfigFilename() string {
	return c.Theme.File
}

func (c *SveltinConfig) GetThemeComponentsPath() string {
	return c.Theme.Components
}

func (c *SveltinConfig) GetThemePartialsPath() string {
	return c.Theme.Partials
}

func (c *SveltinConfig) GetIndexPageFilename() string {
	return c.Pages.Index
}

func (c *SveltinConfig) GetSlugPageFilename() string {
	return c.Pages.Slug
}

func (c *SveltinConfig) GetContentPageFilename() string {
	return c.Pages.Content
}

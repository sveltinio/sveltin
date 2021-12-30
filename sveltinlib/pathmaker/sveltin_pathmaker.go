/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package pathmaker

import (
	"path/filepath"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/utils"
)

type SveltinPathMaker struct {
	c *config.SveltinConfig
}

func NewSveltinPathMaker(conf *config.SveltinConfig) SveltinPathMaker {
	return SveltinPathMaker{
		c: conf,
	}
}

// ------------------ PROJECT ------------------
func (maker *SveltinPathMaker) GetProjectRoot(project string) string {
	return filepath.Join(maker.c.GetProjectRoot(), project)
}

func (maker *SveltinPathMaker) GetProjectConfigFolder(project string) string {
	return filepath.Join(maker.c.GetProjectRoot(), project, maker.c.Paths.Config)
}

func (maker *SveltinPathMaker) GetProjectContentFolder(project string) string {
	return filepath.Join(maker.c.GetProjectRoot(), project, maker.c.Paths.Content)
}

func (maker *SveltinPathMaker) GetProjectThemesFolder(project string) string {
	return filepath.Join(maker.c.GetProjectRoot(), project, maker.c.Paths.Themes)
}

// ------------------ FOLDERS ------------------
func (maker *SveltinPathMaker) GetRootFolder() string {
	return maker.c.GetProjectRoot()
}

func (maker *SveltinPathMaker) GetConfigFolder() string {
	return maker.c.GetConfigPath()
}

func (maker *SveltinPathMaker) GetContentFolder() string {
	return maker.c.GetContentPath()
}

func (maker *SveltinPathMaker) GetRoutesFolder() string {
	return maker.c.GetRoutesPath()
}

func (maker *SveltinPathMaker) GetAPIFolder() string {
	return filepath.Join(maker.c.GetAPIPath(), maker.c.GetAPIVersion())
}

func (maker *SveltinPathMaker) GetLibFolder() string {
	return maker.c.GetLibPath()
}

func (maker *SveltinPathMaker) GetStaticFolder() string {
	return maker.c.GetStaticPath()
}

func (maker *SveltinPathMaker) GetThemesFolder() string {
	return maker.c.GetThemesPath()
}

func (maker *SveltinPathMaker) GetThemeComponentsFolder() string {
	return maker.c.GetThemeComponentsPath()
}

func (maker *SveltinPathMaker) GetThemePartialsFolder() string {
	return maker.c.GetThemePartialsPath()
}

// ------------------ EXISTING RESOURCES ------------------
func (maker *SveltinPathMaker) GetPathToPublicPages() string {
	return filepath.Join(maker.c.GetProjectRoot(), maker.c.GetRoutesPath())
}

func (maker *SveltinPathMaker) GetPathToExistingResources() string {
	return filepath.Join(maker.c.GetProjectRoot(), maker.c.GetContentPath())
}

// ------------------ FILES ------------------

func (maker *SveltinPathMaker) GetResourceLibFilename(artifact string) string {
	return utils.ToLibFilename(artifact)
}

func (maker *SveltinPathMaker) GetResourceContentFilename() string {
	return maker.c.GetContentPageFilename()
}

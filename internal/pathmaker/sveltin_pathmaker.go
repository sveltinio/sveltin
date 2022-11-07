/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package pathmaker defines utility functions used to navigate through the project structure.
package pathmaker

import (
	"path/filepath"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/utils"
)

// SveltinPathMaker is the main struct to deal with path within a Sveltin/SvelteKit project structure.
type SveltinPathMaker struct {
	s *config.SveltinSettings
}

// NewSveltinPathMaker returns a SveltinPathMaker struct.
func NewSveltinPathMaker(settings *config.SveltinSettings) *SveltinPathMaker {
	return &SveltinPathMaker{
		s: settings,
	}
}

// ------------------ PROJECT ------------------

// GetProjectRoot returns a string representing the path to the sveltin project folder
// relative to the current working directory.
func (maker *SveltinPathMaker) GetProjectRoot(project string) string {
	return filepath.Join(maker.s.GetProjectRoot(), project)
}

// GetProjectConfigFolder returns a string representing the path to the 'config' folder
// for a sveltin project relative to the current working directory.
func (maker *SveltinPathMaker) GetProjectConfigFolder(project string) string {
	return filepath.Join(maker.s.GetProjectRoot(), project, maker.s.Paths.Config)
}

// GetProjectContentFolder returns a string representing the path to the 'content' folder
// for a sveltin project relative to the current working directory.
func (maker *SveltinPathMaker) GetProjectContentFolder(project string) string {
	return filepath.Join(maker.s.GetProjectRoot(), project, maker.s.Paths.Content)
}

// GetProjectThemesFolder returns a string representing the path to the 'themes' folder
// for a sveltin project relative to the current working directory.
func (maker *SveltinPathMaker) GetProjectThemesFolder(project string) string {
	return filepath.Join(maker.s.GetProjectRoot(), project, maker.s.Paths.Themes)
}

// ------------------ FOLDERS ------------------

// GetRootFolder returns a string representing the path to the project root folder.
func (maker *SveltinPathMaker) GetRootFolder() string {
	return maker.s.GetProjectRoot()
}

// GetConfigFolder returns a string representing the path to the 'config' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetConfigFolder() string {
	return maker.s.GetConfigPath()
}

// GetContentFolder returns a string representing the path to the 'content' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetContentFolder() string {
	return maker.s.GetContentPath()
}

// GetRoutesFolder returns a string representing the path to the 'src/routes' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetRoutesFolder() string {
	return maker.s.GetRoutesPath()
}

// GetLibFolder returns a string representing the path to the 'src/lib' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetLibFolder() string {
	return maker.s.GetLibPath()
}

// GetParamsFolder returns a string representing the path to the 'src/params' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetParamsFolder() string {
	return maker.s.GetParamsPath()
}

// GetAPIFolder returns a string representing the path to the 'src/routes/api' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetAPIFolder() string {
	return filepath.Join(maker.s.GetAPIPath(), maker.s.GetAPIVersion())
}

// GetStaticFolder returns a string representing the path to the 'static' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetStaticFolder() string {
	return maker.s.GetStaticPath()
}

// GetThemesFolder returns a string representing the path to the 'themes' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetThemesFolder() string {
	return maker.s.GetThemesPath()
}

// GetThemeComponentsFolder returns a string representing the path to the 'themes/<theme>/components'
// folder for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetThemeComponentsFolder() string {
	return maker.s.GetThemeComponentsPath()
}

// GetThemePartialsFolder returns a string representing the path to the 'themes/<theme>/partials'
// folder for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetThemePartialsFolder() string {
	return maker.s.GetThemePartialsPath()
}

// ------------------ EXISTING RESOURCES ------------------

// GetPathToPublicPages returns a string representing the path to the 'src/routes' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetPathToPublicPages() string {
	return filepath.Join(maker.s.GetProjectRoot(), maker.s.GetRoutesPath())
}

// GetPathToRoutes returns a string representing the path to the 'src/routes' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetPathToRoutes() string {
	return filepath.Join(maker.s.GetProjectRoot(), maker.s.GetRoutesPath())
}

// GetPathToExistingResources returns a string representing the path to the 'content' folder
// for a sveltin project relative to the project root folder.
func (maker *SveltinPathMaker) GetPathToExistingResources() string {
	return filepath.Join(maker.s.GetProjectRoot(), maker.s.GetContentPath())
}

// ------------------ FILES ------------------

// GetResourceLibFilename returns a string representing the path to the resource lib file
// for a sveltin project relative to the current working directory.
func (maker *SveltinPathMaker) GetResourceLibFilename(artifact string) string {
	return utils.ToLibFile(artifact)
}

// GetResourceContentFilename returns a string representing the filename
// for a resource content page.
func (maker *SveltinPathMaker) GetResourceContentFilename() string {
	return maker.s.GetContentPageFilename()
}

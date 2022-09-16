/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package fsm ...
package fsm

import (
	"path/filepath"
	"strings"

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/internal/composer"
	"github.com/sveltinio/sveltin/internal/pathmaker"
	"github.com/sveltinio/sveltin/internal/tpltypes"
)

// SveltinFSManager is the struct for a pathmaker.
type SveltinFSManager struct {
	maker *pathmaker.SveltinPathMaker
}

// NewSveltinFSManager returns a pointer to a SveltinFSManager struct.
func NewSveltinFSManager(maker *pathmaker.SveltinPathMaker) *SveltinFSManager {
	return &SveltinFSManager{
		maker: maker,
	}
}

// GetFolder returns a Folder struct for the provided folder name.
func (s *SveltinFSManager) GetFolder(name string) *composer.Folder {
	switch name {
	case "root":
		return composer.GetRootFolder(s.maker)
	case "config":
		return composer.GetConfigFolder(s.maker)
	case "content":
		return composer.GetContentFolder(s.maker)
	case "routes":
		return composer.GetRoutesFolder(s.maker)
	case "params":
		return composer.GetParamsFolder(s.maker)
	case "api":
		return composer.GetAPIFolder(s.maker)
	case "lib":
		return composer.GetLibFolder(s.maker)
	case "static":
		return composer.GetStaticFolder(s.maker)
	case "themes":
		return composer.GetThemesFolder(s.maker)
	default:
		return composer.NewFolder(name)
	}
}

// NewResourceContentFolder returns a pointer to the 'resource content' Folder.
func (s *SveltinFSManager) NewResourceContentFolder(name string, resource string) *composer.Folder {
	return composer.NewFolder(filepath.Join(resource, name))
}

// NewResourceContentFile returns a pointer to the 'resource content' File.
func (s *SveltinFSManager) NewResourceContentFile(name string, template string) *composer.File {
	return &composer.File{
		Name:       s.maker.GetResourceContentFilename(),
		TemplateID: template,
		TemplateData: &config.TemplateData{
			Name: name,
		},
	}
}

// NewPublicPage returns a pointer to a new 'public page' File.
func (s *SveltinFSManager) NewPublicPage(name string, language string) *composer.File {
	return &composer.File{
		Name:       helpers.PublicPageFilename(language),
		TemplateID: language,
		TemplateData: &config.TemplateData{
			Name: name,
		},
	}
}

// NewNoPageFile returns a pointer to a 'no-public page' File.
func (s *SveltinFSManager) NewNoPageFile(name string, projectConfig *tpltypes.ProjectData, resources []string, contents map[string][]string) *composer.File {
	return &composer.File{
		Name:       name + ".xml",
		TemplateID: name,
		TemplateData: &config.TemplateData{
			NoPageData: &tpltypes.NoPageData{
				Config: projectConfig,
				Items:  helpers.NewNoPageItems(resources, contents),
			},
		},
	}
}

// NewMenuFile returns a pointer to a 'no-public page' File.
func (s *SveltinFSManager) NewMenuFile(name string, projectConfig *tpltypes.ProjectData, resources []string, contents map[string][]string, withContentFlag bool) *composer.File {
	return &composer.File{
		Name:       name + ".js.ts",
		TemplateID: name,
		TemplateData: &config.TemplateData{
			MenuData: &tpltypes.MenuData{
				Items:       helpers.NewMenuItems(resources, contents),
				WithContent: withContentFlag,
			},
		},
	}
}

// NewConfigFile returns a pointer to a new 'config' File.
func (s *SveltinFSManager) NewConfigFile(projectName string, name string, cliVersion string) *composer.File {
	filename := strings.ToLower(name) + ".js.ts"
	return &composer.File{
		Name:       filename,
		TemplateID: name,
		TemplateData: &config.TemplateData{
			ProjectName: projectName,
			Name:        filename,
			Misc:        cliVersion,
		},
	}
}

// NewDotEnvFile returns a pointer to a new '.env' File.
func (s *SveltinFSManager) NewDotEnvFile(projectName string, tplData *config.TemplateData) *composer.File {
	return &composer.File{
		Name:         tplData.Name,
		TemplateID:   "dotenv",
		TemplateData: tplData,
	}
}

// NewContentFile returns a pointer to a new 'content' File.
func (s *SveltinFSManager) NewContentFile(name string, template string, resource string) *composer.File {
	return &composer.File{
		Name:       s.maker.GetResourceContentFilename(),
		TemplateID: template,
		TemplateData: &config.TemplateData{
			Name: name,
		},
	}
}

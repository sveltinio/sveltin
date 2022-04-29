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
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/pathmaker"
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
		Name:       helpers.PublicPageFilename(name, language),
		TemplateID: language,
		TemplateData: &config.TemplateData{
			Name: name,
		},
	}
}

// NewNoPage returns a pointer to a 'no-public page' File.
func (s *SveltinFSManager) NewNoPage(name string, projectConfig *config.ProjectConfig, resources []string, contents map[string][]string, metadata map[string][]string, pages []string) *composer.File {
	return &composer.File{
		Name:       name + ".xml",
		TemplateID: name,
		TemplateData: &config.TemplateData{
			NoPage: &config.NoPage{
				Config: projectConfig,
				Items:  helpers.NewNoPageItems(resources, contents, metadata, pages),
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

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package fsm ...
package fsm

import (
	"embed"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
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
func (s *SveltinFSManager) NewResourceContentFolder(contentData *tpltypes.ContentData) *composer.Folder {
	return composer.NewFolder(filepath.Join(contentData.Resource, contentData.Name))
}

// NewResourceContentFile returns a pointer to the 'resource content' File.
func (s *SveltinFSManager) NewResourceContentFile(contentData *tpltypes.ContentData) *composer.File {
	return &composer.File{
		Name:       s.maker.GetResourceContentFilename(),
		TemplateID: contentData.Type,
		TemplateData: &config.TemplateData{
			Content: contentData,
		},
	}
}

// NewPublicPageFile returns a pointer to a new 'public page' File.
func (s *SveltinFSManager) NewPublicPageFile(pageData *tpltypes.PageData, projectSettings *tpltypes.ProjectSettings) *composer.File {
	return &composer.File{
		Name:       helpers.PublicPageFilename(pageData.Type),
		TemplateID: pageData.Type,
		TemplateData: &config.TemplateData{
			Page:            pageData,
			ProjectSettings: projectSettings,
		},
	}
}

// NewNoPageFile returns a pointer to a 'no-public page' File.
func (s *SveltinFSManager) NewNoPageFile(name string, data *tpltypes.ProjectSettings, resources []string, contents map[string][]string) *composer.File {
	return &composer.File{
		Name:       name + ".xml",
		TemplateID: name,
		TemplateData: &config.TemplateData{
			NoPage: &tpltypes.NoPageData{
				Data:  data,
				Items: helpers.NewNoPageItems(resources, contents),
			},
		},
	}
}

// NewMenuFile returns a pointer to a 'no-public page' File.
func (s *SveltinFSManager) NewMenuFile(name string, resources []string, contents map[string][]string, withContentFlag bool) *composer.File {
	return &composer.File{
		Name:       name + ".js.ts",
		TemplateID: name,
		TemplateData: &config.TemplateData{
			Menu: &tpltypes.MenuData{
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
			Misc: &tpltypes.MiscFileData{
				Name: filename,
				Info: cliVersion,
			},
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

// NewJSONConfigFile returns a pointer to a new '.json' File.
func (s *SveltinFSManager) NewJSONConfigFile(tplData *config.TemplateData) *composer.File {
	return &composer.File{
		Name:         tplData.Name,
		TemplateID:   "project_settings",
		TemplateData: tplData,
	}
}

// CopyFileFromEmbed move file from embedded FS to the output folder.
func (s *SveltinFSManager) CopyFileFromEmbed(efs *embed.FS, fs afero.Fs, embeddedResourcesMap map[string]string, embeddedFileID, output string) error {
	sourceFile := embeddedResourcesMap[embeddedFileID]
	saveAs := filepath.Join(output, filepath.Base(sourceFile))
	if err := common.MoveFile(efs, fs, sourceFile, saveAs, false); err != nil {
		return err
	}
	return nil
}

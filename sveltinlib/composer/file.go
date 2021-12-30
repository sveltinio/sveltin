/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package composer

import (
	"errors"
	"path/filepath"

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
)

type File struct {
	Name         string
	path         string
	TemplateId   string
	TemplateData *config.TemplateData
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) SetName(name string) {
	f.Name = name
}

func (f *File) GetPath() string {
	return f.path
}

func (f *File) SetPath(path string) {
	f.path = path
}

func (f *File) GetTemplateId() string {
	return f.TemplateId
}

func (f *File) GetTemplateData() *config.TemplateData {
	return f.TemplateData
}

func (f *File) Create(sf *factory.Artifact) error {
	//fmt.Println(filepath.Join(f.GetPath(), f.GetName()))
	preparedContent := helpers.PrepareContent(sf.GetBuilder(), sf.GetResources(), f.GetTemplateId(), f.GetTemplateData())
	fileContent := helpers.MakeFileContent(sf.GetEFS(), preparedContent)
	saveAs := filepath.Join(f.GetPath(), f.GetName())

	if err := helpers.WriteContentToDisk(sf.GetFS(), saveAs, fileContent); err != nil {
		nErr := errors.New("something went wrong: " + err.Error())
		return common.NewDefaultError(nErr)
	}
	return nil
}

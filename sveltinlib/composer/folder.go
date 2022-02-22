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
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

type Folder struct {
	Name       string
	path       string
	components []component
}

func NewFolder(name string) *Folder {
	return &Folder{
		Name: name,
	}
}

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) SetName(name string) {
	f.Name = name
}

func (f *Folder) GetPath() string {
	return f.path
}

func (f *Folder) SetPath(path string) {
	f.path = path
}

func (f *Folder) GetComponents() []component {
	return f.components
}

func (f *Folder) Create(sf *factory.Artifact) error {
	if !common.DirExists(sf.GetFS(), f.GetPath()) {
		if err := common.MkDir(sf.GetFS(), f.GetPath()); err != nil {
			return err
		}
	}

	for _, composite := range f.GetComponents() {
		switch composite.(type) {
		case *Folder:
			composite.SetPath(filepath.Join(f.GetPath(), composite.GetName()))
		case *File:
			composite.SetPath(f.GetPath())
		default:
			errN := errors.New("composite type not valid")
			return sveltinerr.NewDefaultError(errN)
		}
		if err := composite.Create(sf); err != nil {
			return err
		}
	}
	return nil
}

func (f *Folder) Add(c component) {
	f.components = append(f.components, c)
}

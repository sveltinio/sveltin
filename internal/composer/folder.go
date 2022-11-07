/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package composer

import (
	"errors"
	"path/filepath"

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/helpers/factory"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

// Folder is the struct representing a folder on the file system.
type Folder struct {
	Name       string
	path       string
	components []Component
}

// NewFolder returns a pointer to a Folder struct.
func NewFolder(name string) *Folder {
	return &Folder{
		Name: name,
	}
}

// GetName returns the Folder name as string.
func (f *Folder) GetName() string {
	return f.Name
}

// SetName sets the folder name.
func (f *Folder) SetName(name string) {
	f.Name = name
}

// GetPath returns the Folder path as string.
func (f *Folder) GetPath() string {
	return f.path
}

// SetPath sets the Folder path.
func (f *Folder) SetPath(path string) {
	f.path = path
}

// GetComponents returns the slice of component for a Folder struct.
func (f *Folder) GetComponents() []Component {
	return f.components
}

// Create is the function to create the entire folder structure on the file system.
func (f *Folder) Create(sf *factory.Artifact) error {
	if !common.DirExists(sf.GetFS(), f.GetPath()) {
		err := common.MkDir(sf.GetFS(), f.GetPath())
		utils.IsError(err, false)
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

// Add append an item to the Folder's components list.
func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}

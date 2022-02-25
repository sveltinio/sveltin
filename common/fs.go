/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package common ...
package common

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

// MkDir is a wrapper for afero MkdirAll to create folder structure on the file system.
func MkDir(fs afero.Fs, x ...string) error {
	p := filepath.Join(x...)
	fmt.Println(p)

	if err := fs.MkdirAll(p, os.ModePerm); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// DirExists returns true if the directory/folder exists.
func DirExists(fs afero.Fs, path string) bool {
	exists, _ := afero.DirExists(fs, path)
	return exists
}

// FileExists returns true if the file exists.
func FileExists(fs afero.Fs, path string) (bool, error) {
	exists, _ := afero.Exists(fs, path)
	if !exists {
		return false, sveltinerr.NewFileNotFoundError()
	}

	isDir, _ := afero.IsDir(fs, path)
	if isDir {
		return false, sveltinerr.NewDirInsteadOfFileError()
	}

	return true, nil
}

// WriteToDisk writes content to a file on the file system. Returns error if something went wrong.
func WriteToDisk(fs afero.Fs, inpath string, r io.Reader) (err error) {
	return afero.WriteReader(fs, inpath, r)
}

// TouchFile is used to create a file on the file system.
func TouchFile(fs afero.Fs, x ...string) error {
	inpath := filepath.Join(x...)
	if err := MkDir(fs, filepath.Dir(inpath)); err != nil {
		return err
	}
	if err := WriteToDisk(fs, inpath, bytes.NewReader([]byte{})); err != nil {
		return err
	}
	return nil
}

/*
func TouchFileSingle(fs afero.Fs, name string) error {
	if err := WriteToDisk(fs, name, bytes.NewReader([]byte{})); err != nil {
		return err
	}
	return nil
}*/

// CopyFileFromEmbeddedFS copy files from the embedded file system to the actual file system.
func CopyFileFromEmbeddedFS(efs *embed.FS, fs afero.Fs, pathToFile string, saveTo string) error {
	content, err := efs.ReadFile(pathToFile)
	if err != nil {
		return sveltinerr.NewFileNotFoundError()
	}
	pathToSaveFile := filepath.Join(saveTo)
	if err := WriteToDisk(fs, pathToSaveFile, bytes.NewReader(content)); err != nil {
		return sveltinerr.NewDefaultError(err)
	}
	return nil
}

/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package common

import (
	"bytes"
	"embed"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

func MkDir(fs afero.Fs, x ...string) error {
	p := filepath.Join(x...)

	if err := fs.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func DirExists(fs afero.Fs, path string) bool {
	exists, _ := afero.DirExists(fs, path)
	return exists
}

func FileExists(fs afero.Fs, path string) (bool, error) {
	exists, _ := afero.Exists(fs, path)
	if !exists {
		return false, NewFileNotFoundError()
	}

	isDir, _ := afero.IsDir(fs, path)
	if isDir {
		return false, NewDirInsteadOfFileError()
	}

	return true, nil
}

func WriteToDisk(fs afero.Fs, inpath string, r io.Reader) (err error) {
	return afero.WriteReader(fs, inpath, r)
}

func TouchFile(fs afero.Fs, x ...string) error {
	inpath := filepath.Join(x...)
	MkDir(fs, filepath.Dir(inpath))
	if err := WriteToDisk(fs, inpath, bytes.NewReader([]byte{})); err != nil {
		return err
	}
	return nil
}

func TouchFileSingle(fs afero.Fs, name string) error {
	if err := WriteToDisk(fs, name, bytes.NewReader([]byte{})); err != nil {
		return err
	}
	return nil
}

func CopyFileFromEmbeddedFS(efs *embed.FS, fs afero.Fs, pathToFile string, saveTo string) error {
	content, err := efs.ReadFile(pathToFile)
	if err != nil {
		return NewFileNotFoundError()
	}
	pathToSaveFile := filepath.Join(saveTo)
	if err := WriteToDisk(fs, pathToSaveFile, bytes.NewReader(content)); err != nil {
		return NewDefaultError(err)
	}
	return nil
}

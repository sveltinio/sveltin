/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package common

import (
	"bufio"
	"bytes"
	"embed"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/afero"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

// MkDir is a wrapper for afero MkdirAll to create folder structure on the file system.
func MkDir(fs afero.Fs, x ...string) error {
	p := filepath.Join(x...)
	if err := fs.MkdirAll(p, os.ModePerm); err != nil {
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

// MoveFile copy files from the embedded file system to the actual file system and can backups them.
func MoveFile(efs *embed.FS, fs afero.Fs, sourceFile string, saveTo string, backup bool) error {
	if backup {
		destFileExists, err := FileExists(fs, saveTo)
		if destFileExists && err == nil {
			if err := fs.Rename(saveTo, saveTo+`_backup_`+time.Now().Format("2006-01-02_15:04:05")); err != nil {
				return err
			}
			if err := fs.Remove(saveTo); err != nil {
				return err
			}
		}
	}
	err := CopyFileFromEmbeddedFS(efs, fs, sourceFile, saveTo)
	if err != nil {
		return sveltinerr.NewMoveFileError(sourceFile, saveTo)
	}
	return nil
}

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

// ReadFileLineByLine returns a slice of strings representing lines of the input file.
func ReadFileLineByLine(appFs afero.Fs, filepath string) ([]string, error) {
	file, err := appFs.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines, nil
}

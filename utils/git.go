/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package utils

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/sveltinio/sveltin/config"
)

func GitClone(sveltinTemplate *config.AppTemplate, inpath string) {
	_, err := git.PlainClone(inpath, false, &git.CloneOptions{
		URL:      sveltinTemplate.URL,
		Progress: nil,
	})

	CheckIfError(err)
	cleanGitRepository(inpath, []string{".git"})
}

func cleanGitRepository(inpath string, foldersToRemove []string) error {
	var err error
	for _, folder := range foldersToRemove {
		err = os.RemoveAll(filepath.Join(inpath, folder))
		if err != nil {
			return err
		}
	}
	return nil
}

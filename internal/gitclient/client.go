/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package gitclient

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/utils"
)

const (
	MasterBranch = "master"
	MainBranch   = "main"
)

type Release struct {
	version    string
	commitHash string
}

// RunInit initialize an empty git repository at the given path.
func RunInit(localPath, defaultBranch string) error {
	if localPath == "" {
		_err := fmt.Errorf("%s canno be empty", localPath)
		return sveltinerr.NewDefaultError(_err)
	}

	repoPath := strings.Join([]string{localPath, git.GitDirName}, "/")
	_, err := git.PlainInit(repoPath, true)
	utils.ExitIfError(err)

	/**
	 * go-git Init and PlainInit create "master" as default branch and it is not configurable:
	 * https://github.com/go-git/go-git/blob/3f1cfde283c93f33218c807602e93d47f72f7b90/repository.go#L88
	 */
	err = renameGitBranch(localPath, MasterBranch, defaultBranch)
	utils.ExitIfError(err)

	return nil
}

// RunGitClone execute 'git clone' command.
func RunClone(repoURL, tag, inpath string) error {
	if repoURL == "" || inpath == "" {
		return sveltinerr.NewNotValidArgumentsError()
	}

	_, err := utils.NewGitHubURLParser(repoURL)
	if err != nil {
		return sveltinerr.NewNotValidGitHubRepoURL(repoURL)
	}

	r, err := git.PlainClone(inpath, false, &git.CloneOptions{
		URL: repoURL,
	})
	utils.ExitIfError(err)

	tagrefs, err := r.Tags()
	utils.ExitIfError(err)

	release := &Release{}
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		if strings.Contains(t.Name().Short(), tag) {
			release.version = t.Name().Short()
			release.commitHash = t.Hash().String()
		}
		return nil
	})
	utils.ExitIfError(err)

	w, err := r.Worktree()
	utils.ExitIfError(err)

	// if tag found, checkout it otherwise use the latest tag
	if release.commitHash != "" {
		err = w.Checkout(&git.CheckoutOptions{
			Hash: plumbing.NewHash(release.commitHash),
		})
		utils.ExitIfError(err)
	}

	// remove ".git" folder
	if err := cleanupGitRepository(inpath, git.GitDirName); err != nil {
		return err
	}

	return nil
}

func renameGitBranch(reposPath, old, new string) error {
	err := os.Chdir(reposPath)
	if err != nil {
		return err
	}

	cmdString := fmt.Sprintf("git branch -m %s %s", old, new)
	_, err = script.Exec(cmdString).Stdout()
	return err
}

func cleanupGitRepository(localPath string, f ...string) error {
	path := filepath.Join(f...)

	err := os.RemoveAll(filepath.Join(localPath, path))
	if err != nil {
		return err
	}

	return nil
}

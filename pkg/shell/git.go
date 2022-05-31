/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package shell ...
package shell

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

// GitBin is the git binary name.
const GitBin = "git"

// GitShell is a Shell implementation used to interact with a npnClient.
type GitShell struct {
	shell Shell
}

// NewGitClient returns a pointer to a NodePackageManager struct.
func NewGitClient() *GitShell {
	return &GitShell{
		shell: &LocalShell{},
	}
}

// GetShell returns a Shell.
func (s *GitShell) GetShell() Shell {
	return s.shell
}

// RunInit execute 'git init' command to initialize an empty git repository.
func (s *GitShell) RunInit(localPath string, silentMode bool) error {
	if localPath == "" {
		return sveltinerr.NewExecSystemCommandError(GitBin, "")
	}

	gitOpt := strings.Join([]string{"init", localPath}, " ")
	err := s.GetShell().Execute(GitBin, gitOpt, silentMode)
	if err != nil {
		return sveltinerr.NewExecSystemCommandError(GitBin, gitOpt)
	}

	return nil
}

// RunGitClone execute 'git clone' command.
func (s *GitShell) RunGitClone(repoURL, inpath string, silentMode bool) error {
	if repoURL == "" || inpath == "" {
		return sveltinerr.NewNotValidArgumentsError()
	}

	_, err := utils.NewGitHubURLParser(repoURL)
	if err != nil {
		return sveltinerr.NewNotValidGitHubRepoURL(repoURL)
	}
	gitOpt := strings.Join([]string{"clone", "-q", repoURL, inpath}, " ")
	err = s.GetShell().Execute(GitBin, gitOpt, silentMode)
	if err != nil {
		return sveltinerr.NewExecSystemCommandError(GitBin, gitOpt)
	}

	if err := cleanGitRepository(inpath, []string{".git"}); err != nil {
		return err
	}

	return nil
}

// RunSubmodule execute the 'git submodule add' command.
func (s *GitShell) RunSubmodule(repoURL, inpath string, silentMode bool) error {
	if repoURL == "" || inpath == "" {
		return sveltinerr.NewNotValidArgumentsError()
	}

	_, err := utils.NewGitHubURLParser(repoURL)
	if err != nil {
		return sveltinerr.NewNotValidGitHubRepoURL(repoURL)
	}
	gitOpt := strings.Join([]string{"submodule", "add", repoURL, inpath}, " ")
	err = s.GetShell().Execute(GitBin, gitOpt, silentMode)
	if err != nil {
		return sveltinerr.NewExecSystemCommandError(GitBin, gitOpt)
	}

	return nil
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

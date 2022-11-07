/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import (
	"net"
	"net/url"
	"strings"

	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

// IsValidURL returns true if the input string is a well-structured url
func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// GitHubRepo is the struct representing a GitHub repository
type GitHubRepo struct {
	host string
	user string
	repo string
}

// NewGitHubURLParser takes a github url and returns the repository info as GitHubRepo struct
func NewGitHubURLParser(input string) (*GitHubRepo, error) {
	if !isGitHubURL(input) {
		return nil, sveltinerr.NewNotValidGitHubURL(input)
	}

	parsedURL, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	f := func(c rune) bool { return c == '/' }
	splittedURL := strings.FieldsFunc(parsedURL.Path, f)
	if len(splittedURL) < 2 {
		return nil, sveltinerr.NewNotValidGitHubRepoURL(input)
	}

	repoUser := splittedURL[0]
	repoName := strings.ReplaceAll(splittedURL[1], ".git", "")

	return &GitHubRepo{
		host: "github.com",
		user: repoUser,
		repo: repoName,
	}, nil
}

// GetHost returns the host as string
func (gh *GitHubRepo) GetHost() string {
	return gh.host
}

// GetUser returns the repo's owner as string
func (gh *GitHubRepo) GetUser() string {
	return gh.user
}

// GetRepo returns the repo name as string
func (gh *GitHubRepo) GetRepo() string {
	return gh.repo
}

func isGitHubURL(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		return false
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return false
		}
	}
	return host == "github.com"
}

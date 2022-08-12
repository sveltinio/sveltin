package utils

import (
	"testing"

	"github.com/matryer/is"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
)

func TestNotGitHubURL(t *testing.T) {
	notValidURL := "https://example.com"
	_, err := NewGitHubURLParser(notValidURL)
	is := is.New(t)
	is.Equal(sveltinerr.NewNotValidGitHubURL(notValidURL), err)
}

func TestNotValidGitHubRepoURL(t *testing.T) {
	is := is.New(t)

	notValidURL := "https://github.com/sveltinio"
	_, err := NewGitHubURLParser(notValidURL)
	is.Equal(sveltinerr.NewNotValidGitHubRepoURL(notValidURL), err)

	notValidURL = "https://github.com"
	_, err = NewGitHubURLParser(notValidURL)
	is.Equal(sveltinerr.NewNotValidGitHubRepoURL(notValidURL), err)
}

func TestGitHubURLParser(t *testing.T) {
	type GHRepo struct {
		host string
		user string
		repo string
	}
	tests := []struct {
		repoURL string
		wanted  *GHRepo
	}{
		{
			repoURL: "https://github.com/sveltinio/sveltin",
			wanted: &GHRepo{
				host: "github.com",
				user: "sveltinio",
				repo: "sveltin",
			},
		},
		{
			repoURL: "https://github.com/sveltejs/kit",
			wanted: &GHRepo{
				host: "github.com",
				user: "sveltejs",
				repo: "kit",
			},
		},
	}

	for _, tc := range tests {
		is := is.New(t)
		ghParser, err := NewGitHubURLParser(tc.repoURL)
		is.NoErr(err)
		is.Equal(tc.wanted.host, ghParser.GetHost())
		is.Equal(tc.wanted.user, ghParser.GetUser())
		is.Equal(tc.wanted.repo, ghParser.GetRepo())
	}
}

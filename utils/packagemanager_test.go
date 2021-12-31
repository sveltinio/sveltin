package utils

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sveltinio/sveltin/common"
)

func TestPackageManager(t *testing.T) {
	tests := []struct {
		pm     string
		wanted bool
	}{
		{pm: "npm", wanted: true},
		{pm: "yarn", wanted: true},
		{pm: "pnpm", wanted: true},
	}

	for _, tc := range tests {
		is := is.New(t)
		exists := isCommandAvailable(tc.pm)
		is.Equal(exists, tc.wanted)
	}
}

func TestGetAvailablePackageManager(t *testing.T) {
	tests := []struct {
		pm     string
		wanted bool
	}{
		{pm: "npm", wanted: true},
		{pm: "yarn", wanted: true},
		{pm: "pnpm", wanted: true},
	}

	for _, tc := range tests {
		is := is.New(t)
		got := common.Contains(GetAvailablePackageMangerList(), tc.pm)
		is.Equal(got, tc.wanted)
	}
}

package utils

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sveltinio/sveltin/common"
)

func TestPackageManager(t *testing.T) {
	tests := []struct {
		npmClient string
		wanted    bool
	}{
		{npmClient: "npm", wanted: true},
	}

	for _, tc := range tests {
		is := is.New(t)
		exists := isCommandAvailable(tc.npmClient)
		is.Equal(exists, tc.wanted)
	}
}

func TestGetAvailablePackageManager(t *testing.T) {
	tests := []struct {
		npmClient string
		wanted    bool
	}{
		{npmClient: "npm", wanted: true},
	}

	for _, tc := range tests {
		is := is.New(t)
		got := common.Contains(GetAvailableNPMClientList(), tc.npmClient)
		is.Equal(got, tc.wanted)
	}
}

package utils

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sveltinio/sveltin/common"
)

func TestPackageManager(t *testing.T) {
	tests := []struct {
		npmClient NPMClient
		wanted    bool
	}{
		{
			npmClient: NPMClient{
				Name:    "npm",
				Version: "1",
			},
			wanted: true},
	}

	for _, tc := range tests {
		is := is.New(t)
		exists := isCommandAvailable(tc.npmClient.Name)
		is.Equal(exists, tc.wanted)
	}
}

func TestGetAvailableNPMClient(t *testing.T) {
	tests := []struct {
		npmClient NPMClient
		wanted    bool
	}{
		{npmClient: NPMClient{
			Name:    "npm",
			Version: "1",
		}, wanted: true},
	}

	for _, tc := range tests {
		is := is.New(t)
		installed := GetInstalledNPMClientList()
		got := common.Contains(GetNPMClientNames(installed), tc.npmClient.Name)
		is.Equal(got, tc.wanted)
	}
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import (
	"os/exec"

	"github.com/spf13/afero"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/npmclient"
)

// GetInstalledNPMClientList returns the list of installed npmClient as slice of NPMClientInfo.
func GetInstalledNPMClientList() []npmclient.NPMClientInfo {
	valid := []string{npmclient.Npm, npmclient.Yarn, npmclient.Pnpm}

	npmClientList := []npmclient.NPMClientInfo{}
	for _, pm := range valid {
		valid, version := GetNPMClientInfo(pm)
		if valid {
			a := npmclient.NPMClientInfo{
				Name:    pm,
				Desc:    pm,
				Version: version,
			}
			npmClientList = append(npmClientList, a)
		}
	}
	return npmClientList
}

// GetNPMClientNames returns the list of installed npmClient as slice of strings.
func GetNPMClientNames(items []npmclient.NPMClientInfo) []string {
	npmClientNames := []string{}
	for _, v := range items {
		npmClientNames = append(npmClientNames, v.Name)
	}
	return npmClientNames
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name, "-v")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// GetNPMClientInfo returns true and npm client version as string.
func GetNPMClientInfo(name string) (bool, string) {
	out, err := exec.Command(name, "-v").Output()
	if err != nil {
		return false, ""
	}
	return true, string(out)
}

func filter(in []npmclient.NPMClientInfo, name string) npmclient.NPMClientInfo {
	var out npmclient.NPMClientInfo
	for _, each := range in {
		if each.Name == name {
			out = each
		}
	}
	return out
}

// GetSelectedNPMClient returns the selected NPMClientInfo struct out of the available ones.
func GetSelectedNPMClient(in []npmclient.NPMClientInfo, name string) npmclient.NPMClientInfo {
	return filter(in, name)
}

// RetrieveProjectName returns the project name as string parsing the package.json file.
func RetrieveProjectName(appFS afero.Fs, pathToPkgJSON string) (string, error) {
	pkgFileContent, err := afero.ReadFile(appFS, pathToPkgJSON)
	ExitIfError(err)
	pkgParsed := npmclient.Parse(pkgFileContent)
	if pkgParsed.Name != "" {
		return pkgParsed.Name, nil
	}

	return "", sveltinerr.NewProjectNameNotFoundError()
}

// RetrievePackageManagerFromPkgJSON returns NPMClientInfo struct parsing the package.json file.
func RetrievePackageManagerFromPkgJSON(appFS afero.Fs, pathToPkgJSON string) (npmclient.NPMClientInfo, error) {
	pkgFileContent, err := afero.ReadFile(appFS, pathToPkgJSON)
	ExitIfError(err)
	pkgParsed := npmclient.Parse(pkgFileContent)
	if pkgParsed.PackageManager != "" {
		pmInfoString := npmclient.NPMClientInfoStr(pkgParsed.PackageManager)
		return pmInfoString.ToNPMClientInfo(), nil
	}
	return npmclient.NPMClientInfo{}, sveltinerr.NewPackageManagerKeyNotFoundOnPackageJSONFile()
}

// RetrieveCSSLib returns the css lib name used by the project.
func RetrieveCSSLib(appFS afero.Fs, pathToPkgJSON string) (string, error) {
	pkgFileContent, err := afero.ReadFile(appFS, pathToPkgJSON)
	if err != nil {
		return "", err
	}
	pkgParsed := npmclient.Parse(pkgFileContent)

	if isTailwindCSS(pkgParsed) {
		return "tailwindcss", nil
	} else if isBootstrap(pkgParsed) {
		return "bootstrap", nil
	} else if isBulma(pkgParsed) {
		return "bulma", nil
	} else if isSass(pkgParsed) {
		return "sass", nil
	} else {
		return "vanillacss", nil
	}
}

func isTailwindCSS(p *npmclient.PackageJSON) bool {
	if _, exists := p.DevDependencies["tailwindcss"]; exists {
		return true
	}
	return false
}

func isBootstrap(p *npmclient.PackageJSON) bool {
	if _, exists := p.DevDependencies["bootstrap"]; exists {
		return true
	}
	return false
}

func isBulma(p *npmclient.PackageJSON) bool {
	if _, exists := p.DevDependencies["bulma"]; exists {
		return true
	}
	return false
}

func isSass(p *npmclient.PackageJSON) bool {
	if _, exists := p.DevDependencies["tailwindcss"]; exists {
		return true
	}
	return false
}

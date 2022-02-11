/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package utils

import (
	"os/exec"
	"strings"
)

type NPMClient struct {
	Name    string
	Version string
}

func (n NPMClient) ToString() string {
	return TrimmedSuffix(n.Name + "@" + n.Version)
}

func GetInstalledNPMClientList() []NPMClient {
	valid := []string{"npm", "yarn", "pnpm"}
	npmClientList := []NPMClient{}
	for _, pm := range valid {
		valid, version := GetNPMClientInfo(pm)
		if valid {
			a := NPMClient{
				Name:    pm,
				Version: version,
			}
			npmClientList = append(npmClientList, a)
		}
	}
	return npmClientList
}

func GetNPMClientNames(items []NPMClient) []string {
	npmClientNames := []string{}
	for _, v := range items {
		npmClientNames = append(npmClientNames, v.Name)
	}
	return npmClientNames
}

func GetNPMClientName(infoStr string) string {
	return strings.Split(infoStr, "@")[0]
}

func GetNPMClient(infoStr string) NPMClient {
	splitted := strings.Split(infoStr, "@")
	return NPMClient{
		Name:    splitted[0],
		Version: splitted[1],
	}
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name, "-v")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func GetNPMClientInfo(name string) (bool, string) {
	out, err := exec.Command(name, "-v").Output()
	if err != nil {
		return false, ""
	}
	return true, string(out)
}

func filter(in []NPMClient, name string) NPMClient {
	var out NPMClient
	for _, each := range in {
		if each.Name == name {
			out = each
		}
	}
	return out
}

func GetSelectedNPMClient(in []NPMClient, name string) NPMClient {
	return filter(in, name)
}

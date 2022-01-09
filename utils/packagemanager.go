/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package utils

import (
	"os/exec"
)

func GetAvailableNPMClientList() []string {
	valid := []string{"npm", "yarn", "pnpm"}
	available := []string{}

	for _, pm := range valid {
		if isCommandAvailable(pm) {
			available = append(available, pm)
		}
	}
	return available
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name, "-v")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

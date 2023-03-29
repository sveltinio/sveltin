/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package npmclient

import "os"

/*
func setNodeForceColor_16() {
	setNodeForceColor("1")
}
*/

func setNodeForceColor_256() {
	setNodeForceColor("2")
}

/*
func setNodeForceColor_16M() {
	setNodeForceColor("3")
}
*/

func setNodeForceColor(value string) {
	checkAndSetEnvVar("FORCE_COLOR", value)
}

// checkAndSetEnvVar
func checkAndSetEnvVar(name, val string) {
	_, present := os.LookupEnv(name)
	if !present {
		os.Setenv(name, val)
	}
}

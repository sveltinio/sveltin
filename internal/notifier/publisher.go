/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package notifier implements the Observer design pattern used to notify users when a new sveltin release is available.
package notifier

// Publisher interface declares the events of interest to other objects.
type Publisher interface {
	Register(observer Observer) (bool, error)
	Deregister(observer Observer) (bool, error)
	Notify() (bool, error)
}

const (
	// InitCmd identifies the sveltin init command.
	InitCmd = "initCmd"
	// AllExceptInitCmd identifies all sveltin commands excepts the init one.
	AllExceptInitCmd = "othersCmd"
)

// MaxDaysBetween is the max number of days used to trigger a check for latest release.
const MaxDaysBetween = 7

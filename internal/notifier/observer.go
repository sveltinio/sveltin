/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package notifier

// Observer interface declares the notification interface.
type Observer interface {
	Update(string, string) (bool, error)
	GetId() string
}

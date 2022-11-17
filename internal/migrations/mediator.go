/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

// IMigrationMediator interface declares methods of communication with components.
type IMigrationMediator interface {
	canRun(IMigration) bool
	notifyAboutCompletion() error
}

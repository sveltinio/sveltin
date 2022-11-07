/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package migrations implements the Mediator design pattern used to manage migrations over sveltin versions.
package migrations

// Migration is the interface defining the methods to be implemented by single migration.
type Migration interface {
	Execute() error
	allowUp() error
}

// MigrationData is the struct with data used by migrations.
type MigrationData struct {
	PathToFile        string
	CliVersion        string
	ProjectCliVersion string
}

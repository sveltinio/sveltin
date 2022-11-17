/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import "fmt"

// Migration identifiers.
const (
	ProjectSettingsMigrationID string = "projectSettings"
	DefaultsConfigMigrationID  string = "defaultsConfig"
	ThemeConfigMigrationID     string = "themeConfig"
	DotEnvMigrationID          string = "dotenv"
	PackageJSONMigrationID     string = "packagejson"
	MDsveXMigrationID          string = "mdsvex"
)

// IMigrationFactory declares a set of methods for creating each of the abstract products.
type IMigrationFactory interface {
	MakeMigration(*MigrationManager, *MigrationServices, *MigrationData) IMigration
}

//=============================================================================

// GetMigrationFactory picks the migration factory depending on the migration id.
func GetMigrationFactory(id string) (IMigrationFactory, error) {
	switch id {
	case ProjectSettingsMigrationID:
		return &ProjectSettingsMigration{}, nil
	case DefaultsConfigMigrationID:
		return &UpdateDefaultsConfigMigration{}, nil
	case ThemeConfigMigrationID:
		return &UpdateThemeConfigMigration{}, nil
	case DotEnvMigrationID:
		return &UpdateDotEnvMigration{}, nil
	case PackageJSONMigrationID:
		return &UpdatePkgJSONMigration{}, nil
	case MDsveXMigrationID:
		return &UpdateMDsveXMigration{}, nil
	default:
		return nil, fmt.Errorf("wrong migration id used")
	}
}

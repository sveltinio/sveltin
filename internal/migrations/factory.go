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
	ProjectSettingsMigrationId string = "projectSettings"
	DefaultsConfigMigrationId  string = "defaultsConfig"
	ThemeConfigMigrationId     string = "themeConfig"
	DotEnvMigrationId          string = "dotenv"
	WebSiteTSMigrationId       string = "websitets"
	MenuTSMigrationId          string = "menuts"
	ResourceLibs               string = "resource-libs"
	SveltinDTSMigrationId      string = "sveltindts"
	PackageJSONMigrationId     string = "packagejson"
	MDsveXMigrationId          string = "mdsvex"
	SvelteConfigMigrationId    string = "svelteconfig"
	LayoutMigrationId          string = "layout"
	SvelteFilesMigrationId     string = "svelte-files"
	PageServerTSMigrationId    string = "page-server-ts"
)

var migrationMap = map[string]IMigrationFactory{
	ProjectSettingsMigrationId: &ProjectSettingsMigration{},
	DefaultsConfigMigrationId:  &UpdateDefaultsConfigMigration{},
	ThemeConfigMigrationId:     &UpdateThemeConfigMigration{},
	DotEnvMigrationId:          &UpdateDotEnvMigration{},
	WebSiteTSMigrationId:       &UpdateWebSiteTSMigration{},
	MenuTSMigrationId:          &UpdateMenuTSMigration{},
	ResourceLibs:               &UpdateResourceLibsMigration{},
	SveltinDTSMigrationId:      &SveltinDTSMigration{},
	PackageJSONMigrationId:     &UpdatePkgJSONMigration{},
	MDsveXMigrationId:          &UpdateMDsveXMigration{},
	SvelteConfigMigrationId:    &UpdateSvelteConfigMigration{},
	LayoutMigrationId:          &UpdateLayoutTSMigration{},
	SvelteFilesMigrationId:     &SvelteFilesMigration{},
	PageServerTSMigrationId:    &UpdatePageServerTSMigration{},
}

// IMigrationFactory declares a set of methods for creating each of the abstract migrations.
type IMigrationFactory interface {
	MakeMigration(*MigrationManager, *MigrationServices, *MigrationData) IMigration
}

//=============================================================================

// GetMigrationFactory picks the migration factory depending on the migration id.
func GetMigrationFactory(id string) (IMigrationFactory, error) {
	if migration, exists := migrationMap[id]; exists {
		return migration, nil
	}

	return nil, fmt.Errorf("wrong migration id: %s is not a valid migration", id)
}

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
	PackageJSONMigrationId     string = "packagejson"
	MDsveXMigrationId          string = "mdsvex"
	SvelteConfigMigrationId    string = "svelteconfig"
	LayoutMigrationId          string = "layout"
	HeadingsMigrationId        string = "headingsjs"
	WebSiteTSMigrationId       string = "websitets"
	MenuTSMigrationId          string = "menuts"
	StringsTSMigrationId       string = "stringsts"
	SveltinDTSMigrationId      string = "sveltindts"
	SvelteFilesMigrationId     string = "svelte-files"
)

// IMigrationFactory declares a set of methods for creating each of the abstract products.
type IMigrationFactory interface {
	MakeMigration(*MigrationManager, *MigrationServices, *MigrationData) IMigration
}

//=============================================================================

// GetMigrationFactory picks the migration factory depending on the migration id.
func GetMigrationFactory(id string) (IMigrationFactory, error) {
	switch id {
	case ProjectSettingsMigrationId:
		return &ProjectSettingsMigration{}, nil
	case DefaultsConfigMigrationId:
		return &UpdateDefaultsConfigMigration{}, nil
	case ThemeConfigMigrationId:
		return &UpdateThemeConfigMigration{}, nil
	case DotEnvMigrationId:
		return &UpdateDotEnvMigration{}, nil
	case PackageJSONMigrationId:
		return &UpdatePkgJSONMigration{}, nil
	case MDsveXMigrationId:
		return &UpdateMDsveXMigration{}, nil
	case SvelteConfigMigrationId:
		return &UpdateSvelteConfigMigration{}, nil
	case LayoutMigrationId:
		return &UpdateLayoutTSMigration{}, nil
	case HeadingsMigrationId:
		return &UpdateHeadingJSMigration{}, nil
	case WebSiteTSMigrationId:
		return &UpdateWebSiteTSMigration{}, nil
	case MenuTSMigrationId:
		return &UpdateMenuTSMigration{}, nil
	case StringsTSMigrationId:
		return &UpdateStringsTSMigration{}, nil
	case SveltinDTSMigrationId:
		return &SveltinDTSMigration{}, nil
	case SvelteFilesMigrationId:
		return &SvelteFilesMigration{}, nil
	default:
		return nil, fmt.Errorf("wrong migration id: %s is not a valid migration", id)
	}
}

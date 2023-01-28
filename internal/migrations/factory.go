/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import "fmt"

// Migration is the type to identidy a migration.
type Migration int

// Migration identifiers.
const (
	ProjectSettings Migration = iota
	DefaultsConfig
	WebSiteTS
	MenuTS
	SveltinDTS
	ResourceLibs
	Layout
	SvelteFiles
	PageServerTS
	SveltinioComponent
	ThemeConfig
	ThemeSveltinioComponents
	MDsveXConfig
	SvelteConfig
	DotEnv
	ViteConfig
	TSConfig
	PackageJSON
)

var migrationNameMap = map[Migration]string{
	ProjectSettings:          "project-settings",
	DefaultsConfig:           "defaults-ts",
	WebSiteTS:                "website-ts",
	MenuTS:                   "menu-ts",
	SveltinDTS:               "sveltin-dts",
	ResourceLibs:             "lib-files-ts",
	Layout:                   "layout-svelte",
	SvelteFiles:              "pages-svelte",
	PageServerTS:             "page-server-ts",
	SveltinioComponent:       "sveltinio-components",
	ThemeConfig:              "theme-config-js",
	ThemeSveltinioComponents: "theme-sveltinio-components",
	MDsveXConfig:             "mdsvex-config-js",
	SvelteConfig:             "svelte-config-js",
	DotEnv:                   "dotenv",
	ViteConfig:               "vire-config-ts",
	TSConfig:                 "ts-config-ts",
	PackageJSON:              "package-json",
}

var migrationMap = map[Migration]IMigrationFactory{
	ProjectSettings:          &ProjectSettingsMigration{},
	DefaultsConfig:           &UpdateDefaultsConfigMigration{},
	WebSiteTS:                &UpdateWebSiteTSMigration{},
	MenuTS:                   &UpdateMenuTSMigration{},
	SveltinDTS:               &SveltinDTSMigration{},
	ResourceLibs:             &UpdateResourceLibsMigration{},
	Layout:                   &UpdateLayoutTSMigration{},
	SvelteFiles:              &SvelteFilesMigration{},
	PageServerTS:             &UpdatePageServerTSMigration{},
	SveltinioComponent:       &UnhandledMigration{},
	ThemeConfig:              &UpdateThemeConfigMigration{},
	ThemeSveltinioComponents: &UnhandledMigration{},
	MDsveXConfig:             &UpdateMDsveXMigration{},
	SvelteConfig:             &UpdateSvelteConfigMigration{},
	DotEnv:                   &UpdateDotEnvMigration{},
	ViteConfig:               &UpdateViteConfigMigration{},
	TSConfig:                 &UpdateTSConfigMigration{},
	PackageJSON:              &UpdatePkgJSONMigration{},
}

// IMigrationFactory declares a set of methods for creating each of the abstract migrations.
type IMigrationFactory interface {
	MakeMigration(*MigrationManager, *MigrationServices, *MigrationData) IMigration
}

//=============================================================================

// GetMigrationFactory picks the migration factory depending on the migration id.
func GetMigrationFactory(id Migration) (IMigrationFactory, error) {
	if migration, exists := migrationMap[id]; exists {
		return migration, nil
	}

	return nil, fmt.Errorf("unknown migration id: %s", migrationNameMap[id])
}

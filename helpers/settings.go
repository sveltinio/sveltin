package helpers

import "github.com/sveltinio/sveltin/config"

func NewSveltinSettings(pmName string) config.SveltinSettings {
	return config.SveltinSettings{
		Item: config.SettingItem{
			PackageManager: pmName,
		},
	}
}

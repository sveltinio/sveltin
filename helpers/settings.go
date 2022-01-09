package helpers

import "github.com/sveltinio/sveltin/config"

func NewSveltinSettings(npmClientName string) config.SveltinSettings {
	return config.SveltinSettings{
		Item: config.SettingItem{
			NPMClient: npmClientName,
		},
	}
}

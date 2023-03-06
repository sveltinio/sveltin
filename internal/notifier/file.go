/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package notifier

import (
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/utils"
)

// SveltinJsonObserver is the concrete Observer.
type SveltinJsonObserver struct {
	Id         string
	Fs         afero.Fs
	TargetPath string
}

// Update perform the actions in response to notifications issued by the Publisher.
func (o *SveltinJsonObserver) Update(current, lastCheck string) (bool, error) {
	if o.GetId() == AllExceptInitCmd && lastCheck != "" {
		daysBetween, err := utils.DaysBetween(lastCheck, utils.TodayISO())
		if err != nil {
			return false, err
		}

		if daysBetween > MaxDaysBetween {
			content, err := afero.ReadFile(o.Fs, o.TargetPath)
			if err != nil {
				return false, err
			}

			newContent, err := utils.SetJsonStringValue(content, "sveltin.lastCheck", utils.TodayISO())
			if err != nil {
				return false, err
			}

			if err = afero.WriteFile(o.Fs, o.TargetPath, newContent, 0666); err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return true, nil
}

// GetId returns the Observer identifier.
func (o *SveltinJsonObserver) GetId() string {
	return o.Id
}

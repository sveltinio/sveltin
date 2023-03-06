/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package notifier

import (
	"github.com/sveltinio/sveltin/tui/prompts"
	"github.com/sveltinio/sveltin/utils"
)

// CommandObserver is the concrete Observer.
type CommandObserver struct {
	Id string
}

// Update perform the actions in response to notifications issued by the Publisher.
func (o *CommandObserver) Update(current, lastCheck string) (bool, error) {
	var latest string
	switch o.Id {
	case InitCmd:
		_latest, err := utils.RetrieveLatestReleaseVersion()
		if err != nil {
			return false, err
		}
		latest = _latest
	case AllExceptInitCmd:
		daysBetween, err := utils.DaysBetween(lastCheck, utils.TodayISO())
		if err != nil {
			return false, err
		}
		if daysBetween > MaxDaysBetween {
			latest, err = utils.RetrieveLatestReleaseVersion()
			if err != nil {
				return false, err
			}
		}
	}

	currentAsFloat := utils.SemVersionToFloat(current)
	latestAsFloat := utils.SemVersionToFloat(latest)
	if currentAsFloat < latestAsFloat {
		return prompts.ContinueOnUpdateAvailable(current, latest)
	}

	return true, nil
}

// GetId returns the Observer identifier.
func (o *CommandObserver) GetId() string {
	return o.Id
}

/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import "time"

const (
	layoutISO        = "2006-01-02"
	layoutMonthShort = "02-Jan-2006"
)

// Today returns the current date as formatted string "DD-ShortMonth-YYYY".
func Today() string {
	return time.Now().Format(layoutMonthShort)
}

// TodayISO returns the current UTC date as formatted string "YYYY-MM-DD".
func TodayISO() string {
	currentTime := time.Now()
	return currentTime.Format(layoutISO)
}

// DaysBetween returns the difference in days between two dates as strings.
func DaysBetween(t1, t2 string) (float64, error) {
	t1Parsed, err := time.Parse(layoutISO, t1)
	if err != nil {
		return 0, err
	}
	t2Parsed, err := time.Parse(layoutISO, t2)
	if err != nil {
		return 0, err
	}

	days := t2Parsed.Sub(t1Parsed).Hours() / 24
	return days, nil
}

// CurrentYear returns the current calendar year as a string.
func CurrentYear() string {
	return time.Now().Format("2006")
}

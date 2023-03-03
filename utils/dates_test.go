package utils

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestDatesUtils(t *testing.T) {
	is := is.New(t)

	is.Equal(time.Now().Format("02-Jan-2006"), Today())
	is.Equal("2023", CurrentYear())
	is.Equal(time.Now().Format("2006-01-02"), TodayISO())

	day1 := "2023-02-22"
	day2 := "2023-02-26"

	days, _ := DaysBetween(day1, day2)
	is.Equal(4.0, days)
}

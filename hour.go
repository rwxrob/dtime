package timefmt

import (
	"time"
)

// HourOf returns the Hour of the day passed as time rounded to the
// beginning of that hour.
func HourOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInHourOf returns the exact same minutes and secondsd but in a different hour of the same day. For beginning of hour use without SameTimeIn.
func SameTimeInHourOf(t *time.Time, hour string) *time.Time {
	// TODO
	//	nt := t.Add((time.Duration(time.Tuesday-t.Weekday()) )
	return nil
}

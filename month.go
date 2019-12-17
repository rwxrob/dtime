package timefmt

import (
	"time"
)

// JanuaryOf returns the beginning of the first day of January of the given
// year.
func JanuaryOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInJanuaryOf(t))
}

// SameTimeInJanuaryOf returns the exact same month day and time but for the
// month of January instead.
func SameTimeInJanuaryOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// MonthOf returns the beginning of the first day of the specified month of the
// given year.
func MonthOf(t *time.Time, month string) *time.Time {
	switch month {
	case "jan", "january", "Jan", "January":
		return JanuaryOf(t)
		// TODO add more
	}
	return nil
}

// SameTimeInMonthOf returns the exact same month day and time but for the
// specified month instead.
func SameTimeInMonthOf(t *time.Time, month string) *time.Time {
	switch month {
	case "jan", "january", "Jan", "January":
		return SameTimeInJanuaryOf(t)
		// TODO add more
	}
	return nil
}

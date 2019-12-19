package timefmt

import "time"

// StartOfMinute returns the start of the given minute.
func StartOfMinute(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	return &nt
}

var MinuteOf = StartOfMinute

// StartOfDay returns the start of the given hour.
func StartOfHour(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	return &nt
}

var HourOf = StartOfHour

// StartOfDay returns the start of the given day.
func StartOfDay(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return &nt
}

var DayOf = StartOfDay

// StartOfWeek returns the start of the given week.
func StartOfWeek(t *time.Time) *time.Time {
	return MondayOf(t)
}

var WeekOf = StartOfWeek

// StartOfMonth returns the start of the month.
func StartOfMonth(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return &nt
}

var MonthOf = StartOfMonth

// StartOfYear returns the start of the month.
func StartOfYear(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	return &nt
}

var YearOf = StartOfYear

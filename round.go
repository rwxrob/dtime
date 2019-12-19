package htime

import "time"

// MinuteOf returns the start of the given minute.
func MinuteOf(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	return &nt
}

// HourOf returns the start of the given hour.
func HourOf(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	return &nt
}

// DayOf returns the start of the given day.
func DayOf(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return &nt
}

// WeekOf returns the start of the given week.
func WeekOf(t *time.Time) *time.Time {
	return MondayOf(t)
}

// StartOfMonth returns the start of the month.
func MonthOf(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return &nt
}

// YearOf returns the start of the month.
func YearOf(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	return &nt
}

// Today returns the start of the current date.
func Today() *time.Time {
	now := time.Now()
	return DayOf(&now)
}

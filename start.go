package timefmt

import "time"

// StartOfMinute
func StartOfMinute(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	return &nt
}

// StartOfHour
func StartOfHour(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	return &nt
}

// StartOfDay rounds down to the beginning of the day passed as time returning
// a new time.
func StartOfDay(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return &nt
}

// StartOfWeek rounds down to the beginning of the week passed as time returning
// a new time.
func StartOfWeek(t *time.Time) *time.Time {
	return MondayOf(t)
}

// StartOfMonth returns the start of the month.
func StartOfMonth(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return &nt
}

// StartOfYear returns the start of the month.
func StartOfYear(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	return &nt
}

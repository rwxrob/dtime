package timefmt

import "time"

// StartOfDay rounds down to the beginning of the day passed as time returning
// a new time.
func StartOfDay(t *time.Time) *time.Time {
	nt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return &nt
}

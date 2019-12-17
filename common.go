package timefmt

import (
	"time"
)

// Today returns the start of the current date.
func Today() *time.Time {
	now := time.Now()
	return StartOfDay(&now)
}

// Tomorrow returns the start of tomorrow in a way that is safe from leap-year
// and daylight-savings-time changes.
// TODO

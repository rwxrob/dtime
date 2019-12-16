package timefmt

import (
	"time"
)

// Parse infers the format and returns a Time.
func Parse(s string) *time.Time {

	n := len(s)

	// no formats have only one character
	if n < 2 {
		return nil
	}

	// now

	_now := time.Now()
	now := &_now
	//year, month, day := now.Date()
	if s == "now" {
		return now
	}

	// weekday & month

	if n == 3 {
		switch s[0:3] {
		case "mon", "tue", "wed", "thu", "fri", "sat", "sun":
			return WeekdayOf(now, s[0:3])
		case "jan", "feb", "mar", "apr", "may", "jun", "jul",
			"aug", "sep", "oct", "nov", "dec":
		}
	}

	// time

	if len(s) == 2 {
		switch s[1] {
		case 'a', 'p':
			d := now.Format("2006-Jan-2 -0700")
			d += " " + s + "m"
			t, err := time.Parse("2006-Jan-2 -0700 3pm", d)
			if err == nil {
				return &t
			}
			return nil
		}
	}

	// weekday-time

	// month-name

	// month-day

	// month-day-time

	/*
		switch s[0:3] {
		case "now":
			t := time.Now()
			return &t
		case "mon", "tue", "wed", "thu", "fri", "sat", "sun":
		case "jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec":
		}
	*/
	return nil
}

package timefmt

import (
	"time"
)

// Parse infers the format and returns a Time.
func Parse(s string) *time.Time {

	// use of simple logic lighter weight (faster)
	// than regular expressions

	n := len(s)

	// no formats have only one character
	if n < 2 {
		return nil
	}

	// now

	_now := time.Now()
	now := &_now
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
			// TODO return MonthOf(now, s[0:3])
		}
	}

	// time

	if n >= 2 && n <= 4 {
		switch s[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return ParseTime(s)
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

// ParseTime specifically parses a time element from the given string. Any of
// the following formats are valid:
//
//    15
//    1504
//    3p
//    304p
//
// The rest of the time will be assumed to be the beginning of the current day
// resulting in a time on the current day in local time.
//
// Seconds are not supported (nor will they ever be) since the use case for
// them being entered by a user is extremely rare (if ever). Instead use Go's
// standard time parsing when such is needed.
func ParseTime(s string) *time.Time {
	n := len(s)

	if n < 2 {
		return nil
	}

	var t time.Time
	var err error

	today := Today()
	base, _ := time.Parse("3pm", "0am")
	last := s[n-1]

	switch {
	case last == 'a' || last == 'p':
		switch n {
		case 2, 3:
			t, err = time.Parse("3pm", s+"m")
		case 4:
			t, err = time.Parse("3:04pm", s[0:1]+":"+s[1:]+"m")
		}
	case n == 2:
		t, err = time.Parse("15", s)
	case n == 4:
		t, err = time.Parse("1504", s)
	}

	if err != nil {
		return nil
	}

	nt := today.Add(t.Sub(base))
	return &nt
}

// ParseMonth returns the beginning of the month indicated. Any of the
// following formats are valid:
//
//     jan
//     Jan
//     january
//     January
//
// The rest of the time will be assumed to be the beginning of the current year
// in local time.
func ParseMonth(s string) *time.Time {
	switch s {
	case "jan", "Jan", "january", "January":
	case "feb", "Feb", "february", "February":
	case "mar", "Mar", "march", "March":
	case "apr", "Apr", "april", "April":
	case "may", "May":
	case "jun", "Jun", "june", "June":
	case "jul", "Jul", "july", "July":
	case "aug", "Aug", "august", "August":
	case "sep", "Sep", "september", "September":
	case "oct", "Oct", "october", "October":
	case "nov", "Nov", "november", "November":
	case "dec", "Dec", "december", "December":
	}
	return nil
}

package htime

import (
	"strconv"
	"time"
)

// Parse is the main function in this package and provides a minimal format for
// entering date and time information in a practical way. It's primary use case
// is when a user needs to enter such data quickly and regularly from the
// command line or into mobile and other devices where human input speed is
// limited to what can be tapped out on the screen. The characters used in the
// formatting only characters that appear on most default keyboards without
// shifting or switching to symbolic input. The package provides no method of
// specifying timezone, which falls out of the scope of this package. All times
// are therefore assumed to be local. For a precise specification of the format
// see the timefmt.abnf time included with the package source code.
func Parse(s string) (*time.Time, *time.Duration) {
	// TODO
	return nil, nil
}

// ParseDateTime returns a time from an input string that is any of the
// following date and time combinations:
//
//     mon
//     jan
//     3p 304p 15 1504
//     15,mon
//     1504,mon
//     304p,mon
//     jan2
//     15,jan2
//     1504,jan2
//     304p,jan2
//     ,2006
//     jan,2006
//     jan2,2006
//     15,jan2,2006
//     304p,jan2,2006
//
// See Parse*() for additional format possibilities.
//
func ParseDateTime(s string) *time.Time {

	n := len(s)

	// no formats have only one character
	if n < 2 {
		return nil
	}

	_now := time.Now()
	now := &_now

	// now

	if s == "now" {
		return now
	}

	// weekday & month

	if n == 3 {
		switch s[0:3] {
		case "mon", "tue", "wed", "thu", "fri", "sat", "sun":
			return DayOfWeek(now, s[0:3])
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
		// TODO more times
	}

	// time-weekday

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

// ParseWeekday return the beginning of the weekday indicated. Any of the
// following formats are valid:
//
//     mon
//     Mon
//     monday
//     Monday
//     ...
//
// The time is assumed to be within the current week.  See WeekdayOf().
func ParseWeekday(s string) *time.Time {
	_now := time.Now()
	now := &_now
	return DayOfWeek(now, s)
}

// ParseMonth returns the beginning of the month indicated. Any of the
// following formats are valid:
//
//     jan
//     Jan
//     january
//     January
//     ...
//
// The time is assumed to be within the current week.  See MonthOf().
func ParseMonth(s string) *time.Time {
	_now := time.Now()
	now := &_now
	return MonthOfYear(now, s)
}

// ParseYear returns the beginning of the indicated year as a local time.
// The time is assumed to be within the current year.  See YearOf().
func ParseYear(s string) *time.Time {
	if len(s) < 4 {
		return nil
	}
	t, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	now := time.Now()
	nt := time.Date(t, 1, 1, 0, 0, 0, 0, now.Location())
	return &nt
}

// Package htime enables easy human input of times, dates, and durations. It
// also includes many convenience functions for rouding time duration
// boundaries as is frequently needed for scheduling and time-based search
// applications.
//
// Pointers to time.Time are used throughout the package since <nil> is usually
// a more desirable zero value than that for time.Time without a pointer.  */
package htime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	SECOND float64 = 1000000000
	MINUTE float64 = 60000000000
	HOUR   float64 = 3600000000000
	DAY    float64 = 86400000000000
	WEEK   float64 = 604800000000000
	YEAR   float64 = 31536000000000000
)

func dump(i interface{}) {
	fmt.Printf("%v\n", i)
}

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
	/*
		p := strings.IndexAny(s, "+-")

		// just a datetime
		if p < 0 {
			return ParseDateTime(s), nil
		}

		// just an offset
		if p == 0 {
			_now := time.Now()
			return &_now, ParseOffset(s)
		}

		// both datetime and offset
		dt := ParseDateTime(s[0:p])
		off := ParseOffset(s[p:])
		if dt == nil || off == nil {
			return nil, nil
		}
		return dt, off
	*/
	p := new(parser)
	p.Buffer = s
	p.Init()
	p.Parse()
	p.Execute()
	t := time.Duration(int64(p.offset))
	dump(t)
	return nil, &t
}

// ParseOffset returns a pointer to time.Duration. The input strings are
// a similar to those supported as arguments to time.ParseDuration() but
// must always begin with a mandatory plus (+) or minus (-), for example:
//
//     +10m
//     -3h
//     +10s
//
// Note that the any valid Go duration is also acceptable even though simply
// durations are the most popular and encouraged:
//
//     +1m20s
//     +1.5h
//
// The following are just aliases really and are not safe for bounding
// daylight savings and leap time:
//
//     +2d (days, one day = 24h)
//     -1w (weeks, one week = 7 days, 168h)
//
// Note that months were omitted due their large differences in number of days.
// Use weeks instead.
func ParseOffset(s string) *time.Duration {
	if len(s) < 3 || (s[0] != '-' && s[0] != '+') {
		return nil
	}
	var hours float64
	if strings.Index(s, "w") > 0 {
		if strings.Index(s, "h") > 0 {
			// TODO get the h stuff convert to h
		}
		// TODO cut out w stuff
		//TODO
		dump(hours)
	}
	if strings.Index(s, "d") >= 0 {
		// TODO
		// TODO get the d stuff convert to h
		// TODO cut out d stuff
	}
	d, _ := time.ParseDuration(s)
	return &d
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

// MonthOf returns the start of the month.
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

// JanuaryOf returns the beginning of the first day of January of the given
// year.
func JanuaryOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInJanuaryOf(t))
}

// FebruaryOf returns the beginning of the first day of February of the given
// year.
func FebruaryOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInFebruaryOf(t))
}

// MarchOf returns the beginning of the first day of March of the given
// year.
func MarchOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInMarchOf(t))
}

// AprilOf returns the beginning of the first day of April of the given
// year.
func AprilOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInAprilOf(t))
}

// MayOf returns the beginning of the first day of May of the given
// year.
func MayOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInMayOf(t))
}

// JuneOf returns the beginning of the first day of June of the given
// year.
func JuneOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInJuneOf(t))
}

// JulyOf returns the beginning of the first day of July of the given
// year.
func JulyOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInJulyOf(t))
}

// AugustOf returns the beginning of the first day of August of the given
// year.
func AugustOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInAugustOf(t))
}

// SeptemberOf returns the beginning of the first day of September of the given
// year.
func SeptemberOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInSeptemberOf(t))
}

// OctoberOf returns the beginning of the first day of October of the given
// year.
func OctoberOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInOctoberOf(t))
}

// NovemberOf returns the beginning of the first day of November of the given
// year.
func NovemberOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInNovemberOf(t))
}

// DecemberOf returns the beginning of the first day of December of the given
// year.
func DecemberOf(t *time.Time) *time.Time {
	return MonthOf(SameTimeInDecemberOf(t))
}

func samemonth(t *time.Time, month int) *time.Time {
	d := time.Date(
		t.Year(),
		time.Month(month),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		t.Location(),
	)
	return &d
}

// SameTimeInJanuaryOf returns the exact same month day and time but for the
// month of January instead.
func SameTimeInJanuaryOf(t *time.Time) *time.Time {
	return samemonth(t, 1)
}

// SameTimeInFebruaryOf returns the exact same month day and time but for the
// month of February instead.
func SameTimeInFebruaryOf(t *time.Time) *time.Time {
	return samemonth(t, 2)
}

// SameTimeInMarchOf returns the exact same month day and time but for the
// month of March instead.
func SameTimeInMarchOf(t *time.Time) *time.Time {
	return samemonth(t, 3)
}

// SameTimeInAprilOf returns the exact same month day and time but for the
// month of April instead.
func SameTimeInAprilOf(t *time.Time) *time.Time {
	return samemonth(t, 4)
}

// SameTimeInMayOf returns the exact same month day and time but for the
// month of May instead.
func SameTimeInMayOf(t *time.Time) *time.Time {
	return samemonth(t, 5)
}

// SameTimeInJuneOf returns the exact same month day and time but for the
// month of June instead.
func SameTimeInJuneOf(t *time.Time) *time.Time {
	return samemonth(t, 6)
}

// SameTimeInJulyOf returns the exact same month day and time but for the
// month of July instead.
func SameTimeInJulyOf(t *time.Time) *time.Time {
	return samemonth(t, 7)
}

// SameTimeInAugustOf returns the exact same month day and time but for the
// month of August instead.
func SameTimeInAugustOf(t *time.Time) *time.Time {
	return samemonth(t, 8)
}

// SameTimeInSeptemberOf returns the exact same month day and time but for the
// month of September instead.
func SameTimeInSeptemberOf(t *time.Time) *time.Time {
	return samemonth(t, 9)
}

// SameTimeInOctoberOf returns the exact same month day and time but for the
// month of October instead.
func SameTimeInOctoberOf(t *time.Time) *time.Time {
	return samemonth(t, 10)
}

// SameTimeInNovemberOf returns the exact same month day and time but for the
// month of November instead.
func SameTimeInNovemberOf(t *time.Time) *time.Time {
	return samemonth(t, 11)
}

// SameTimeInDecemberOf returns the exact same month day and time but for the
// month of December instead.
func SameTimeInDecemberOf(t *time.Time) *time.Time {
	return samemonth(t, 12)
}

// MonthOfYear returns the beginning of the first day of the specified month of the
// given year.
func MonthOfYear(t *time.Time, month string) *time.Time {
	switch month {
	case "jan", "january", "Jan", "January":
		return JanuaryOf(t)
	case "feb", "Feb", "february", "February":
		return FebruaryOf(t)
	case "mar", "Mar", "march", "March":
		return MarchOf(t)
	case "apr", "Apr", "april", "April":
		return AprilOf(t)
	case "may", "May":
		return MayOf(t)
	case "jun", "Jun", "june", "June":
		return JuneOf(t)
	case "jul", "Jul", "july", "July":
		return JulyOf(t)
	case "aug", "Aug", "august", "August":
		return AugustOf(t)
	case "sep", "Sep", "september", "September":
		return SeptemberOf(t)
	case "oct", "Oct", "october", "October":
		return OctoberOf(t)
	case "nov", "Nov", "november", "November":
		return NovemberOf(t)
	case "dec", "Dec", "december", "December":
		return DecemberOf(t)
	}
	return nil
}

// SameTimeInMonthOfYear returns the exact same month day and time but for the
// specified month instead.
func SameTimeInMonthOfYear(t *time.Time, month string) *time.Time {
	switch month {
	case "jan", "Jan", "january", "January":
		return SameTimeInJanuaryOf(t)
	case "feb", "Feb", "february", "February":
		return SameTimeInFebruaryOf(t)
	case "mar", "Mar", "march", "March":
		return SameTimeInMarchOf(t)
	case "apr", "Apr", "april", "April":
		return SameTimeInAprilOf(t)
	case "may", "May":
		return SameTimeInMayOf(t)
	case "jun", "Jun", "june", "June":
		return SameTimeInJuneOf(t)
	case "jul", "Jul", "july", "July":
		return SameTimeInJulyOf(t)
	case "aug", "Aug", "august", "August":
		return SameTimeInAugustOf(t)
	case "sep", "Sep", "september", "September":
		return SameTimeInSeptemberOf(t)
	case "oct", "Oct", "october", "October":
		return SameTimeInOctoberOf(t)
	case "nov", "Nov", "november", "November":
		return SameTimeInNovemberOf(t)
	case "dec", "Dec", "december", "December":
		return SameTimeInDecemberOf(t)
	}
	return nil
}

// MondayOf returns the Monday of the week passed as time rounded to the
// beginning of the day.
func MondayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnMondayOf(t))
}

// TuesdayOf returns the Tuesday of the week passed as time rounded to the
// beginning of the day.
func TuesdayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnTuesdayOf(t))
}

// WednesdayOf returns the Wednesday of the week passed as time rounded to the
// beginning of the day.
func WednesdayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnWednesdayOf(t))
}

// ThursdayOf returns the Thursday of the week passed as time rounded to the
// beginning of the day.
func ThursdayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnThursdayOf(t))
}

// FridayOf returns the Friday of the week passed as time rounded to the
// beginning of the day.
func FridayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnFridayOf(t))
}

// SaturdayOf returns the Saturday of the week passed as time rounded to the
// beginning of the day.
func SaturdayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnSaturdayOf(t))
}

// SundayOf returns the Sunday of the week passed as time rounded to the
// beginning of the day.
func SundayOf(t *time.Time) *time.Time {
	return DayOf(SameTimeOnSundayOf(t))
}

// SameTimeOnMondayOf returns the exact same time but on the Monday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnMondayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Monday-t.Weekday()) * 24 * time.Hour))
	return &nt
}

// SameTimeOnTuesdayOf returns the exact same time but on the Tuesday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnTuesdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Tuesday-t.Weekday()) * 24 * time.Hour))
	return &nt
}

// SameTimeOnWednesdayOf returns the exact same time but on the Wednesday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnWednesdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Wednesday-t.Weekday()) * 24 * time.Hour))
	return &nt
}

// SameTimeOnThursdayOf returns the exact same time but on the Thursday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnThursdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Thursday-t.Weekday()) * 24 * time.Hour))
	return &nt
}

// SameTimeOnFridayOf returns the exact same time but on the Friday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnFridayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Friday-t.Weekday()) * 24 * time.Hour))
	return &nt
}

// SameTimeOnSaturdayOf returns the exact same time but on the Saturday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnSaturdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Saturday-t.Weekday()) * 24 * time.Hour))
	return &nt
}

// SameTimeOnSundayOf returns the exact same time but on the Sunday of the week
// indicated. For beginning of day use without SameTimeOn.
func SameTimeOnSundayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(time.Sunday-t.Weekday()+7) * 24 * time.Hour))
	return &nt
}

// DayOfWeek returns the day of the week passed rounded to the beginning of the
// weekday indicated.
func DayOfWeek(t *time.Time, day string) *time.Time {
	switch day {
	case "mon", "monday", "Mon", "Monday":
		return MondayOf(t)
	case "tue", "tuesday", "Tue", "Tuesday":
		return TuesdayOf(t)
	case "wed", "wednesday", "Wed", "Wednesday":
		return WednesdayOf(t)
	case "thu", "thursday", "Thu", "Thursday":
		return ThursdayOf(t)
	case "fri", "friday", "Fri", "Friday":
		return FridayOf(t)
	case "sat", "saturday", "Sat", "Saturday":
		return SaturdayOf(t)
	case "sun", "sunday", "Sun", "Sunday":
		return SundayOf(t)
	}
	return nil
}

// SameTimeOnDayOfWeek returns the day of the week passed rounded to the
// beginning of the week day indicated.
func SameTimeOnDayOfWeek(t *time.Time, day string) *time.Time {
	switch day {
	case "mon", "monday", "Mon", "Monday":
		return SameTimeOnMondayOf(t)
	case "tue", "tuesday", "Tue", "Tuesday":
		return SameTimeOnTuesdayOf(t)
	case "wed", "wednesday", "Wed", "Wednesday":
		return SameTimeOnWednesdayOf(t)
	case "thu", "thursday", "Thu", "Thursday":
		return SameTimeOnThursdayOf(t)
	case "fri", "friday", "Fri", "Friday":
		return SameTimeOnFridayOf(t)
	case "sat", "saturday", "Sat", "Saturday":
		return SameTimeOnSaturdayOf(t)
	case "sun", "sunday", "Sun", "Sunday":
		return SameTimeOnSundayOf(t)
	}
	return nil
}

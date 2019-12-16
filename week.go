package timefmt

import (
	"time"
)

// MondayOf returns the Monday of the week passed as time rounded to the
// beginning of the day.
func MondayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeMondayOf(t))
}

// TuesdayOf returns the Tuesday of the week passed as time rounded to the
// beginning of the day.
func TuesdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeTuesdayOf(t))
}

// WednesdayOf returns the Wednesday of the week passed as time rounded to the
// beginning of the day.
func WednesdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeWednesdayOf(t))
}

// ThursdayOf returns the Thursday of the week passed as time rounded to the
// beginning of the day.
func ThursdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeThursdayOf(t))
}

// FridayOf returns the Friday of the week passed as time rounded to the
// beginning of the day.
func FridayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeFridayOf(t))
}

// SaturdayOf returns the Saturday of the week passed as time rounded to the
// beginning of the day.
func SaturdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeSaturdayOf(t))
}

// SundayOf returns the Sunday of the week passed as time rounded to the
// beginning of the day.
func SundayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeSundayOf(t))
}

// SameTimeMondayOf returns the exact same time but on the Monday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeMondayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Monday) * 24 * time.Hour))
	return &nt
}

// SameTimeTuesdayOf returns the exact same time but on the Tuesday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeTuesdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Tuesday) * 24 * time.Hour))
	return &nt
}

// SameTimeWednesdayOf returns the exact same time but on the Wednesday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeWednesdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Wednesday) * 24 * time.Hour))
	return &nt
}

// SameTimeThursdayOf returns the exact same time but on the Thursday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeThursdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Thursday) * 24 * time.Hour))
	return &nt
}

// SameTimeFridayOf returns the exact same time but on the Friday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeFridayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Friday) * 24 * time.Hour))
	return &nt
}

// SameTimeSaturdayOf returns the exact same time but on the Saturday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeSaturdayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Saturday) * 24 * time.Hour))
	return &nt
}

// SameTimeSundayOf returns the exact same time but on the Sunday of the week
// indicated. For beginning of day use without SameTime.
func SameTimeSundayOf(t *time.Time) *time.Time {
	nt := t.Add((time.Duration(t.Weekday()-time.Sunday) * 24 * time.Hour))
	return &nt
}

// WeekdayOf returns the day of the week passed rounded to the beginning of the
// week day indicated.
func WeekdayOf(t *time.Time, day string) *time.Time {
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

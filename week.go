package timefmt

import (
	"time"
)

// MondayOf returns the Monday of the week passed as time rounded to the
// beginning of the day.
func MondayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnMondayOf(t))
}

// TuesdayOf returns the Tuesday of the week passed as time rounded to the
// beginning of the day.
func TuesdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnTuesdayOf(t))
}

// WednesdayOf returns the Wednesday of the week passed as time rounded to the
// beginning of the day.
func WednesdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnWednesdayOf(t))
}

// ThursdayOf returns the Thursday of the week passed as time rounded to the
// beginning of the day.
func ThursdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnThursdayOf(t))
}

// FridayOf returns the Friday of the week passed as time rounded to the
// beginning of the day.
func FridayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnFridayOf(t))
}

// SaturdayOf returns the Saturday of the week passed as time rounded to the
// beginning of the day.
func SaturdayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnSaturdayOf(t))
}

// SundayOf returns the Sunday of the week passed as time rounded to the
// beginning of the day.
func SundayOf(t *time.Time) *time.Time {
	return StartOfDay(SameTimeOnSundayOf(t))
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
	nt := t.Add((time.Duration(time.Sunday-t.Weekday()) * 24 * time.Hour))
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

// SameTimeOnWeekdayOf returns the day of the week passed rounded to the
// beginning of the week day indicated.
func SameTimeOnWeekdayOf(t *time.Time, day string) *time.Time {
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

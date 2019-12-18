package timefmt

import (
	"time"
)

// JanuaryOf returns the beginning of the first day of January of the given
// year.
func JanuaryOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInJanuaryOf(t))
}

// FebruaryOf returns the beginning of the first day of February of the given
// year.
func FebruaryOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInFebruaryOf(t))
}

// MarchOf returns the beginning of the first day of March of the given
// year.
func MarchOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInMarchOf(t))
}

// AprilOf returns the beginning of the first day of April of the given
// year.
func AprilOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInAprilOf(t))
}

// MayOf returns the beginning of the first day of May of the given
// year.
func MayOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInMayOf(t))
}

// JuneOf returns the beginning of the first day of June of the given
// year.
func JuneOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInJuneOf(t))
}

// JulyOf returns the beginning of the first day of July of the given
// year.
func JulyOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInJulyOf(t))
}

// AugustOf returns the beginning of the first day of August of the given
// year.
func AugustOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInAugustOf(t))
}

// SeptemberOf returns the beginning of the first day of September of the given
// year.
func SeptemberOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInSeptemberOf(t))
}

// OctoberOf returns the beginning of the first day of October of the given
// year.
func OctoberOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInOctoberOf(t))
}

// NovemberOf returns the beginning of the first day of November of the given
// year.
func NovemberOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInNovemberOf(t))
}

// DecemberOf returns the beginning of the first day of December of the given
// year.
func DecemberOf(t *time.Time) *time.Time {
	return StartOfMonth(SameTimeInDecemberOf(t))
}

// SameTimeInJanuaryOf returns the exact same month day and time but for the
// month of January instead.
func SameTimeInJanuaryOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInFebruaryOf returns the exact same month day and time but for the
// month of February instead.
func SameTimeInFebruaryOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInMarchOf returns the exact same month day and time but for the
// month of March instead.
func SameTimeInMarchOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInAprilOf returns the exact same month day and time but for the
// month of April instead.
func SameTimeInAprilOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInMayOf returns the exact same month day and time but for the
// month of May instead.
func SameTimeInMayOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInJuneOf returns the exact same month day and time but for the
// month of June instead.
func SameTimeInJuneOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInJulyOf returns the exact same month day and time but for the
// month of July instead.
func SameTimeInJulyOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInAugustOf returns the exact same month day and time but for the
// month of August instead.
func SameTimeInAugustOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInSeptemberOf returns the exact same month day and time but for the
// month of September instead.
func SameTimeInSeptemberOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInOctoberOf returns the exact same month day and time but for the
// month of October instead.
func SameTimeInOctoberOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInNovemberOf returns the exact same month day and time but for the
// month of November instead.
func SameTimeInNovemberOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// SameTimeInDecemberOf returns the exact same month day and time but for the
// month of December instead.
func SameTimeInDecemberOf(t *time.Time) *time.Time {
	// TODO
	return nil
}

// MonthOf returns the beginning of the first day of the specified month of the
// given year.
func MonthOf(t *time.Time, month string) *time.Time {
	switch month {
	case "jan", "january", "Jan", "January":
		return JanuaryOf(t)
		// TODO add more
	}
	return nil
}

// SameTimeInMonthOf returns the exact same month day and time but for the
// specified month instead.
func SameTimeInMonthOf(t *time.Time, month string) *time.Time {
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

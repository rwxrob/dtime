package htime

import (
	"testing"
	"time"
)

func TestMondayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(MondayOf(&then))
}

func TestTuesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(TuesdayOf(&then))
}

func TestWednesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(WednesdayOf(&then))
}

func TestThursdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(ThursdayOf(&then))
}

func TestFridayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(FridayOf(&then))
}

func TestSaturdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SaturdayOf(&then))
}

func TestSundayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SundayOf(&then))
}

func TestSameTimeOnMondayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnMondayOf(&then))
}

func TestSameTimeOnTuesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnTuesdayOf(&then))
}

func TestSameTimeOnWednesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnWednesdayOf(&then))
}

func TestSameTimeOnThursdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnThursdayOf(&then))
}

func TestSameTimeOnFridayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnFridayOf(&then))
}

func TestSameTimeOnSaturdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnSaturdayOf(&then))
}

func TestSameTimeOnSundayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnSundayOf(&then))
}

func TestDayOfWeek(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(DayOfWeek(&then, "tue"))
}

func TestSameTimeOnWeekdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeOnWeekdayOf(&then, "tue"))
}

func TestSameTimeInJanuaryOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeInJanuaryOf(&then))
}

func TestParse(t *testing.T) {
	t.Log(Parse("tue"))
}

func TestParseTime(t *testing.T) {
	t.Log(ParseTime("8p"))
	t.Log(ParseTime("11p"))
	t.Log(ParseTime("16"))
	t.Log(ParseTime("215p"))
	t.Log(ParseTime("2006"))
}

func TestParseYear(t *testing.T) {
	t.Log(ParseYear("1980"))
}

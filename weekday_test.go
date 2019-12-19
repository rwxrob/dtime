package htime_test

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/skilstak/go/htime"
)

func ExampleDayOfWeek() {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	fmt.Println(then)
	fmt.Println(htime.DayOfWeek(&then, "tue"))
	fmt.Println(htime.DayOfWeek(&then, "fri"))
	again, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-11 14:34 -0500")
	fmt.Println(again)
	fmt.Println(htime.DayOfWeek(&again, "tue"))
	fmt.Println(htime.DayOfWeek(&again, "fri"))
	// Output:
	// 2020-05-13 14:34:00 -0500 -0500
	// 2020-05-12 00:00:00 -0500 -0500
	// 2020-05-15 00:00:00 -0500 -0500
	// 2020-05-11 14:34:00 -0500 -0500
	// 2020-05-12 00:00:00 -0500 -0500
	// 2020-05-15 00:00:00 -0500 -0500
}

func TestMondayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.MondayOf(&then))
}

func TestTuesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.TuesdayOf(&then))
}

func TestWednesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.WednesdayOf(&then))
}

func TestThursdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.ThursdayOf(&then))
}

func TestFridayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.FridayOf(&then))
}

func TestSaturdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SaturdayOf(&then))
}

func TestSundayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SundayOf(&then))
}

func TestSameTimeOnMondayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnMondayOf(&then))
}

func TestSameTimeOnTuesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnTuesdayOf(&then))
}

func TestSameTimeOnWednesdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnWednesdayOf(&then))
}

func TestSameTimeOnThursdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnThursdayOf(&then))
}

func TestSameTimeOnFridayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnFridayOf(&then))
}

func TestSameTimeOnSaturdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnSaturdayOf(&then))
}

func TestSameTimeOnSundayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnSundayOf(&then))
}

func TestDayOfWeek(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.DayOfWeek(&then, "tue"))
}

func TestSameTimeOnWeekdayOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(htime.SameTimeOnWeekdayOf(&then, "tue"))
}

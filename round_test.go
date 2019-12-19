package htime_test

import (
	"fmt"
	"time"

	"gitlab.com/skilstak/go/htime"
)

func ExampleMinuteOf() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(htime.MinuteOf(&t))

	// Output:
	// 2020-05-13 14:34:00 -0500 -0500
}

func ExampleHourOf() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(htime.HourOf(&t))

	// Output:
	// 2020-05-13 14:00:00 -0500 -0500
}

func ExampleDayOf() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(htime.DayOf(&t))

	// Output:
	// 2020-05-13 00:00:00 -0500 -0500
}

func ExampleWeekOf() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(htime.WeekOf(&t))

	// Output:
	// 2020-05-11 00:00:00 -0500 -0500
}

func ExampleMonthOf() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(htime.MonthOf(&t))

	// Output:
	// 2020-05-01 00:00:00 -0500 -0500
}

func ExampleYearOf() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(htime.YearOf(&t))

	// Output:
	// 2020-01-01 00:00:00 -0500 -0500
}

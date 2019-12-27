package htime_test

import (
	"fmt"
	"time"

	"gitlab.com/skilstak/go/htime"
)

const format = "2006-01-02 15:04 -0700"

func ExampleSpan() {
	then, _ := time.Parse(format, "2013-05-13 14:34 -0500")
	htime.DefaultTime = &then
	spans := []string{
		"mon",
		"jan",

		"3p",
		"3pm",
		"3Pm",
		"3PM",
		"3pM",

		"304p",
		"304pm",
		"304Pm",
		"304PM",
		"304pM",

		"15",
		"1504",

		"15,mon",
		"1504,mon",
		"304p,mon",
		"304pm,mon",

		"jan2",

		"15,jan2",
		"1504,jan2",
		"304p,jan2",

		",2006",

		"jan,2006",
		"jan2,2006",
		"15,jan2,2006",
		"304p,jan2,2006",

		// ------------- invalid

		"26",   // attempted day, sees as bad hour
		"2006", // attempted year, sees as hour and minute

	}
	for _, s := range spans {
		first, last := htime.Span(s)
		fmt.Printf("span:  %v\nfirst: %v\nlast:  %v\n\n", s, first, last)
	}

	// Output:
	// span:  mon
	// first: 2020-05-11 00:00:00 -0500 -0500
	// last:  <nil>
}

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

package htime_test

import (
	"fmt"
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

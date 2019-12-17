package timefmt_test

import (
	"fmt"
	"time"

	"gitlab.com/skilstak/go/timefmt"
)

func ExampleStartOfMinute() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(timefmt.StartOfMinute(&t))

	// Output:
	// 2020-05-13 14:34:00 -0500 -0500
}

func ExampleStartOfHour() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(timefmt.StartOfHour(&t))

	// Output:
	// 2020-05-13 14:00:00 -0500 -0500
}

func ExampleStartOfDay() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(timefmt.StartOfDay(&t))

	// Output:
	// 2020-05-13 00:00:00 -0500 -0500
}

func ExampleStartOfWeek() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(timefmt.StartOfWeek(&t))

	// Output:
	// 2020-05-11 00:00:00 -0500 -0500
}

func ExampleStartOfMonth() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(timefmt.StartOfMonth(&t))

	// Output:
	// 2020-05-01 00:00:00 -0500 -0500
}

func ExampleStartOfYear() {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2020-05-13 14:34:56 -0500")
	fmt.Println(timefmt.StartOfYear(&t))

	// Output:
	// 2020-01-01 00:00:00 -0500 -0500
}

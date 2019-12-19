package timefmt

import (
	"testing"
	"time"
)

func TestSameTimeInJanuaryOf(t *testing.T) {
	then, _ := time.Parse("2006-01-02 15:04 -0700", "2020-05-13 14:34 -0500")
	t.Log(SameTimeInJanuaryOf(&then))
}

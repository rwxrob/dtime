package timefmt

import "testing"

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
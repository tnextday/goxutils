package xduration

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	testDate := map[string]time.Duration{
		"1h":                 1 * time.Hour,
		"0.5h":               time.Duration(0.5 * float64(time.Hour)),
		"3year2d":            3*365*24*time.Hour + 2*24*time.Hour,
		"4month0.5d3m":       4*30*24*time.Hour + time.Duration(0.5*24.0*float64(time.Hour)) + 3*time.Minute,
		"-1day":              -24 * time.Hour,
		"4 month  0.5d  3 m": 4*30*24*time.Hour + time.Duration(0.5*24.0*float64(time.Hour)) + 3*time.Minute,
		"99year":             99 * 365 * 24 * time.Hour,
	}
	for s, d := range testDate {
		pd, e := Parse(s)
		if e != nil {
			t.Fatal(e)
		}
		t.Log(s, "=", pd)
		if pd != d {
			t.Fatalf("Parse %s = %d, right is %d", s, pd, d)
		}
	}
}

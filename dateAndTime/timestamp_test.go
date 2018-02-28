package dateAndTime

import (
	"math"
	"testing"
	"time"
	"github.com/araddon/dateparse"
)

func assertStrEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatalf("Expected: %s; Got: %s", expected, actual)
	}
}

func assertIntEqual(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Fatalf("Expected: %d; Got: %d", expected, actual)
	}
}

func assertIntArrayEqual(t *testing.T, expected []int, actual []int) {
	for idx, ac := range actual {
		ex := expected[idx]
		if ex != ac {
			t.Fatalf("Expected: %s; Got: %s", expected, actual)
		}
	}
}

func assertAlmostEqual(t *testing.T, expected float64, actual float64) {
	if math.Abs(expected-actual) > 1e-7 {
		t.Fatalf("Expected: %f; Got: %f", expected, actual)
	}
}

func TestConvertStringToTimestamp24HoursFormat(t *testing.T) {
	tm, err := ConvertStringToTime("2018-02-27 20:13:34")
	if err != nil {
		t.Fatal(err)
	}
	assertIntEqual(t, 2018, tm.Year())
	assertIntEqual(t, 2, int(tm.Month()))
	assertIntEqual(t, 27, tm.Day())
	assertIntEqual(t, 20, tm.Hour())
	assertIntEqual(t, 13, tm.Minute())
	assertIntEqual(t, 34, tm.Second())
}

func TestGivenPartialTimeStringExpectDefaultFieldValues(t *testing.T) {
	candiates := map[string][]int{
		"2018-01-01": {2018, 1, 1, 0, 0, 0},
		"2018-02-27 15:58:46": {2018, 2, 27, 15, 58, 46},
		"2018-02-27 03:58:46 PM": {2018, 2, 27, 15, 58, 46},
		"2018-02-27 18:58:46": {2018, 2, 27, 18, 58, 46},
		"2018-02-27T15:58:46": {2018, 2, 27, 15, 58, 46},
		"2018-02-27T18:58:46": {2018, 2, 27, 18, 58, 46},
	}
	for s, expected := range candiates {
		tm, err := ConvertStringToTime(s)
		if err != nil {
			t.Fatal(err)
		}
		actual := []int{tm.Year(), int(tm.Month()), tm.Day(), tm.Hour(), tm.Minute(), tm.Second()}
		assertIntArrayEqual(t, expected, actual)
	}
}

func TestExpectUnixLocalTime(t *testing.T) {
	candiates := map[string]int{
		"2018-02-27 11:54:44": 1519685684,
		"2018-02-20 12:58:23": 1519084703,
		"2018-01-02": 1514804400,
	}
	for s, expected := range candiates {
		tm, err := ConvertStringToTime(s)
		if err != nil {
			t.Fatal(err)
		}
		actual := tm.Local().Unix()
		assertIntEqual(t, expected, int(actual))
	}
}

// implicitly set missing fields to their default values
// implicitly uses 24 hour format
func ConvertStringToTime(s string) (time.Time, error) {
	return dateparse.ParseLocal(s)
}

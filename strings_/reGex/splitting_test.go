package reGex

import (
	"testing"
	"regexp"
)

func assertEqual(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Fatalf("Expected: %s; Got: %s", expected, actual)
	}
	for idx, lhs := range expected {
		rhs := actual[idx]
		if lhs != rhs {
			t.Fatalf("#%d members are different. Expected: %s; Got: %s", idx, lhs, rhs)
		}
	}
}

func assertEmpty(t *testing.T, v []string) {
	if len(v) != 0 {
		t.Fatalf("Not empty! %s", v)
	}
}

func groups(s string, re *regexp.Regexp) []string {
	result := re.FindAllStringSubmatch(s, -1)
	if len(result) == 1 {
		return result[0][1: ]
	}
	return nil
}

func TestSplitDeterministic(t *testing.T) {
	re := regexp.MustCompile(`^(.+):(\d+)$`)
	assertEmpty(t, groups("", re))
	assertEmpty(t, groups("....", re))
	assertEmpty(t, groups("::::", re))
	assertEmpty(t, groups(":0", re))
	assertEqual(t, []string{" ", "0"}, groups(" :0", re))
	assertEqual(t, []string{"asd", "10"}, groups("asd:10", re))
	assertEqual(t, []string{"/doom/e1m1/imp.id1", "123"}, groups("/doom/e1m1/imp.id1:123", re))
}

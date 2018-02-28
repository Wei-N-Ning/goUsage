package dateAndTime

import (
	"testing"
	"time"
)

func assertIntLessThen(t *testing.T, lhs int64, rhs int64) {
	if lhs >= rhs {
		t.Fatalf("lhs (%d) is not less than rhs (%d)", lhs, rhs)
	}
}

func assertTimeEqual(t *testing.T, lhs time.Time, rhs time.Time) {
	if lhs != rhs {
		t.Fatalf("lhs (%s) != rhs (%s)", lhs, rhs)
	}
}

func TestAddDuration(t *testing.T) {
	now := time.Now()
	then := now.Add(-10 * time.Minute)
	assertIntLessThen(t, int64(then.Unix()), int64(now.Unix()))
	assertTimeEqual(t, now, then.Add(10 * time.Minute))
}

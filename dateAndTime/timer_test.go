package dateAndTime

import "testing"
import "time"


func TestElapsedTime(t *testing.T) {
	start := time.Now()
	time.Sleep(1)
	end := time.Now()
	delta := end.Sub(start)
	if (delta < 0) {
		t.Error("not expected")
	}
}
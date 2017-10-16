package containers

import (
	"testing"
)


func TestTwoDimensionalArray(t *testing.T) {
	var twoD [3][3]int
	twoD[0][0] = 111
	if 3 != len(twoD) {
		t.Error(len(twoD))
	}
}


func TestDeclAndInitArray(t *testing.T) {
	nums := [5]int{1, 2, 2}
	if 5 != len(nums) {
		t.Error()
	}
}

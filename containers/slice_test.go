package containers

import (
	"testing"
)


func TestSliceConcept(t *testing.T) {
	nums := []int{}
	if 0 != len(nums) {
		t.Error(nums)
	}
	nums = append(nums, 1, 2)
	if 2 != len(nums) {
		t.Error(nums)
	}
	if 2 != nums[1] {
		t.Error(nums)
	}
}


func TestCopySlice(t *testing.T) {
	nums := []int{1, 2, 3}
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	if 3 != len(numsCopy) {
		t.Error(numsCopy)
	}
}

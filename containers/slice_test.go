package containers

import (
	"testing"
)


func TestSliceConcept(t *testing.T) {
	nums := []int{}
	t.Log(len(nums), nums)
	nums = append(nums, 1, 2)
	t.Log(len(nums), nums)
}


func TestCopySlice(t *testing.T) {
	nums := []int{1, 2, 3}
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	t.Log(numsCopy)
}
package leetcode

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		k int
		result bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 0, false},
		{[]int{1, 2, 1}, 2, true},
		{[]int{1, 2, 3, 1, 1}, 2, true},
	}
	for _, test := range tests {
		result := containsNearbyDuplicate(test.nums, test.k)
		if result != test.result {
			t.Errorf("containsNearbyDuplicate(%v, %v) return error result %v", test.nums, test.k, result)
		}
	}
}

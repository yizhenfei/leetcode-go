package leetcode

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		result bool
	}{
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 2}, false},
		{[]int{1, 2, 2}, true},
		{[]int{2, 1, 2}, true},
	}
	for _, test := range tests {
		result := containsDuplicate(test.nums)
		if result != test.result {
			t.Errorf("containsDuplicate(%v) return error result %v", test.nums, result)
		}
	}
}

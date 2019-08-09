package leetcode

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var tests = []struct {
		nums []int
		sum int
		result int
	}{
		{[]int{}, 1, 0},
		{[]int{2, 1, 1}, 2, 1},
		{[]int{1, 2, 1, 1}, 2, 1},
		{[]int{1, 1, 2}, 2, 1},
		{[]int{1, 1, 1, 2, 1}, 3, 2},
		{[]int{1, 2, 1, 1, 1}, 3, 2},
	}
	for _, test := range tests {
		result := minSubArrayLen(test.sum, test.nums)
		if result != test.result {
			t.Errorf("minSubArrayLen(%v, %v) return error result: %v, expected: %v",
				test.nums, test.sum, result, test.result)
		}
	}
}

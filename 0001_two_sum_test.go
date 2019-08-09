package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums []int
		target int
	}{
		{[]int{0, 1}, 1},
		{[]int{0, 1, 2}, 1},
		{[]int{0, 1, 2}, 3},
		{[]int{0, 1, 2, 4}, 3},
		{[]int{0, 1, 1, 3}, 2},
		{[]int{3, 1, 1, 0}, 2},
		// Additional
		{[]int{3, 2, 4}, 6},
	}
	for _, test := range tests {
		tmp_nums := make([]int, len(test.nums))
		copy(tmp_nums, test.nums)
		result := twoSum(tmp_nums, test.target)
		if test.nums[result[0]] + test.nums[result[1]] != test.target {
			t.Errorf("twoSum(%v, %v) return error result %v", test.nums, test.target, result)
		}
	}
}

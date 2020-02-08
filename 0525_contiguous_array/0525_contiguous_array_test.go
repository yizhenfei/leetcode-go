package leetcode

import "testing"

func TestFindMaxLength(t *testing.T) {
	var tests = []struct {
		nums      []int
		maxLength int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 0}, 2},
		{[]int{1, 1, 1, 0, 0, 1}, 4},
	}
	for _, test := range tests {
		result := findMaxLength(test.nums)
		if test.maxLength != result {
			t.Errorf("findMaxLength(%+v, %v) return error result %v",
				test.nums, test.maxLength, result)
		}
	}
}

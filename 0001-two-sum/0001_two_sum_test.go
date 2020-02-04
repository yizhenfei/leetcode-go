package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
	}{
		{[]int{0, 1}, 1},       // no move
		{[]int{0, 1, 2}, 1},    // move towards left
		{[]int{0, 1, 2}, 3},    // move towards right
		{[]int{0, 1, 2, 4}, 3}, // move towards both directions
		{[]int{0, 1, 1, 3}, 2}, // identity numbers
		{[]int{3, 1, 1, 0}, 2}, // numbers with random order
		// Additional
		{[]int{3, 2, 4}, 6},
	}
	for _, test := range tests {
		input := append(test.nums[:0:0], test.nums...)
		result := twoSum(input, test.target)
		if test.nums[result[0]]+test.nums[result[1]] != test.target {
			t.Errorf("twoSum(%v, %v) return error result %v", test.nums, test.target, result)
		}
	}
}

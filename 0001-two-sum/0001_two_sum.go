package leetcode

import "sort"

type number struct {
	value int
	index int
}

func twoSum(nums []int, target int) []int {
	// Check invalid input
	if len(nums) < 2 {
		panic(false)
	}

	// Convert input into number structs
	numbers := []number{}
	for idx, val := range nums {
		numbers = append(numbers, number{val, idx})
	}

	// Sort nums in ascending order
	sort.Slice(numbers, func(i, j int) bool { return numbers[i].value < numbers[j].value })

	// Scan from both ends to find the match
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left].value + numbers[right].value
		if sum == target {
			return []int{numbers[left].index, numbers[right].index}
		} else if sum < target {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	panic(false)
}

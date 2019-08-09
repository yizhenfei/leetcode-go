package leetcode

import "sort"

type NumWithIdx struct {
	num int
	idx int
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		panic(false)
	}

	// Store index with numbers
	numidxs := []NumWithIdx{}
	for idx, num := range nums {
		numidxs = append(numidxs, NumWithIdx{num, idx})
	}
	// Sort nums in ascending order
	sort.Slice(numidxs, func(i, j int) bool { return numidxs[i].num < numidxs[j].num })
	// Scan from both ends to find the match
	left := 0
	right := len(numidxs) - 1
	for left < right {
		sum := numidxs[left].num + numidxs[right].num
		if sum == target {
			return []int{numidxs[left].idx, numidxs[right].idx}
		} else if sum < target {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	panic(false)
}

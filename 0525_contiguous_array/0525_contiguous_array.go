package leetcode

func findMaxLength(nums []int) int {
	// Special cases
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// Setup basic variables
	maxLength := 0
	count := 0
	occurs := make(map[int]int)
	occurs[0] = -1

	// Iterate over each number
	for idx, num := range nums {
		// Update count
		if num == 0 {
			count--
		} else {
			count++
		}

		// Find in occurence map
		firstIdx, ok := occurs[count]
		if ok {
			// Occured before, then update maxLenth if needed
			if idx-firstIdx > maxLength {
				maxLength = idx - firstIdx
			}
		} else {
			occurs[count] = idx
		}
	}

	// Return result
	return maxLength
}

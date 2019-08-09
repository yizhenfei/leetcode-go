package leetcode

func minSubArrayLen(s int, nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	head := 0
	tail := 0
	sum := nums[0]
	minLen := 0
	
	for ;; {
		// move cursor
		if sum >= s {
			if head < tail {
				sum -= nums[head]
				head = head + 1
			} else {
				return 1
			}
		} else {
			if tail < len(nums)-1 {
				tail = tail + 1
				sum += nums[tail]
			} else {
				return minLen
			}
		}
		// update
		if sum >= s {
			if tail-head+1 < minLen || minLen == 0 {
				minLen = tail-head+1
			}
		}
	}

	panic(false)
}

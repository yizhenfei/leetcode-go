package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	if k < 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= k; j++ {
			if i+j >= len(nums) {
				break
			}
			if nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}

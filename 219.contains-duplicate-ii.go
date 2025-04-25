package leetcode

// @leet start
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)
	n := len(nums)

	for i := 0; i < k; i++ {
		if i >= n {
			return false
		}
		hash[nums[i]]++
		if hash[nums[i]] > 1 {
			return true
		}
	}

	for i := k; i < n; i++ {
		hash[nums[i]]++
		if hash[nums[i]] > 1 {
			return true
		}
		hash[nums[i-k]]--
	}

	return false
}

// @leet end

// Keynold

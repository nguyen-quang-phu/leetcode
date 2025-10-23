package leetcode

import "fmt"

// @leet start
func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	start := 0
	a := start
	b := a + k

	for a + 2*k <=n {
		if nums[a] < nums[a+1] && nums[b] < nums[b+1] {
			a++
		}
	}

	return false
}

// @leet end

// Keynold

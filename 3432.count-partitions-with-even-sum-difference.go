package leetcode

// @leet start
func countPartitions(nums []int) int {
	n := len(nums)

	sum := 0

	for i := 1; i <= n; i++ {
		sum += nums[i-1]
	}

	if sum%2 != 0 {
		return 0
	}

	return n - 1
}
// @leet end

// Keynold

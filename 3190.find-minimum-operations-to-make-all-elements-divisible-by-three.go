package leetcode

// @leet start
func minimumOperations(nums []int) int {
	n := len(nums)
	result := 0

	for i := 0; i < n; i++ {
		if nums[i]%3 != 0 {
			result++
		}
	}
	return result
}
// @leet end

// Keynold

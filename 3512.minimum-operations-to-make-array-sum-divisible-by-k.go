package leetcode

// @leet start
func minOperations(nums []int, k int) int {
	result := 0
	sum := 0

	for _, num := range nums {
		sum += num
	}

	for sum%k != 0 {
		sum -= 1
		result++
	}

	return result
}
// @leet end

// Keynold

package leetcode

import "fmt"

// @leet start
func pivotIndex(nums []int) int {
	sum := 0
	n:= len(nums)

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	leftSum := 0
	rightSum := sum - nums[0]

	for i := 0; i < n; i++ {
		if i == 0 {
			leftSum = 0
			rightSum = sum - nums[0]
		} else if i == n-1 {
			rightSum = 0
			leftSum = sum - nums[n-1]
		} else {
			leftSum += nums[i-1]
			rightSum -= nums[i]
		}
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}
// @leet end

// Keynold

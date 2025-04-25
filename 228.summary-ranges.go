package leetcode

import "fmt"

// @leet start
func stringFormat(nums []int, left int, right int) string {
	max := right
	n := len(nums)
	if right == n {
		max = n
	}
	if right-left == 1 {
		return fmt.Sprintf("%d", nums[left])
	}
	return fmt.Sprintf("%d->%d", nums[left], nums[max-1])
}

func summaryRanges(nums []int) []string {
	result := []string{}
	left, right, n := 0, 0, len(nums)

	for i := range nums {
		right++
		if right >= n || nums[right]-nums[i] != 1 {
			item := stringFormat(nums, left, right)
			result = append(result, item)
			left = i + 1
			right = left
		}
	}

	return result
}

// @leet end

// Keynold

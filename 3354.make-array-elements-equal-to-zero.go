package leetcode

import "fmt"

// @leet start
func countValidSelections(nums []int) (count int) {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	left := 0
	right:= sum

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if left-right >=0 && left-right <= 1 {
				count++
			}

			if right-left >=0 && right-left <= 1 {
        count++
      }

		} else {
			left += nums[i]
			right -= nums[i]
		}
	}

	return count
}
// @leet end

// Keynold

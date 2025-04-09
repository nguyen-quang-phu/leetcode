package leetcode

// @leet start

func minOperations(nums []int, k int) int {
	hash := make(map[int]bool)
	min := nums[0]
	for _, num := range nums {
		hash[num] = true
		if min > num {
			min = num
		}
	}

	if min < k {
		return -1
	}

	result := len(hash)

	if min == k {
		return result - 1
	}

	return result
}

// @leet end

// Keynold

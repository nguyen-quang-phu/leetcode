package leetcode

// @leet start
func minimumOperations(nums []int) int {
	hash := make(map[int]bool)
	n := len(nums)
	index := 0

	for i := n - 1; i >= 0; i-- {
		if _, ok := hash[nums[i]]; ok {
			index = i + 1
			break
		} else {
			hash[nums[i]] = true
		}
	}

	return (index + 2) / 3
}

// @leet end

// Keynold

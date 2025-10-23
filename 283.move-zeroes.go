package leetcode

// @leet start
func moveZeroes(nums []int)  {
	n := len(nums)
	sentinel := 0

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[sentinel] = nums[i]
			sentinel++
		} else {
			continue
		}
	}

	for i := sentinel; i < n; i++ {
		nums[i] = 0
	}
}
// @leet end

// Keynold

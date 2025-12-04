package leetcode

// @leet start
func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	curr := 0
	res := make([]bool, n)

	for i := 0; i < n; i++ {
		curr = ((curr << 1) + nums[i] ) % 5
		res[i] = curr%5 == 0
	}
	return res
}
// @leet end

// Keynold

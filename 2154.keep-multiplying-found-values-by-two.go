package leetcode

// @leet start
func findFinalValue(nums []int, original int) int {
	n:= len(nums)
	curr := original

	freq := make(map[int]bool, n)

	for _, v := range nums {
		freq[v] = true
	}

	for {
		if _, ok := freq[curr]; !ok {
			return curr
		}

		curr *= 2
	}
}
// @leet end

// Keynold

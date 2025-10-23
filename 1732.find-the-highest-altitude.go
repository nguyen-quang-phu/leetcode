package leetcode

// @leet start
func largestAltitude(gain []int) int {
	result := 0
	sum := 0
	n := len(gain)

	for i := 0; i < n; i++ {
		sum += gain[i]
		result = max(result, sum)
	}

	return result
}
// @leet end

// Keynold

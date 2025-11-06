package leetcode

// @leet start
func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i & (i-1) == 0 {
			result[i] = 1
		} else {
			if i%2==0 {
				result[i] = result[i/2]
			} else {
				result[i] = result[i/2] + 1
			}
		}
	}
	return result
}
// @leet end

// Keynold

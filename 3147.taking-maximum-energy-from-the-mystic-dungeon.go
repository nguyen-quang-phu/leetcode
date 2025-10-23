package leetcode

// @leet start
func maximumEnergy(energy []int, k int) int {

	n := len(energy)
	result := energy[n-1]
	dp := make([]int, n)

	for i := 0; i < k; i++ {
		dp[n-1-i] = energy[n-1-i]
		if dp[n-1-i] > result {
			result = dp[n-1-i]
		}
	}

	for i := n - k - 1; i >= 0; i-- {
		dp[i] = energy[i] + dp[i+k]
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

// @leet end
// Keynold

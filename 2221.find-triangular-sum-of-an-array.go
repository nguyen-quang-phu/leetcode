package leetcode

// @leet start
func triangularSum(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
    dp[i] = make([]int, i+1)
  }
  dp[0][0] = nums[0]
  for i := 1; i < len(nums); i++ {
    for j := 0; j < i+1; j++ {
      dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]) % 10
    }
  }
  return dp[len(nums)-1][0]

}
// @leet end

// Keynold

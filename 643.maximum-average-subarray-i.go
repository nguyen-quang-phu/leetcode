package leetcode
// @leet start
func findMaxAverage(nums []int, k int) float64 {
	left, right := 0,k
	n := len(nums)
	sum := 0

	for i:= 0; i < k; i++ {
		sum += nums[i]
	}

	result := sum
	for right < n {
		sum = sum - nums[left] + nums[right]
		result = max(result, sum)
		left++
		right++
	}


	return float64(result) / float64(k)
}
// @leet end

// Keynold

package leetcode

// @leet start
func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	minPrefSum := make([]int64, k)

	var prefSum int64 = 0
	for i := 0; i < k-1; i++ {
		prefSum += int64(nums[i])
		minPrefSum[i+1] = prefSum
	}

	res := prefSum + int64(nums[k-1])
	for i := k-1; i < n; i++ {
		prefSum += int64(nums[i])
		cur := prefSum - minPrefSum[(i+1)%k]
		if cur > res {
			res = cur
		}
		if prefSum < minPrefSum[(i+1)%k] {
			minPrefSum[(i+1)%k] = prefSum
		}
	}

	return res
}

// @leet end

// Keynold

package leetcode

// @leet start
func repeatedNTimes(nums []int) int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++

		if freq[num] > 1 {
			return num
		}
	}

	return -1
}
// @leet end

// Keynold

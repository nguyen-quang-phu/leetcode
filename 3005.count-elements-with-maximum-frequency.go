package leetcode

// @leet start
func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	maxFreq := 0
	hashFreq := make(map[int]int)
	for _, count := range freq {
		hashFreq[count]++
    if count > maxFreq {
      maxFreq = count
    }
	}

	return hashFreq[maxFreq] * maxFreq
}
// @leet end

// Keynold

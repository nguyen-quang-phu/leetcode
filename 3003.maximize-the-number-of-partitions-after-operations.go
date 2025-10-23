package leetcode

// @leet start
func maxPartitionsAfterOperations(s string, k int) int {
	freq := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}


}
// @leet end

// Keynold

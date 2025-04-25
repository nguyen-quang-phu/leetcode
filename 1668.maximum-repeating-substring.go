package leetcode

// @leet start
func maxRepeating(sequence string, word string) int {
	n := len(sequence)
	m := len(word)
	count := 0
	result := 0

	for i := 0; i <= n-m; i++ {
		if sequence[i:i+m] == word {
			count++
			i += m - 1
		} else {
			i -= count * m
			count = 0
		}

		if result < count {
			result = count
		}
	}

	return result
}

// @leet end

// Keynold

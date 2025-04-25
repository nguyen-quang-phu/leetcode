package leetcode

// @leet start
func numRabbits(answers []int) int {
	hash := make(map[int]int)
	result := 0

	for _, answer := range answers {
		hash[answer]++
	}

	for answer, count := range hash {
		result += (count + answer) / (answer + 1) * (answer + 1)
	}
	return result
}

// @leet end

// Keynold

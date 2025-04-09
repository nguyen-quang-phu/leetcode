package leetcode

// @leet start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	result := n
	for {
		if result%2 == 0 {
			result /= 2
		} else if result%3 == 0 {
			result /= 3
		} else if result%5 == 0 {
			result /= 5
		} else {
			break
		}
	}

	return result == 1
}

// @leet end

// Keynold

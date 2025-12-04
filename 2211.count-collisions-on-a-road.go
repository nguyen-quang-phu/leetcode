package leetcode

// @leet start
func countCollisions(directions string) int {
	n := len(directions)
	left, right := 0, n-1

	// Skip all 'L' from the left
	for left < n && directions[left] == 'L' {
		left++
	}

	// Skip all 'R' from the right
	for right >= 0 && directions[right] == 'R' {
		right--
	}

	count := 0
	for i := left; i <= right; i++ {
		if directions[i] != 'S' {
			count++
		}
	}
	return count
}
// @leet end

// Keynold

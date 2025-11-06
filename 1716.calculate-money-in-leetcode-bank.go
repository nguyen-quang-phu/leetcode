package leetcode

// @leet start
func totalMoney(n int) int {
	start := -1
	resutl := 0

	for i := 0; i < n; i++ {
		if i%7 == 0 {
			start++
		}

		resutl += start + (i % 7) + 1
  }

	return resutl
}

// @leet end

// Keynold

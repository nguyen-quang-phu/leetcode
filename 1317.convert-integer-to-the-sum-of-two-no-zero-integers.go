package leetcode

// @leet start
func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n/2; i++ {
		if isNoZero(i) && isNoZero(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{}
}

func isNoZero(i int) bool {
	for i > 0 {
		if i%10 == 0 {
			return false
		}
		i /= 10
	}

	return true
}

// @leet end

// Keynold

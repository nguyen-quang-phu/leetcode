package leetcode

// @leet start
func smallestNumber(n int) int {
	res := 1
	for {
		if res-1 >= n {
			return res - 1
		}

		res = res << 1
	}
}

// @leet end

// Keynold

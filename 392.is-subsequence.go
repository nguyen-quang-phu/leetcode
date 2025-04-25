package leetcode

// @leet start
func isSubsequence(s string, t string) bool {
	index := 0
	n := len(t)
	m := len(s)
	if m == 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if index < m && s[index] == t[i] {
			index++
		}

		if index == m {
			return true
		}
	}

	return false
}

// @leet end

// Keynold

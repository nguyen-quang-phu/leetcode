package leetcode

// @leet start
func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	s_index := 0

	for i := 0; i < tLen; i++ {
		if s_index == sLen {
			break
		}
		if t[i] == s[s_index] {
			s_index++
		}
	}

	return s_index == sLen
}
// @leet end

// Keynold

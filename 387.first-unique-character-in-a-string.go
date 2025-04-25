package leetcode

// @leet start
func firstUniqChar(s string) int {
	hash := make([]int, 26)
	n := len(s)

	for i := 0; i < n; i++ {
		hash[s[i]-'a']++
	}

	for i := 0; i < n; i++ {
		if hash[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// @leet end

// Keynold

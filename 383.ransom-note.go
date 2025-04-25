package leetcode

// @leet start
func canConstruct(ransomNote string, magazine string) bool {
	n := len(magazine)
	m := len(ransomNote)
	if m > n {
		return false
	}
	hash := make([]int, 26)

	for i := 0; i < n; i++ {
		hash[magazine[i]-'a']++
	}

	for i := 0; i < m; i++ {
		hash[ransomNote[i]-'a']--

		if hash[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

// @leet end

// Keynold

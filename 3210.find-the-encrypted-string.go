package leetcode

// @leet start
func getEncryptedString(s string, k int) string {
	n := len(s)
	encrypted := []byte(s)
	hash := s + s

	for index := range s {
		encrypted[index] = hash[index+k%n]
	}

	return string(encrypted)
}

// @leet end

// Keynold

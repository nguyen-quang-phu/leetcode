package leetcode

import (
	"fmt"
	"strings"
)

// @leet start
func canBeTypedWords(text string, brokenLetters string) int {
	hash := make(map[byte]int)

	for i := 0; i < len(brokenLetters); i++ {
		hash[brokenLetters[i]]++
	}

	words := strings.Split(text, " ")
	res := 0

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if _, ok := hash[words[i][j]]; ok {
				res++
				break
			}
		}
	}

	return len(words) - res
}
// @leet end

// Keynold

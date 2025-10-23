package leetcode

// @leet start
func reverseWords(s string) string {
	charArray := []rune(s)
	n := len(charArray)
	left, right := 0, 0
	i:=0

	for i := 0; i < n/2; i++ {
		charArray[i], charArray[n-i-1] = charArray[n-i-1], charArray[i]
	}


}
// @leet end

// Keynold

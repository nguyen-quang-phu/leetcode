package leetcode


// @leet start

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	res := gcd(len(str1), len(str2))

	return str1[:res]
}

// @leet end

// Keynold

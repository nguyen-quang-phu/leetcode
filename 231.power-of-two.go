package leetcode

// @leet start
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
// @leet end

// Keynold

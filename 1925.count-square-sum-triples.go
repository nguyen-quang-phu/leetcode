package leetcode

// @leet start
func countTriples(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			k := i*i + j*j
			if k > n*n {
				break
			}
			if isPerfectSquare(k) {
				count++
			}
		}
	}
	return count * 2
}

func isPerfectSquare(x int) bool {
	if x < 0 {
		return false
	}
	s := int(math.Sqrt(float64(x)))
	return s*s == x
}
// @leet end

// Keynold

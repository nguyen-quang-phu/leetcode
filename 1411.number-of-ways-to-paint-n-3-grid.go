package leetcode

// @leet start
func numOfWays(n int) int {
	const MOD int64 = 1000000007
	var A, B int64 = 6, 6

	for i := 2; i <= n; i++ {
		newA := (2*A + 2*B) % MOD
		newB := (2*A + 3*B) % MOD
		A, B = newA, newB
	}

	return int((A + B) % MOD)
}
// @leet end

// Keynold

package leetcode

// @leet start
func numSub(s string) int {
	n := len(s)

	res := 0

	count := 0
	mod :=  int(1e9+7)

	for i := 0; i <= n; i++ {
		if i == n {
			res += count * (count + 1) / 2
			res = res % mod
			return res
		}

		if s[i] == '1' {
			count++
		} else {

			res += count * (count + 1) / 2
			res = res % mod
			count = 0
		}
	}

	return res
}
// @leet end

// Keynold

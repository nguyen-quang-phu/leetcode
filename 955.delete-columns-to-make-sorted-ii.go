package leetcode

// @leet start

func minDeletionSize(strs []string) int {
	l := len(strs)
	n := len(strs[0])
	res := 0
	flags := make([]bool, n)
LOOP:
	for i := 1; i < l; i++ {
		for j := 0; j < n; j++ {
			if flags[j] || strs[i][j] == strs[i-1][j] {
				continue
			}
			if strs[i][j] < strs[i-1][j] {
				flags[j] = true
				res++
				i = 0
			}
			continue LOOP
		}
	}
	return res
}
// @leet end

// Keynold

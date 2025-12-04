package leetcode

// @leet start
func numberOfPaths(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	d := make([][]int, m)
	prev := make([][]int, m)
	mod := 1e9 + 7
	for i := 0; i < m; i += 1 {
		d[i] = make([]int, k)
		prev[i] = make([]int, k)
	}
	prev[0][0] = 1
	for i := 0; i < n; i += 1 {
		for j := 0; j < m; j += 1 {
			for l := 0; l < k; l += 1 {
				l2 := (grid[i][j] + l) % k
				d[j][l2] = prev[j][l]
				if j > 0 {
					d[j][l2] = (d[j][l2] + d[j-1][l]) % mod
				}
			}
		}
		prev, d = d, prev
	}
	return prev[m-1][0]
}
// @leet end

// Keynold

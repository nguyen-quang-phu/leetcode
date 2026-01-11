package leetcode

// @leet start
func countNegatives(grid [][]int) int {
	count := 0
	m := len(grid)
	n := len(grid[0])

	for i,j:= m-1,0; i >= 0 && j < n; {
		if grid[i][j] < 0 {
			count += n - j
			i--
		} else {
			j++
		}
	}
	return count
}
// @leet end

// Keynold


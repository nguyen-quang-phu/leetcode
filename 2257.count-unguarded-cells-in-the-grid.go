package leetcode

// @leet start


func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	const (
		UNGUARDED = 0
		GUARDED   = 1
		GUARD     = 2
		WALL      = 3
	)

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	var dfs func(int, int, byte)

	dfs = func(row, col int, direction byte) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
			return
		}
		if grid[row][col] == GUARD || grid[row][col] == WALL {
			return
		}
		if grid[row][col] == GUARDED {
			// Đã được canh rồi thì không cần set lại
		}
		grid[row][col] = GUARDED
		switch direction {
		case 'U':
			dfs(row-1, col, 'U')
		case 'D':
			dfs(row+1, col, 'D')
		case 'L':
			dfs(row, col-1, 'L')
		case 'R':
			dfs(row, col+1, 'R')
		}
	}

	// Mark guards
	for _, g := range guards {
		grid[g[0]][g[1]] = GUARD
	}

	// Mark walls
	for _, w := range walls {
		grid[w[0]][w[1]] = WALL
	}

	// Spread guarded areas
	for _, g := range guards {
		dfs(g[0]-1, g[1], 'U')
		dfs(g[0]+1, g[1], 'D')
		dfs(g[0], g[1]-1, 'L')
		dfs(g[0], g[1]+1, 'R')
	}

	// Count unguarded cells
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == UNGUARDED {
				count++
			}
		}
	}

	return count
}
// @leet end

// Keynold

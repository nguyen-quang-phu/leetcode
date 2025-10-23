package leetcode

// @leet start
func dfs(grid [][]int, visited [][]bool, x, y, limit int) bool {
	n := len(grid)
	if x==n-1 && y==n-1 {
    return true
  }
  visited[x][y] = true
  for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
    newX, newY := x+dir[0], y+dir[1]
    if newX < 0 || newX >= n || newY < 0 || newY >= n || visited[newX][newY] || grid[newX][newY] > limit {
      continue
    }
    if dfs(grid, visited, newX, newY, limit) {
      return true
    }
  }
  return false

}
func swimInWater(grid [][]int) int {
	n := len(grid)

	visited := make([][]bool, n)

	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	time := 0
  for time < grid[0][0] {
    if dfs(grid, visited, 0, 0, time) {
      return time
    }
    time++
  }
  return time
}
// @leet end

// Keynold

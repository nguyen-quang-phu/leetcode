package leetcode

// @leet start
func flowerGame(n int, m int) int64 {
	odd_left := (n + 1)/2
	even_left := n - odd_left
	odd_right := (m + 1)/2
	even_right := m - odd_right

	return int64(odd_left * even_right + even_left * odd_right)

}
// @leet end

// Keynold

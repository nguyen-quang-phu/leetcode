package leetcode

// @leet start
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
			if !isValidUnit(board[i]) {
					return false
			}
			col := make([]byte, 9)
			for j := 0; j < 9; j++ {
					col[j] = board[j][i]
			}
			if !isValidUnit(col) {
					return false
			}
			box := make([]byte, 9)
			for j := 0; j < 3; j++ {
					for k := 0; k < 3; k++ {
							box[j*3+k] = board[(i/3)*3+j][(i%3)*3+k]
					}
			}
			if !isValidUnit(box) {
					return false
			}
	}
	return true

}

func isValidUnit(unit []byte) bool {
		seen := make(map[byte]bool)
		for _, v := range unit {
				if v != '.' {
						if seen[v] {
								return false
						}
						seen[v] = true
				}
		}
		return true
}
// @leet end

// Keynold

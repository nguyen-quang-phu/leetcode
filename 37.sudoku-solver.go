package leetcode

// @leet start
func solveSudoku(board [][]byte) {
	rows, cols, boxes := [9]int{}, [9]int{}, [9]int{}
	empties := make([][2]int, 0)

	place := func(r, c, num int) {
		m := 1 << num
		rows[r] |= m
		cols[c] |= m
		boxes[(r/3)*3+c/3] |= m
	}
	remove := func(r, c, num int) {
		m := ^(1 << num)
		rows[r] &= m
		cols[c] &= m
		boxes[(r/3)*3+c/3] &= m
	}
	countBits := func(n int) int {
		c := 0
		for n > 0 {
			n &= n - 1
			c++
		}
		return c
	}
	bitPos := func(mask int) int {
		p := 0
		for (1 << p) != mask {
			p++
		}
		return p
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				empties = append(empties, [2]int{i, j})
			} else {
				place(i, j, int(board[i][j]-'1'))
			}
		}
	}

	var backtrack func() bool
	backtrack = func() bool {
		if len(empties) == 0 {
			return true
		}
		minOpt, idx, mask := 10, -1, 0
		for k, cell := range empties {
			r, c := cell[0], cell[1]
			b := (r/3)*3 + c/3
			used := rows[r] | cols[c] | boxes[b]
			opt := 9 - countBits(used)
			if opt < minOpt {
				minOpt, idx, mask = opt, k, (^used)&0x1FF
				if opt == 1 {
					break
				}
			}
		}
		r, c := empties[idx][0], empties[idx][1]
		empties = append(empties[:idx], empties[idx+1:]...)
		for mask > 0 {
			pick := mask & -mask
			num := bitPos(pick)
			place(r, c, num)
			board[r][c] = byte(num + '1')
			if backtrack() {
				return true
			}
			remove(r, c, num)
			board[r][c] = '.'
			mask -= pick
		}
		empties = append(empties[:idx], append([][2]int{{r, c}}, empties[idx:]...)...)
		return false
	}
	backtrack()
}

// @leet end

// Keynold

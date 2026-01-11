package leetcode

// @leet start
func maxMatrixSum(matrix [][]int) int64 {
	var total int64 = 0
	var minAbs int64 = -1
	odd := false

	for _, row := range matrix {
		for _, val := range row {
			absVal := val
			if val < 0 {
				absVal = -val
				odd = !odd
			}
			total += int64(absVal)

			if minAbs == -1 || int64(absVal) < minAbs {
				minAbs = int64(absVal)
			}
		}
	}

	if odd {
		total -= 2 * minAbs
	}
	return total
}
// @leet end

// Keynold

package leetcode

// @leet start
func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiaSq := 0
	maxArea := 0
	for _, dim := range dimensions {
		l := dim[0]
		w := dim[1]
		diaSq := l*l + w*w
		area := l * w
		if diaSq > maxDiaSq {
			maxDiaSq = diaSq
			maxArea = area
		} else if diaSq == maxDiaSq {
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
// @leet end

// Keynold

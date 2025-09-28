package leetcode

import "math"

// @leet start
func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			for k := j + 1; k < len(points); k++ {
				x3, y3 := points[k][0], points[k][1]
				area := 0.5 * math.Abs(float64(x1*y2+x2*y3+x3*y1-y1*x2-y2*x3-y3*x1))
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
// @leet end

// Keynold

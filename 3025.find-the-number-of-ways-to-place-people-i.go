package leetcode

import (
	"fmt"
	"sort"
)

// @leet start
func numberOfPairs(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] > points[j][1] // y descending
        }
        return points[i][0] < points[j][0]     // x ascending
    })

		fmt.Println("points", points)
    n := len(points)
    result := 0

    for i := 0; i < n; i++ {
        top := points[i][1]
        bot := -1 << 31 // INT_MIN
        for j := i + 1; j < n; j++ {
            y := points[j][1]
						fmt.Println("i, j, points[i], points[j]", i, j, points[i], points[j])
            if bot < y && y <= top {
								fmt.Println(i, j, points[i], points[j])
                result++
                bot = y
                if bot == top {
                    break
                }
            }
        }
    }
    return result
}
// @leet end

// Keynold

package leetcode

// @leet start
import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	n := len(intervals)
	todo := make([]int, n)
	for i := range todo {
		todo[i] = 2
	}

	ans := 0

	for t := n - 1; t >= 0; t-- {
		start := intervals[t][0]
		// e := intervals[t][1]
		m := todo[t] // số điểm còn thiếu của interval t

		for p := start; p < start+m; p++ {
			for i := 0; i <= t; i++ {
				if todo[i] > 0 && p <= intervals[i][1] {
					todo[i]--
				}
			}
			ans++
		}
	}

	return ans
}
// @leet end

// Keynold

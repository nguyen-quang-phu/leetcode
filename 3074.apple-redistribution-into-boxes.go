package leetcode

import "sort"

// @leet start
func minimumBoxes(apple []int, capacity []int) int {
	sum := 0

	for i := 0; i < len(apple); i++ {
		sum += apple[i]
	}

	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	count := 0

	for _, cap := range capacity {
		sum -= cap

		if sum <= 0 {
			return count + 1
		} else {
			count++
		}
	}

	return count
}
// @leet end

// Keynold

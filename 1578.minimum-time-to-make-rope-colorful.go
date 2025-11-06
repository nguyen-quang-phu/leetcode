package leetcode

// @leet start
func minCost(colors string, neededTime []int) int {
	result := 0
	n := len(colors)

	curr := 0
	localMax := neededTime[curr]
	localSum := neededTime[curr]

	for i := 1; i <= n; i++ {
		if i == n {
			result += localSum - localMax
		} else if colors[i] == colors[i-1] {
			localMax = max(localMax, neededTime[i])
			localSum += neededTime[i]
		} else {
			curr = i
			result += localSum - localMax
			localMax = neededTime[curr]
			localSum = neededTime[curr]
		}
	}

	return result
}

// @leet end

// Keynold

package leetcode

// @leet start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0

	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	threshold := max - extraCandies
	res := make([]bool, len(candies))

	for i, c := range candies {
		if c >= threshold {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res;
}
// @leet end

// Keynold

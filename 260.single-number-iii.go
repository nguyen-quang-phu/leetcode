package leetcode

// @leet start
func singleNumber(nums []int) []int {
	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}

	result := []int{}

	for num, count := range hash {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

// @leet end

// Keynold

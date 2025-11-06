package leetcode

// @leet start
func getSneakyNumbers(nums []int) []int {
	freq := make(map[int]int)
	res := []int{}

	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			res = append(res, num)
		}
	}

	return res
}
// @leet end

// Keynold

package leetcode

// @leet start
func isValidTriangleNumber(nums []int) bool {
	if len(nums) != 3 {
		return false
	}

	if nums[0]+nums[1] > nums[2] && nums[1]+nums[2] > nums[0] && nums[0]+nums[2] > nums[1] {
		return true
	}

	return false
}
func triangleNumber(nums []int) int {
	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}


}
// @leet end

// Keynold

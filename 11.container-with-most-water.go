package leetcode

// @leet start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		result = max(result, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
// @leet end

// Keynold

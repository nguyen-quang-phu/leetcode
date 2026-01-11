package leetcode

// @leet start
func plusOne(digits []int) []int {
	n := len(digits)
	carry := true

	for i := n - 1; i >= 0; i-- {
		if carry {
			digits[i]++
			if digits[i] == 10 {
				digits[i] = 0
				carry = true
			} else {
				carry = false
			}
		} else {
			break
		}
	}

	if carry {
		digits = append([]int{1}, digits...)
	}

	return digits
}
// @leet end

// Keynold

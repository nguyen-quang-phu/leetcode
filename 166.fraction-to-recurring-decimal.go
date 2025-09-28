package leetcode

import "strconv"

// @leet start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator % denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	return strconv.Itoa(numerator/denominator) + "." + strconv.Itoa(numerator*10/denominator)
}
// @leet end

// Keynold

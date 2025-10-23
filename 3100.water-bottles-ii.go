package leetcode

// @leet start
func maxBottlesDrunk(numBottles int, numExchange int) int {
	res := numBottles

	for numBottles >= numExchange {
		res++
		numBottles -= numExchange - 1
		numExchange++
	}
	return res
}
// @leet end

// Keynold

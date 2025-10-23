package leetcode

import "runtime/debug"

// @leet start
var FREQ [100000]int

func init() {
	debug.SetMemoryLimit(9_000_000)
}

func findSmallestInteger(nums []int, value int) int {
	n := len(nums)
	if value == 1 {
		return n
	}
	for i := range value {
		FREQ[i] = 0
	}

	for _, num := range nums {
		num %= value
		if num < 0 {
			num += value
		}
		FREQ[num]++
	}

	mex, minfreq := 0, n
	for i := 0; i < value && minfreq != 0; i++ {
		if minfreq > FREQ[i] {
			minfreq = FREQ[i]
			mex = i
		}
	}
	return mex + value*minfreq
}

// @leet end

// Keynold

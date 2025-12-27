package leetcode

import "sort"

// @leet start
func maximumHappinessSum(happiness []int, k int) int64 {
    n, count := len(happiness), 1
    sort.Ints(happiness)
    res := happiness[n - 1]
    for i := n - 2; i >= n - k; i-- {
        if happiness[i] - count > 0 {
            res += happiness[i] - count
        }
        count++
    }
    return int64(res)
}

// @leet end

// Keynold

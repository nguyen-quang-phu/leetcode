package leetcode

// @leet start
func findXSum(nums []int, k int, x int) []int {
    freq := make(map[int]int)
    for i := range k - 1 {
        freq[nums[i]]++
    }

    var ans []int
    for i := k - 1 ; i < len(nums) ; i++ {
        freq[nums[i]]++
        ans = append(ans, find(freq, x))
        freq[nums[i - k + 1]]--
    }

    return ans
}

func find(freq map[int]int, x int) int {
    var keys []int
    for k := range freq {
        keys = append(keys, k)
    }

    sort.Slice(keys, func(i, j int) bool {
        if freq[keys[i]] == freq[keys[j]] {
            return keys[i] > keys[j]
        }

        return freq[keys[i]] > freq[keys[j]]
    })

    ans := 0
    for i := range min(x, len(keys)) {
        ans += (keys[i] * freq[keys[i]])
    }

    return ans
}
// @leet end

// Keynold

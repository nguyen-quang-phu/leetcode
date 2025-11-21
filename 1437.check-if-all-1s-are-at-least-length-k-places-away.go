package leetcode

// @leet start
func kLengthApart(nums []int, k int) bool {
	n := len(nums)
	count := 0
	firstOneFound := -1
	lastOneFound := -1


	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			firstOneFound = i
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if nums[i] == 1 {
			lastOneFound = i
			break
		}
	}

	if firstOneFound == -1 || firstOneFound == lastOneFound {
		return true
	}

	for i := firstOneFound+1; i < lastOneFound; i++ {
		if nums[i] == 0 {
			count++
		} else {
			if count < k {
				return false
			}
			count = 0
		}
	}

	if count < k {
		return false
	}

	return true
}
// @leet end

// Keynold

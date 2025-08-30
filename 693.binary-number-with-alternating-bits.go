package leetcode

// @leet start
func hasAlternatingBits(n int) bool {
	m:= n
	if m % 2 == 0 {
		m = m >> 1
	}

	flag := 1
	for i:=0; i< 16; i++ {
		if m ^ flag == 0 {
			return true
		}
		flag = flag << 2 + 1
	}
	return false
}
// @leet end

// Keynold

package leetcode

// @leet start
func numberOfWays(corridor string) int {
	mod := int(1e9 + 7)
	seats := make([]int, 0, len(corridor))
	for i, ch := range corridor {
		if ch == 'S' {
			seats = append(seats, i)
		}
	}
	n := len(seats)
	if n <= 1 || n%2 != 0 {
		return 0
	}
	ways := 1
	for i := 1; i < n-1; i += 2 {
		ways = (ways * (seats[i+1] - seats[i])) % mod
	}
	return ways
}
// @leet end

// Keynold

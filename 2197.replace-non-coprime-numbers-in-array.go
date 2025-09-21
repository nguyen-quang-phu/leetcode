package leetcode

// @leet start
func replaceNonCoprimes(nums []int) []int {
	ans := []int{}
	for _, num := range nums {
		for len(ans) > 0 {
			last := ans[len(ans)-1]
			g := gcd(last, num)
			if g > 1 {
				num = last / g * num
				ans = ans[:len(ans)-1]
			} else {
				break
			}
		}
		ans = append(ans, num)
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
// @leet end

// Keynold

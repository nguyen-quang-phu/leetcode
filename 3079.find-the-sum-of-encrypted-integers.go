package leetcode

// @leet start
func sumOfEncryptedInt(nums []int) int {
	res:= 0
	for _, num := range nums {
		res += encrypt(num)
	}
	return res

}

func encrypt(num int) int {
	if num < 10 {
		return num
	}

	n := 1
	max:=0

	for num > 0 {
		temp:= num % 10
		if temp > max {
			max = temp
		}
		num /= 10
		n++
	}

	res:=0

	for n > 1 {
		res = res * 10 + max
		n--
	}


	return res
}
// @leet end

// Keynold

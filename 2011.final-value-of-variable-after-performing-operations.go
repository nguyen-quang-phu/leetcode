package leetcode

// @leet start
func finalValueAfterOperations(operations []string) (x int) {
	for i := 0; i < len(operations); i++ {
    if operations[i][1] == '+' {
      x++
    } else {
      x--
    }
  }

	return x
}
// @leet end

// Keynold

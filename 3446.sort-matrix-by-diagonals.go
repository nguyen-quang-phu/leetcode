package leetcode

// @leet start
func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	m := 2*n - 1
	store := make([][]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			store[i-j] = append(store[i-j], grid[i][j])
		}
	}
	for k := range store {
		arr := store[k]
		quickSort(arr, 0, len(arr)-1)
		store[k] = arr
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = store[i-j][0]
			store[i-j] = store[i-j][1:]
		}
	}
	return grid
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	pivot := arr[(left+right)>>1]
	for l <= r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if left < r {
		quickSort(arr, left, r)
	}
	if l < right {
		quickSort(arr, l, right)
	}
}
// @leet end

// Keynold

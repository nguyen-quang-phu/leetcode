package leetcode

// @leet start
func canPlaceFlowers(flowerbed []int, n int) bool {
    len := len(flowerbed)
    for i := 0; i < len; i++ {
        left := i == 0 || flowerbed[i-1] == 0
        right := i == len-1 || flowerbed[i+1] == 0

        if left && right && flowerbed[i] == 0 {
            flowerbed[i] = 1
            n--
        }
    }
    return n <= 0
}
// @leet end

// Keynold

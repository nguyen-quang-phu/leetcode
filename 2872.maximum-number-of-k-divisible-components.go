package leetcode

// @leet start
const MAXN = 30000
var adj [MAXN]uint16
var deg [MAXN]uint8
var stack [MAXN/2]uint16

func iszeromask(v uint32) uint32 {
	return ((v | -v) >> 31) - 1
}

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	if n == 1 { // this one doesn't count
		return 1
	}
	for i := range n {
		adj[i] = 0
		deg[i] = 0
	}
	for _, e := range edges {
		adj[e[0]] ^= uint16(e[1])
		adj[e[1]] ^= uint16(e[0])
		deg[e[0]]++
		deg[e[1]]++
	}
	var s uint16
    var result uint32
	pushstack := func(i uint16) {
		mask := uint16(iszeromask(uint32(deg[i] ^ 1)))
		stack[s] = (stack[s] &^ mask) | (i & mask)
		s += 1 & mask
	}
	for i := range uint16(n) {
		pushstack(i)
	}
	for s != 0 {
		s--
		curr := stack[s]
		next := adj[curr]
		adj[next] ^= curr
		deg[next]--
		pushstack(next)
		values[curr] %= k
		result += 1 & iszeromask(uint32(values[curr]))
		values[next] += values[curr]
	}
	return int(result)
}
// @leet end

// Keynold

package leetcode

// @leet start
func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1_000_000_007
	knows := make([]int64, n)
	knows[0] = 1
	shared, total := int64(0), int64(1)

	for day := delay; day < forget && day < n; day++ {
		shared = (shared + knows[day-delay]) % mod
		total = (total + shared) % mod
		knows[day] = shared
	}

	for day := forget; day < n; day++ {
		leavers := knows[day-forget]
		newSharers := knows[day-delay]

		shared = (shared + newSharers - leavers + mod) % mod
		total = (total + shared - leavers + mod) % mod
		knows[day] = shared
	}

	return int((total + mod) % mod)
}
// @leet end

// Keynold

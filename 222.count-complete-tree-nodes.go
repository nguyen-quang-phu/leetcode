package leetcode

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return count(root.Left) + count(root.Right) + 1
}

func countNodes(root *TreeNode) int {
	return count(root)
}

// @leet end

// Keynold

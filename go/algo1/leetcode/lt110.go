package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, ans := depth110(root)
	return ans
}

func depth110(root *TreeNode) (dp int, balanced bool) {
	if root == nil {
		return 0, true
	}
	ld, lb := depth110(root.Left)
	rd, rb := depth110(root.Right)

	return max(ld, rd) + 1, lb && rb && (max(ld, rd)-min(ld, rd) <= 1)
}

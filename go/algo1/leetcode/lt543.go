package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans
}
func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	ld := depth(root.Left, ans)
	rd := depth(root.Right, ans)
	*ans = max(*ans, ld+rd+1)
	return max(ld, rd) + 1
}

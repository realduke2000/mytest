package leetcode

import (
	"fmt"
)

func height655(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	leftHeight := 0
	if root.Left != nil {
		leftHeight = height655(root.Left)
	}
	rightHeight := 0
	if root.Right != nil {
		rightHeight = height655(root.Right)
	}
	return max(leftHeight, rightHeight) + 1 
}

func printChild(r,c,h int, res [][]string, node *TreeNode ) {
	if node.Left != nil {
		_r := r+1
		_c := c-(1 << (h-r-1))
		res[_r][_c] = fmt.Sprintf("%d", node.Left.Val)
		printChild(_r,_c,h,res, node.Left)
	}
	if node.Right != nil {
		_r := r+1
		_c := c+(1 << (h-r-1))
		res[_r][_c] = fmt.Sprintf("%d", node.Right.Val)
		printChild(_r,_c,h,res, node.Right)
	}
}

func printTree(root *TreeNode) [][]string {
    h := height655(root)
	m := h+1
	n := (1<< (h+1))-1
	res := make([][]string, m)
	for i :=0;i<m;i++ {
		res[i] = make([]string, n)
	}

	res[0][(n-1)/2] = fmt.Sprintf("%d",  root.Val)
	printChild(0, (n-1)/2, h, res, root)
	return res
}
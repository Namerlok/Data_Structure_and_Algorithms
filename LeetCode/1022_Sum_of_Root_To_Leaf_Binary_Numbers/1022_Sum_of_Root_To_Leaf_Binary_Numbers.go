package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return subSumRootToLeaf(root, 0)
}

func subSumRootToLeaf(root *TreeNode, prefix int) int {
	var left, right int
	if root.Left != nil {
		left = subSumRootToLeaf(root.Left, prefix*2+root.Val)
	}
	if root.Right != nil {
		right = subSumRootToLeaf(root.Right, prefix*2+root.Val)
	}
	if root.Left == nil && root.Right == nil {
		return prefix*2 + root.Val
	}
	return left + right
}

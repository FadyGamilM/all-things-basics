package main

import "math"

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if math.Abs(float64(leftDepth)-float64(rightDepth)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

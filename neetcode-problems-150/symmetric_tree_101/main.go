package main

import "log"

func main() {
	a := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
	}

	b := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	c := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	log.Println(isBalanced(a, b))
	log.Println(isBalanced(c, b))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isBalanced(a.Left, b.Left) && isBalanced(a.Right, b.Right)
}

func isSymmetricallyBalanced(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isBalanced(a.Left, b.Right) && isBalanced(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	return isBalanced(root.Left, root.Right)
}

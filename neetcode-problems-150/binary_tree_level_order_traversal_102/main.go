package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)
	tries := 1
	for len(q) > 0 {
		levelNodes := []*TreeNode{}
		levelRes := []int{}

		for tries > 0 {
			poppedItem := q[0]
			updatedQ := q[1:]
			q = updatedQ
			levelNodes = append(levelNodes, poppedItem)
			tries--
		}

		for _, node := range levelNodes {
			if node.Left != nil {
				tries++
				q = append(q, node.Left)
			}
			if node.Right != nil {
				tries++
				q = append(q, node.Right)
			}
		}

		for _, node := range levelNodes {
			levelRes = append(levelRes, node.Val)
		}

		res = append(res, levelRes)
	}

	return res
}

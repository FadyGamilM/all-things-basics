package main

import (
	"all-things-basics/data-structers/queue"
	"all-things-basics/data-structers/stack"
	"log"
)

func main() {
	t1 := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 3,
			},
			right: &Node{
				val: 4,
			},
		},
		right: &Node{
			val: 5,
			left: &Node{
				val: 6,
			},
			right: &Node{
				val: 7,
			}},
	}

	log.Println("========= Pre Order ========")
	pre_order_traversal(t1)
	log.Println("========= In Order ========")
	in_order_traversal(t1)
	log.Println("========= Post Order ========")
	post_order_traversal(t1)

	log.Println("======== BFS ========")
	breadth_first_traversal(t1)
	log.Println(breadth_first_search(t1, 4))

	log.Println("======== DFS =======")
	log.Println(compare_two_trees_values_and_structure(
		&Node{
			val: 1,
			left: &Node{
				val: 2,
			},
			right: &Node{
				val: 3,
			},
		},
		&Node{
			val: 1,
			left: &Node{
				val: 2,
				left: &Node{
					val: 3,
				},
			},
		},
	))
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func pre_order_traversal(root *Node) {
	if root == nil {
		return
	}

	// process
	log.Println(root.val)

	// recurse left
	pre_order_traversal(root.left)

	// recurse right
	pre_order_traversal(root.right)
}

func in_order_traversal(root *Node) {
	if root == nil {
		return
	}

	// recurse left
	in_order_traversal(root.left)

	// process
	log.Println(root.val)

	// recurse right
	in_order_traversal(root.right)
}

func post_order_traversal(root *Node) {
	if root == nil {
		return
	}

	// recurse left
	post_order_traversal(root.left)

	// recurse right
	post_order_traversal(root.right)

	// process
	log.Println(root.val)
}

func breadth_first_search(root *Node, needle int) bool {
	if root == nil {
		return false
	}

	storage := queue.Queue[*Node]{}
	storage.Items = []*Node{root}
	storage.Count++
	for storage.Count > 0 {
		// pop the item
		node := storage.Pop()
		if node.val == needle {
			return true
		}
		// push left then right to storage
		if node.left != nil {
			storage.Push(node.left)
		}

		if node.right != nil {
			storage.Push(node.right)
		}
	}
	return false
}

func breadth_first_traversal(root *Node) {
	if root == nil {
		return
	}

	storage := queue.Queue[*Node]{}
	storage.Items = []*Node{root}
	storage.Count++
	for storage.Count > 0 {
		// pop the item
		node := storage.Pop()
		log.Println(node)
		// push left then right to storage
		if node.left != nil {
			storage.Push(node.left)
		}

		if node.right != nil {
			storage.Push(node.right)
		}
	}
}

func compare_two_trees_values_and_structure(root1, root2 *Node) bool {
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}

	if root1 == nil && root2 == nil {
		return true
	}

	s1 := stack.Stack[*Node]{}
	s1.Push(root1)

	s2 := stack.Stack[*Node]{}
	s2.Push(root2)

	for s1.Count > 0 && s2.Count > 0 {
		n1 := s1.Pop()
		n2 := s2.Pop()
		if n1.val != n2.val {
			return false
		}

		if n1.left != nil {
			s1.Push(n1.left)
		}

		if n1.right != nil {
			s1.Push(n1.right)
		}

		if n2.left != nil {
			s2.Push(n2.left)
		}

		if n2.right != nil {
			s2.Push(n2.right)
		}
	}

	if s1.Count > 0 || s2.Count > 0 {
		return false
	}

	return true
}

func binary_search_tree(root, needle *Node) bool {
	if root == nil || needle == nil {
		return false
	}

	if root.val == needle.val {
		return true
	}

	if root.val < needle.val {
		return binary_search_tree(root.right, needle)
	}

	return binary_search_tree(root.left, needle) 
}

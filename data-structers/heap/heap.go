package heap

import (
	"math"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	Items []T
}

func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		Items: []T{},
	}
}

func (h *MinHeap[T]) Parent(i int) int {
	return int(math.Floor(float64((i - 1)) / float64(2)))
}

func (h *MinHeap[T]) LeftChild(i int) int {
	return int(math.Floor((float64(2*i) + float64(1))))
}

func (h *MinHeap[T]) RightChild(i int) int {
	return int(math.Floor((float64(2*i) + float64(2))))
}

// when we need to insert, we just append to the last index of the array, and then keep hepifying with the parent untill we find a parent that we are >= it
func (h *MinHeap[T]) heapifyUp(newNodeIdx int) {
	if newNodeIdx == 0 {
		return
	}

	parentIdx := h.Parent(newNodeIdx)
	parentVal := h.Items[parentIdx]
	newNodeVal := h.Items[newNodeIdx]

	if newNodeVal < parentVal {
		// swap and heapify up
		h.Items[parentIdx] = newNodeVal
		h.Items[newNodeIdx] = parentVal
		h.heapifyUp(parentIdx)
	}
}

// when we need to delete, we first remove the first node (the min of the heap and its the root) => get the last item in the array and set it as the new root => then hepify this node down
func (h *MinHeap[T]) heapifyDown(nodeIdx int) {
	if nodeIdx >= len(h.Items) {
		// we reached the bottom of the heap
		return
	}
	nodeVal := h.Items[nodeIdx]

	leftIndx := h.LeftChild(nodeIdx)
	var leftVal T
	if leftIndx < len(h.Items) {
		leftVal = h.Items[leftIndx]
	}

	rightIndx := h.RightChild(nodeIdx)
	var rightval T
	if rightIndx < len(h.Items) {
		rightval = h.Items[rightIndx]
	}

	// if the right is smaller than the left and we are larger than this right, so lets heapify down with it
	if leftVal > rightval && nodeVal > rightval {
		h.Items[nodeIdx] = rightval
		h.Items[rightIndx] = nodeVal
		h.heapifyDown(rightIndx)
	} else if leftVal < rightval && nodeVal > leftVal {
		h.Items[nodeIdx] = leftVal
		h.Items[leftIndx] = nodeVal
		h.heapifyDown(leftIndx)
	}

}

func (h *MinHeap[T]) Insert(node T) {
	h.Items = append(h.Items, node)
	h.heapifyUp(len(h.Items) - 1)
}

func (h *MinHeap[T]) Delete() T {
	if len(h.Items) == 0 {
		var zero T
		return zero
	}

	if len(h.Items) == 1 {
		item := h.Items[0]
		h.Items = []T{}
		return item
	}

	out := h.Items[0]
	h.Items[0] = h.Items[len(h.Items)-1]
	updatedItems := h.Items[:len(h.Items)-1]
	h.Items = updatedItems
	h.heapifyDown(0)
	return out
}

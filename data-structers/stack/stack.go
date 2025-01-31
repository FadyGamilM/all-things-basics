package stack

type Stack[T any] struct {
	Count int
	Items []T
}

func (s *Stack[T]) Push(item T) {
	s.Count++
	s.Items = append(s.Items, item)
}

// [1, 2, 3]
// when we pop we should pop 3, then 2, then finally 1
func (s *Stack[T]) Pop() T {
	item := s.Items[s.Count-1]
	updatedItems := s.Items[:s.Count-1]
	s.Items = updatedItems
	s.Count--
	return item
}

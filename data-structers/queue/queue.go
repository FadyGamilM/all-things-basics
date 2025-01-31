package queue

import "sync"

type Queue[T any] struct {
	Count int
	Items []T
	Mutex *sync.Mutex
}

func (q *Queue[T]) Push(item T) {
	q.Count++
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Pop() T {
	q.Count--
	item := q.Items[0]
	updatedQueueItems := q.Items[1:]
	q.Items = updatedQueueItems
	return item
}

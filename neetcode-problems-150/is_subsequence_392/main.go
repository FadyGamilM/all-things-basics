package main

import "log"

type Qu struct {
	items []byte
}

func (q *Qu) push(item byte) {
	q.items = append(q.items, item)
}

func (q *Qu) pop() byte {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Qu) first() byte {
	return q.items[0]
}

func isSubsequence(s string, t string) bool {
	q := &Qu{items: []byte{}}
	tracker := map[byte]int{}
	for i := range s {
		q.push(s[i])
		tracker[s[i]] += 1
	}

	for i := range t {
		if len(q.items) == 0 {
			break
		}
		tItem := t[i]
		if val, exists := tracker[tItem]; exists {
			firstItem := q.first()
			if tItem != firstItem {
				continue
			} else {
				q.pop()
			}
			if val == 0 {
				delete(tracker, tItem)
			} else {
				tracker[tItem] -= 1
				if tracker[tItem] == 0 {
					delete(tracker, tItem)
				}
			}
		}
	}

	if len(q.items) > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	s := "ab"
	t := "baab"
	log.Println(isSubsequence(s, t))
}

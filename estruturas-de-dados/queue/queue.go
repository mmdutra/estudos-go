package queue

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end of the queue
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Dequeue removes the fist item of the queue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

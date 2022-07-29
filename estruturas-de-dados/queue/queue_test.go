package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAddItemToTheQueue(t *testing.T) {
	myQueue := Queue{}

	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)

	items := []int{10, 20, 30}

	assert.Equal(t, items, myQueue.items)
}

func TestShouldRemoveItemOfTheQueue(t *testing.T) {
	myQueue := Queue{}

	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	myQueue.Dequeue()

	items := []int{20, 30}

	assert.Equal(t, items, myQueue.items)
}

package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAddItemToStack(t *testing.T) {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	items := []int{10, 20, 30}

	assert.Equal(t, items, stack.items)
}

func TestShouldPopItemsFromStack(t *testing.T) {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Pop()

	items := []int{10, 20}

	assert.Equal(t, items, stack.items)
}

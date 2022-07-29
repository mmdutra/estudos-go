package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAddTheNodeAtTheBegging(t *testing.T) {
	linkedlist := Linkedlist{}

	linkedlist.Insert(10)
	linkedlist.Insert(20)
	linkedlist.Insert(30)

	result := linkedlist.Print()

	assert.Equal(t, "30 20 10 ", result)
}

func TestShouldRemoveANodeFromList(t *testing.T) {
	linkedlist := Linkedlist{}

	linkedlist.Insert(10)
	linkedlist.Insert(20)
	linkedlist.Insert(30)
	linkedlist.Remove(20)

	result := linkedlist.Print()
	assert.Equal(t, "30 10 ", result)

	assert.Equal(t, "Node not found", linkedlist.Remove(40).Error())
}

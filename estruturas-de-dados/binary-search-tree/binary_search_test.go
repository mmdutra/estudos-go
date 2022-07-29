package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAddDataToTheTree(t *testing.T) {
	node := Node{Key: 10}

	node.Insert(45)
	node.Insert(59)
	node.Insert(23)
	node.Insert(90)

	assert.Equal(t, "10 45 23 90 ", Print(&node))
}

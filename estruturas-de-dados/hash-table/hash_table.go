package hashtable

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 5

// Hashtable will hold an array
type Hashtable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the hashtable
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

func (b *bucket) insert(k string) {

	if b.search(k) {
		fmt.Println("Already exists")
	} else {
		newNode := &bucketNode{key: k}

		newNode.next = b.head
		b.head = newNode
	}
}

func (b *bucket) search(k string) bool {
	curr := b.head
	for curr != nil {
		if curr.key == k {
			return true
		}
		curr = curr.next
	}

	return false
}

func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prev := b.head
	for prev.next != nil {
		if prev.next.key == k {
			prev.next = prev.next.next
		}
		prev = prev.next
	}
}

// Insert will take a key and add it to the hash table array
func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take a key and return true if that key exists in the hash table
func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take a key and delete it from the the hash table
func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func Init() *Hashtable {
	result := &Hashtable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

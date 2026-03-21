package doubly_linked_list

import "errors"

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrValueNotFound   = errors.New("value not found")
	ErrEmptyList       = errors.New("list is empty")
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Prepend inserts a new node at the head in O(1) time.
func (d *DoublyLinkedList[T]) Prepend(val T) {
	newNode := &Node[T]{Value: val}
	if d.size == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.Next = d.head
		d.head.Prev = newNode
		d.head = newNode
	}
	d.size++
}

// Append inserts a new node at the tail in O(1) time.
func (d *DoublyLinkedList[T]) Append(val T) {
	newNode := &Node[T]{Value: val}
	if d.size == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.Prev = d.tail
		d.tail.Next = newNode
		d.tail = newNode
	}
	d.size++
}

// InsertAt inserts a node at the specified index in O(N) time.
func (d *DoublyLinkedList[T]) InsertAt(val T, position int) error {
	if position < 0 || position > d.size {
		return ErrIndexOutOfRange
	}
	if position == 0 {
		d.Prepend(val)
		return nil
	}
	if position == d.size {
		d.Append(val)
		return nil
	}
	newNode := &Node[T]{Value: val}
	curr := d.head
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	newNode.Prev = curr
	curr.Next.Prev = newNode
	curr.Next = newNode
	d.size++
	return nil
}

// RemoveAt removes the node at the specified index in O(N) time.
func (d *DoublyLinkedList[T]) RemoveAt(position int) error {
	if position < 0 || position >= d.size {
		return ErrIndexOutOfRange
	}
	if d.size == 0 {
		return ErrEmptyList
	}
	
	if d.size == 1 {
		d.head = nil
		d.tail = nil
		d.size--
		return nil
	}

	if position == 0 {
		d.head = d.head.Next
		d.head.Prev = nil
		d.size--
		return nil
	}

	if position == d.size-1 {
		d.tail = d.tail.Prev
		d.tail.Next = nil
		d.size--
		return nil
	}

	curr := d.head
	for i := 0; i < position; i++ {
		curr = curr.Next
	}
	
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	d.size--
	return nil
}

// Remove deletes the first occurrence of val in O(N) time.
func (d *DoublyLinkedList[T]) Remove(val T) error {
	if d.size == 0 {
		return ErrEmptyList
	}
	
	curr := d.head
	for curr != nil && curr.Value != val {
		curr = curr.Next
	}
	
	if curr == nil {
		return ErrValueNotFound
	}
	
	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	} else {
		d.head = curr.Next
	}
	
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	} else {
		d.tail = curr.Prev
	}
	d.size--
	return nil
}

// Reverse iteratively reverses the doubly linked list in O(N) time, O(1) space.
func (d *DoublyLinkedList[T]) Reverse() {
	if d.size <= 1 {
		return
	}
	
	curr := d.head
	d.tail = d.head
	var prev *Node[T]
	
	for curr != nil {
		prev = curr.Prev
		curr.Prev = curr.Next
		curr.Next = prev
		curr = curr.Prev // advance since next and prev are swapped
	}
	
	if prev != nil {
		d.head = prev.Prev
	}
}

// Size returns the count of elements in the list.
func (d *DoublyLinkedList[T]) Size() int { return d.size }

// Values returns a slice of all values mapped iteratively.
func (d *DoublyLinkedList[T]) Values() []T {
	res := make([]T, 0, d.size)
	for curr := d.head; curr != nil; curr = curr.Next {
		res = append(res, curr.Value)
	}
	return res
}

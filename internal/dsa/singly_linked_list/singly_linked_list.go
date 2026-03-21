package singly_linked_list

import (
	"errors"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrValueNotFound   = errors.New("value not found")
	ErrEmptyList       = errors.New("list is empty")
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Prepend inserts a new node at the head of the list in O(1) time.
func (s *SinglyLinkedList[T]) Prepend(val T) {
	newNode := &Node[T]{Value: val, Next: s.head}
	s.head = newNode
	if s.size == 0 {
		s.tail = newNode
	}
	s.size++
}

// Append inserts a new node at the tail of the list in O(1) time.
func (s *SinglyLinkedList[T]) Append(val T) {
	newNode := &Node[T]{Value: val}
	if s.size == 0 {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.Next = newNode
		s.tail = newNode
	}
	s.size++
}

// InsertAt inserts a new node at a specific zero-indexed position in O(N) time.
func (s *SinglyLinkedList[T]) InsertAt(val T, position int) error {
	if position < 0 || position > s.size {
		return ErrIndexOutOfRange
	}
	if position == 0 {
		s.Prepend(val)
		return nil
	}
	if position == s.size {
		s.Append(val)
		return nil
	}
	newNode := &Node[T]{Value: val}
	curr := s.head
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	curr.Next = newNode
	s.size++
	return nil
}

// RemoveAt removes the node at a specific zero-indexed position in O(N) time.
func (s *SinglyLinkedList[T]) RemoveAt(position int) error {
	if position < 0 || position >= s.size {
		return ErrIndexOutOfRange
	}
	if s.size == 0 {
		return ErrEmptyList
	}
	if position == 0 {
		s.head = s.head.Next
		s.size--
		if s.size == 0 {
			s.tail = nil
		}
		return nil
	}
	
	curr := s.head
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	if position == s.size-1 {
		s.tail = curr
	}
	s.size--
	return nil
}

// Remove deletes the first occurrence of the specified value in O(N) time.
func (s *SinglyLinkedList[T]) Remove(val T) error {
	if s.size == 0 {
		return ErrEmptyList
	}
	
	if s.head.Value == val {
		s.head = s.head.Next
		s.size--
		if s.size == 0 {
			s.tail = nil
		}
		return nil
	}

	curr := s.head
	for curr.Next != nil && curr.Next.Value != val {
		curr = curr.Next
	}

	if curr.Next == nil {
		return ErrValueNotFound
	}

	curr.Next = curr.Next.Next
	if curr.Next == nil {
		s.tail = curr
	}
	s.size--
	return nil
}

// Reverse iteratively reverses the linked list in place in O(N) time and O(1) space.
func (s *SinglyLinkedList[T]) Reverse() {
	if s.size <= 1 {
		return
	}
	var prev *Node[T]
	curr := s.head
	s.tail = s.head // The new tail will be the old head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	s.head = prev
}

// HasCycle detects whether the linked list has a cycle using Floyd's Tortoise and Hare algorithm.
// Time: O(N), Space: O(1).
func (s *SinglyLinkedList[T]) HasCycle() bool {
	if s.head == nil {
		return false
	}
	slow, fast := s.head, s.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Size returns the count of elements in the list.
func (s *SinglyLinkedList[T]) Size() int { return s.size }

// Values returns a slice of all values present in the linked list sequentially.
func (s *SinglyLinkedList[T]) Values() []T {
	res := make([]T, 0, s.size)
	curr := s.head
	for curr != nil {
		res = append(res, curr.Value)
		curr = curr.Next
	}
	return res
}

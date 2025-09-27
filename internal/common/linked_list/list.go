package linked_list

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type List[T comparable] interface {
	Append(value T)
	Prepend(value T)
	Remove(index int) error
	Find(value T) (int, error)
	Length() int
	Validate() error
	Head() *Node[T]
	Tail() *Node[T]
}

type NonValidatedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (l *NonValidatedList[T]) Append(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  nil,
		Prev:  l.tail,
	}
	l.size++

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.Next = newNode
	l.tail = newNode
}

func (l *NonValidatedList[T]) Prepend(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  l.head,
	}
	l.size++

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.head.Prev = newNode
	l.head = newNode

}

func (l *NonValidatedList[T]) Remove(index int) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("index %d does not exist in the list of size %d", index, l.size)
	}

	// Removing the head node.
	if index == 0 {
		removed := l.head
		l.head = removed.Next
		if l.head != nil {
			l.head.Prev = nil
		} else {
			// List becomes empty; update tail as well.
			l.tail = nil
		}
		l.size--
		return nil
	}

	// Removing the tail node.
	if index == l.size-1 {
		removed := l.tail
		l.tail = removed.Prev
		if l.tail != nil {
			l.tail.Next = nil
		}
		l.size--
		return nil
	}

	// Removing a node from the middle.
	current := l.head
	for range index {
		current = current.Next
	}
	// Update the previous and next nodes to bypass the current one.
	if current.Prev != nil {
		current.Prev.Next = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
	l.size--
	return nil
}

func (l *NonValidatedList[T]) Find(value T) (int, error) {
	current := l.head
	for i := range l.size {
		if current.Value == value {
			return i, nil
		}
		current = current.Next
	}
	return -1, fmt.Errorf("value %v not found in the list", value)
}

func (l *NonValidatedList[T]) Length() int {
	return l.size
}

func (l *NonValidatedList[T]) Validate() error {
	if l.size == 0 {
		return nil // Empty list is valid
	}

	// Validate size consistency
	actualSize := 0
	current := l.head
	for current != nil {
		actualSize++
		if actualSize > l.size {
			return fmt.Errorf("list size inconsistent: actual size exceeds recorded size %d", l.size)
		}
		current = current.Next
	}

	if actualSize != l.size {
		return fmt.Errorf("list size inconsistent: recorded size %d, actual size %d", l.size, actualSize)
	}

	// Validate head/tail consistency
	if l.size == 1 {
		if l.head != l.tail {
			return fmt.Errorf("single-item list head/tail inconsistency")
		}
		if l.head.Next != nil || l.head.Prev != nil {
			return fmt.Errorf("single-item list should have nil Next/Prev pointers")
		}
	}

	// Validate forward/backward link consistency
	current = l.head
	for current != nil && current.Next != nil {
		if current.Next.Prev != current {
			return fmt.Errorf("forward/backward link inconsistency detected")
		}
		current = current.Next
	}

	// Validate that we end at tail
	if current != l.tail {
		return fmt.Errorf("forward traversal doesn't end at tail")
	}

	return nil
}

func (b *NonValidatedList[T]) Head() *Node[T] {
	return b.head
}

func (b *NonValidatedList[T]) Tail() *Node[T] {
	return b.tail
}

func NewList[T comparable]() List[T] {
	return &NonValidatedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

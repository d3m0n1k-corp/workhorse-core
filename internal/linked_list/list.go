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

type baseList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (l *baseList[T]) Append(value T) {
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

}

func (l *baseList[T]) Prepend(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  l.head,
	}

	l.head = newNode
	l.size++
}

func (l *baseList[T]) Remove(index int) error {
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
	for i := 0; i < index; i++ {
		current = current.Next
	}
	if current == nil {
		return fmt.Errorf("unexpected nil node at index %d", index)
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

func (l *baseList[T]) Find(value T) (int, error) {
	current := l.head
	for i := range l.size {
		if current.Value == value {
			return i, nil
		}
		current = current.Next
	}
	return -1, fmt.Errorf("value %v not found in the list", value)
}

func (l *baseList[T]) Length() int {
	return l.size
}

func (l *baseList[T]) Validate() error {
	panic("TODO: Implement")
}

func (b *baseList[T]) Head() *Node[T] {
	return b.head
}

func (b *baseList[T]) Tail() *Node[T] {
	return b.tail
}

func NewList[T comparable]() List[T] {
	return &baseList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

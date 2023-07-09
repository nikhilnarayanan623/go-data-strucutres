package slist

import (
	"errors"
)

// for generic
type dataType interface {
	int | string | float64
}

var (
	ErrInvalidPosition = errors.New("invalid position")
	ErrEmptyList       = errors.New("singly list is empty")
	ErrDataNotExist    = errors.New("data not exist")
)

type node[T dataType] struct {
	data T
	next *node[T]
}

type SinglyLinkedList[T dataType] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// To get a new singly linked list with generics of (int, string, float64)
func NewSinglyLinkedList[T dataType]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// To get the length of the singly linked list
func (c *SinglyLinkedList[T]) Len() int {
	return c.length
}

// To check the singly linked list is empty
func (c *SinglyLinkedList[T]) IsEmpty() bool {
	return c.length == 0
}

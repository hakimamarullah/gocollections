package golinkedlist

import (
	"github.com/hakimamarullah/gocollections/golinkedlist/domain"
	"github.com/hakimamarullah/gocollections/golinkedlist/interfaces"
)

// NewLinkedList constructor of the LinkedList with type List.
func NewLinkedList[T any]() interfaces.List[T] {
	return &domain.LinkedList[T]{Head: nil, Tail: nil}
}

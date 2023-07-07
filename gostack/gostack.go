// Package gostack Go stack data structure implementation
package gostack

import (
	"github.com/hakimamarullah/gocollections/gostack/domain"
	"github.com/hakimamarullah/gocollections/gostack/interfaces"
)

// NewStack constructor to create Stack.
// Returns stack of type GoStack.
func NewStack[T any]() interfaces.GoStack[T] {
	return &domain.Stack[T]{Head: nil, StackSize: 0}
}

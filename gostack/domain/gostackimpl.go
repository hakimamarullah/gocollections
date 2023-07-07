package domain

import (
	"errors"
	"fmt"
	"github.com/hakimamarullah/gocollections/gostack/utils"
	"reflect"
	"strings"
)

// NewNode constructor to create Node.
// Returns pointer of Node.
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data, Next: nil}
}

// Push data to stack and put in on top
func (S *Stack[T]) Push(data T) {
	newNode := NewNode(data)
	if S.Head == nil {
		S.Head = newNode
		S.StackSize++
		return
	}
	newNode.Next = S.Head
	S.Head = newNode
	S.StackSize++

}

// Pop Get first data from the top of stack and remove it
func (S *Stack[T]) Pop() (T, error) {
	if S.Head == nil {
		return zeroValue[T](), errors.New("empty stack")
	}
	data := S.Head

	S.Head = S.Head.Next
	S.StackSize--
	return data.Data, nil
}

// zeroValue obtain zero value for generic type.
func zeroValue[T any]() T {
	return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
}

// Peek Get first data from the top of stack without removing the data
func (S *Stack[T]) Peek() (T, error) {
	if S.Head == nil {
		return zeroValue[T](), errors.New("empty stack")
	}
	data := S.Head
	return data.Data, nil
}

// Print the stack from the top to bottom
func (S *Stack[T]) Print() {
	current := S.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}

	fmt.Println()
}

// IsEmpty Check whether the stack is empty or not
func (S *Stack[T]) IsEmpty() bool {
	return S.Head == nil
}

// Size Return Stack Size
func (S *Stack[T]) Size() uint64 {
	return S.StackSize
}

// Clear Remove all stack element
func (S *Stack[T]) Clear() {
	S.Head = nil
	S.StackSize = 0
}

// Search Returns 1-based position of the stack.
// The position is relative from the top of the stack start with 1.
// Returns -1 if the element doesn't exist in the stack.
func (S *Stack[T]) Search(data T) int64 {
	position := 0
	currentNode := S.Head
	for currentNode != nil {
		position++

		if same, kind := utils.IsSameKind(currentNode.Data, data); same {
			if kind.Name() == "string" {
				x := reflect.ValueOf(currentNode.Data).Interface().(string)
				y := reflect.ValueOf(data).Interface().(string)

				if strings.EqualFold(x, y) {
					return int64(position)
				}

			}
		}
		if reflect.DeepEqual(currentNode.Data, data) {
			return int64(position)
		}
		currentNode = currentNode.Next
	}

	return -1
}

// String implement toString method for Stack.
func (S *Stack[T]) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString("[")

	current := S.Head
	for current != nil {
		if current.Next != nil {
			stringBuilder.WriteString(fmt.Sprintf("%v, ", current.Data))
		} else {
			stringBuilder.WriteString(fmt.Sprintf("%v", current.Data))
		}
		current = current.Next
	}
	stringBuilder.WriteString("]")
	return stringBuilder.String()
}

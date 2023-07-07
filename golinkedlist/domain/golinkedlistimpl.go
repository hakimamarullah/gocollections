package domain

import (
	"errors"
	"fmt"
	"github.com/hakimamarullah/gocollections/golinkedlist/utils"
	"math"
	"reflect"
	"strings"
)

// NewNode constructor for Node.
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data}
}

// Add method to add the data at the end of the list.
func (ll *LinkedList[T]) Add(data T) {
	newNode := NewNode[T](data)
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		ll.ListSize++
		return
	}

	ll.Tail.Next = newNode
	newNode.Prev = ll.Tail
	ll.Tail = newNode
	ll.ListSize++
}

// AddFirst method to add the element at the first position in the list.
func (ll *LinkedList[T]) AddFirst(data T) {
	newNode := NewNode[T](data)
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		ll.ListSize++
		return
	}

	newNode.Next = ll.Head
	ll.Head.Prev = newNode
	ll.Head = newNode
	ll.ListSize++
}

// Get method to get the element by index.
// index could be positive or negative integer.
// if the index is negative then it will start indexing the list from last element.
// the last element of the list has index -1 if accessed from the back of the list.
// Returns (data, node, error)
func (ll *LinkedList[T]) Get(index int) (*T, *Node[T], error) {
	var current *Node[T]

	if ll.Head == nil {
		return nil, nil, errors.New("list empty")
	}

	if index >= int(ll.ListSize) {
		return nil, nil, errors.New(fmt.Sprintf("Index %v out of bound of list size %v", index, ll.ListSize))
	}

	if index < 0 {
		if math.Abs(float64(index)) > float64(ll.ListSize) {
			return nil, nil, errors.New(fmt.Sprintf("Index %v out of bound of list size %v", index, ll.ListSize))
		}
		current = ll.Tail
		for i := 0; i < int(math.Abs(float64(index))); i++ {
			current = current.Prev
		}
	} else {
		current = ll.Head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	}

	return &current.Data, current, nil
}

// Contains checks if the list contains the element given in the parameter.
// Please note that the Contains method is case insensitive for string type.
func (ll *LinkedList[T]) Contains(data T) bool {
	current := ll.Head
	for current != nil {
		if same, kind := utils.IsSameKind(current.Data, data); same {
			if kind.Name() == "string" {
				x := reflect.ValueOf(current.Data).Interface().(string)
				y := reflect.ValueOf(data).Interface().(string)

				if strings.EqualFold(x, y) {
					return true
				}

			}
		}
		if reflect.DeepEqual(current.Data, data) {
			return true
		}
		current = current.Next
	}
	return false
}

// Insert method to insert data at the specified index.
// index should be positive and can't be larger than listSize-1.
// Returns true if the insert operation is successful otherwise false and error will be returned
// instead.
func (ll *LinkedList[T]) Insert(data T, index int) (bool, error) {
	if index == 0 {
		ll.AddFirst(data)
		return true, nil
	}

	if index >= int(ll.ListSize) {
		return false, errors.New(fmt.Sprintf("Index %v out of bound of list size %v", index, ll.ListSize))
	}

	if index > 0 {
		_, node, err := ll.Get(index)
		if err != nil {
			return false, err
		}
		newNode := NewNode[T](data)
		newNode.Next = node
		newNode.Prev = node.Prev
		node.Prev.Next = newNode
		node.Prev = newNode
		ll.ListSize++
		return true, nil
	}

	return false, errors.New("negative index not allowed")
}

// Remove method to remove the given element from the list.
// Please note that the Remove method is case-sensitive for the string type.
// this method returns boolean if the remove operation is successful.
func (ll *LinkedList[T]) Remove(data T) bool {
	if ll.Head == nil {
		return false
	}

	if reflect.DeepEqual(ll.Head.Data, data) {
		ll.Head = ll.Head.Next
		ll.Head.Prev = nil
		ll.ListSize--
		return true
	}

	current := ll.Head.Next
	for current != nil && !reflect.DeepEqual(current.Data, data) {
		current = current.Next
	}

	if current != nil {
		if current.Next == nil {
			current.Prev.Next = nil
			ll.Tail = current.Prev
		} else {
			current.Next.Prev = current.Prev
			current.Prev.Next = current.Next
		}
		ll.ListSize--
		return true
	}
	return false
}

// Clear method to clear all the data in the list.
func (ll *LinkedList[T]) Clear() {
	if ll.Head.Next != nil {
		ll.Head.Next = nil
	}
	ll.Head = nil
	ll.ListSize = 0
}

// Size method to return the size of the list.
func (ll *LinkedList[T]) Size() int {
	return int(ll.ListSize)
}

// String implement toString method for Stack.
func (ll *LinkedList[T]) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString("[")

	current := ll.Head
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

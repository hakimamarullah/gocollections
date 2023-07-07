package domain

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type Stack[T any] struct {
	Head      *Node[T]
	StackSize uint64
}

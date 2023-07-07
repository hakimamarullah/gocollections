package domain

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type LinkedList[T any] struct {
	Head     *Node[T]
	Tail     *Node[T]
	ListSize uint
}

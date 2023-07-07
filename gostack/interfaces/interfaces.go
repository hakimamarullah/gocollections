package interfaces

type GoStack[T any] interface {
	Pop() (T, error)
	Peek() (T, error)
	Push(data T)
	Print()
	IsEmpty() bool
	Size() uint64
	Clear()
	Search(data T) int64
}

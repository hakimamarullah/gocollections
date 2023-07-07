package interfaces

import "github.com/hakimamarullah/gocollections/golinkedlist/domain"

type List[T any] interface {
	Add(data T)
	AddFirst(data T)
	Get(index int) (*T, *domain.Node[T], error)
	Contains(data T) bool
	Insert(data T, index int) (bool, error)
	Remove(data T) bool
	Clear()
	Size() int
}

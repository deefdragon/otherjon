package datastructures

type Collection[T any] interface {
	Add(T)
	Clear()
	IsEmpty() bool
	Size() int
	// GetItterator() ItteratorNode[T]
}

// type ItteratorNode[T any] interface {
// 	HasNext() bool
// 	Next() (ItteratorNode[T], error)
// 	Data() T
// 	InsertBefore(T)
// 	InsertAfter(T)
// 	RemoveBefore() error
// 	RemoveAfter() error
// }
type List[T any] interface {
	Collection[T]
	Set(int, T) error
	RemoveAt(int) error
	AddAt(int, T) error
	Get(int) (T, error)
	GetAsArray() []T
}
type BinaryTree[T any] interface {
	Collection[T]
	GetLeftNode() BinaryTree[T]
	GetRightNode() BinaryTree[T]
	SetLeft(T)
	SetRight(T)
	RemoveLeft()
	RemoveRight()
	Get() T
}

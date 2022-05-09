package datastructures

type Collection[T any] interface {
	Add(T)
	Clear()
	IsEmpty() bool
	Size() int
}

type List[T any] interface {
	Collection[T]
	Set(int, T) error
	RemoveAt(int) error
	AddAt(int, T) error
	Get(int) (T, error)
}

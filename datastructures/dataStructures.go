package datastructures

type Collection[T any] interface {
	Add(T)
	Clear()
	IsEmpty() bool
	Size() int
}

type List[T any] interface {
	Collection[T]
	Set(int, T)
	RemoveAt(int)
	AddAt(int, T)
	Get(int) T
}

package datastructures

var _ List[int] = (*ArrayList[int])(nil)

type ArrayList[T any] struct {
}

func NewArrayList[T any]() *ArrayList[T] {
	return nil
}

// Collection methods
func (a *ArrayList[T]) Add(item T) {
}
func (a *ArrayList[T]) Clear() {
}
func (a *ArrayList[T]) IsEmpty() bool {
	return true
}
func (a *ArrayList[T]) Size() int {
	return 0
}

// List Methods
func (a *ArrayList[T]) Set(pos int, item T) {
}
func (a *ArrayList[T]) RemoveAt(pos int) {
}
func (a *ArrayList[T]) AddAt(pos int, item T) {
}
func (a *ArrayList[T]) Get(pos int) T {
	return *new(T)
}

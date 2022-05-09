package datastructures

var _ List[int] = (*ArrayList[int])(nil)

type ArrayList[T any] struct {
	backer []T
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		backer: make([]T, 0),
	}
}

// Collection methods
func (a *ArrayList[T]) Add(item T) {
	// TODO nutex
	a.backer = append(a.backer, item)
}
func (a *ArrayList[T]) Clear() {
	// TODO nutex
	a.backer = make([]T, 0)
}
func (a *ArrayList[T]) IsEmpty() bool {
	return a.Size() == 0
}
func (a *ArrayList[T]) Size() int {
	return len(a.backer)
}

// List Methods
func (a *ArrayList[T]) Set(pos int, item T) error {
	// TODO nutex
	if pos < 0 || pos >= a.Size() {
		return OutOfBoundsError
	}
	a.backer[pos] = item
	return nil
}
func (a *ArrayList[T]) RemoveAt(pos int) error {
	//TODO nutex
	return nil
}
func (a *ArrayList[T]) AddAt(pos int, item T) error {
	//TODO nutex
	return nil
}
func (a *ArrayList[T]) Get(pos int) (T, error) {
	// TODO nutex
	return *new(T), nil
}

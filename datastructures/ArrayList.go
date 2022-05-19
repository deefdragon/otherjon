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
	if pos >= a.Size() || pos < 0 {
		return OutOfBoundsError
	}
	start := a.backer[:pos]
	end := a.backer[pos+1:]
	for i := 0; len(end) > i; i++ {
		start = append(start, end[i])
	}
	a.backer = start
	return nil
}
func (a *ArrayList[T]) AddAt(pos int, item T) error {
	//TODO nutex
	if pos > a.Size() || pos < 0 {
		return OutOfBoundsError
	}
	start := a.backer[:pos]
	end := a.backer[pos:]
	next := item
	for i := 0; len(end) > i; i++ {
		// old item taken out
		saved := end[i]
		// new item gets put in
		start = append(start, next)
		// new item equals old item
		next = saved
	}
	start = append(start, next)
	a.backer = start
	return nil
}
func (a *ArrayList[T]) Get(pos int) (T, error) {
	// TODO nutex
	if pos >= a.Size() || pos < 0 {
		return *new(T), OutOfBoundsError
	}
	return a.backer[pos], nil
}

// in this getasarray, the format is already array, we only return a.backer
func (a *ArrayList[T]) GetAsArray() []T {
	return a.backer
}

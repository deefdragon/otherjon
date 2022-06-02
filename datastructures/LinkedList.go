package datastructures

var _ List[int] = (*LinkedList[int])(nil)

type LinkedList[T any] struct {
	start *Node[T]
	// end   *Node[T]
	size int
}
type Node[T any] struct {
	data T
	next *Node[T]
	// previous *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Collection methods
func (a *LinkedList[T]) Add(item T) {
	// TODO mutex
	// addAt pos = size-1
	n := &Node[T]{
		data: item,
		next: a.start,
	}
	a.start = n
	a.size++
}
func (a *LinkedList[T]) Clear() {
	// TODO mutex
	a.size = 0
	a.start = nil
	// a.end = nil
}
func (a *LinkedList[T]) IsEmpty() bool {
	return a.size == 0
}
func (a *LinkedList[T]) Size() int {
	return a.size
}

// List Methods
func (a *LinkedList[T]) Set(pos int, item T) error {
	current := a.start
	for i := 0; i < pos; i++ {
		current = current.next
	}
	current.data = item
	return nil
}
func (a *LinkedList[T]) RemoveAt(pos int) error {
	//TODO mutex
	if pos >= a.size || pos < 0 {
		return OutOfBoundsError
	}
	a.size--
	if pos == 0 {
		a.start = a.start.next
		return nil
	}
	current := a.start
	for i := 0; i < pos-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	return nil
}
func (a *LinkedList[T]) AddAt(pos int, item T) error {
	//TODO mutex
	n := &Node[T]{
		data: item,
		next: a.start,
	}
	a.size++
	current := a.start
	if pos == 0 {
		n.next = current
		a.start = n
		return nil
	}
	for i := 0; i < pos-1; i++ {
		current = current.next

	}
	n.next = current.next
	current.next = n

	return nil
}
func (a *LinkedList[T]) Get(pos int) (T, error) {
	// TODO mutex
	current := a.start
	for i := 0; i < pos; i++ {
		current = current.next
	}
	return current.data, nil
}
func (a *LinkedList[T]) GetAsArray() []T {
	arr := []T{}
	nextNode := a.start
	for {
		if nextNode == nil {
			break
		}
		arr = append(arr, nextNode.data)

		nextNode = nextNode.next
	}
	return arr
}

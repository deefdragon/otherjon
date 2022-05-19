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
	// TODO mutex
	// Hint: combine get and add
	return nil
}
func (a *LinkedList[T]) RemoveAt(pos int) error {
	//TODO mutex
	return nil
}
func (a *LinkedList[T]) AddAt(pos int, item T) error {
	//TODO mutex
	return nil
}
func (a *LinkedList[T]) Get(pos int) (T, error) {
	// TODO mutex
	// go from the start, reach the specified position
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
		if len(arr) == a.size  {
			break
		}
		arr = append(arr, nextNode.data)
		nextNode = nextNode.next
	}
	return arr
}

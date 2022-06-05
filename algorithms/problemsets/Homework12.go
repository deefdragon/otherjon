package problemsets

// defines the interface with the functions for our data structure to use
// for the sake of the problem I've decided to make my interface more specific to the methods I write, for explicitness
type D[T any] interface {
	Build(x T)
	InsertAt(pos int, item T)
	DeleteAt(pos int) T
	Reverse(D []T, pos int, numOfItems int)
	Move(D []T, pos1 int, numOfItems int, pos2 int)
}

// Defines an object with our modifiable data/list inside
type List[T any] struct {
	Data []T
}

// Build creates a new array with the inputted item appended to it
func Build[T any](x T) []T {
	return append([]T{}, x)
}

// InsertAt takes in a position and an item then seeks to that position
func (a *List[T]) InsertAt(pos int, item T) {
	start := a.Data[:pos]
	end := a.Data[pos:]
	next := item
	//loop that saves currently indexed item, adds the new item in, then changes the new item to be the saved item
	for i := 0; i < len(end); i++ {
		saved := end[i]
		start = append(start, next)
		next = saved
	}
	// repeats the previous append once more for the last item
	start = append(start, next)
	a.Data = start
}

// DeleteAt takes in a position, deletes the item at the position, shifts everything over, then returns the deleted item
func (a *List[T]) DeleteAt(pos int) T {
	//our variable for storing the deleted item
	deleted := a.Data[pos]
	start := a.Data[:pos]
	end := a.Data[pos+1:]
	for i := 0; i < len(end); i++ {
		start = append(start, end[i])
	}
	return deleted
}

func (a *List[T]) Reverse(D []T, pos int, numOfItems int) {

}
func (a *List[T]) Move(D []T, pos1 int, numOfItems int, pos2 int) {

}

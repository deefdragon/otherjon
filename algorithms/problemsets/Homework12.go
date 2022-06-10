package problemsets

// defines the interface with the functions for our data structure to use
// for the sake of the problem I've decided to make my interface more specific to the methods I write, for explicitness
type D[T any] interface {
	InsertAt(pos int, item T)
	DeleteAt(pos int) T
	Reverse(D []T, pos int, numOfItems int)
	Move(D []T, pos1 int, numOfItems int, pos2 int)
}

// Defines an object with our modifiable data/list inside
type List[T any] struct {
	Data []T
}

/*
InsertAt takes in a position (pos) and an item (item) then seeks to that position using a for loop, saves the item in the current position
that we want to add to, replaces it with our new item, then sets the item variable to be the next data in our line, constantly saving and replacing our data
until we reach the end of the list, where the last piece of data is appended to the end of the list
*/
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

/* DeleteAt takes in a position (pos), saves the data currently at the position we want to delete from,
then seeks to that position using a for loop, shifts over all data AFTER the specified position to the left,
effectively replacing/overwriting the data we want to delete, then returns the removed data in the form of the "deleted" variable
*/
func (a *List[T]) DeleteAt(pos int) T {
	//variable for saving deleted item
	deleted := a.Data[pos]
	//variable declaring what the first half of the data we need is
	start := a.Data[:pos]
	//variable declaring what we want our end of the list to be, pos+1 to ignore the data we're deleting
	end := a.Data[pos+1:]
	// seeks to target position and appends pos+1 to the first half of the list, effectively removing the data at target postion by shifting everything over it
	for i := 0; i < len(end); i++ {
		start = append(start, end[i])
	}
	return deleted
}

/* Reverse takes in a data structure D, a position pos, and a number of effected items numOfItems. First we create a variable save of type T and
create a loop that seeks to target position, then for the position through target number of items, we save the beginning data (at current position) then
the data at the end of the items we want effected (pos + numOfItems) then we swap their places, decrease numOfItems by 1 and incrementing i (pos) by one, moving
closer together until we reach the middle of the list of things we want effected, effectively reversing i through i + k items
*/

func (a *List[T]) Reverse(D []T, pos int, numOfItems int) {
	// variables declared for the first and last item in our subarray
	var saveStart T
	var saveEnd T
	// For loop seeks to target position, incrementing current position until the end of position plus range
	for i := 0; i < pos+numOfItems; i++ {
		// If statement checks for our current position to be in target range
		if i >= pos || i <= pos+numOfItems {
			// Sets saveStart and saveEnd to the first and last positions of our range/subarray respectively
			saveStart = a.DeleteAt(i)
			saveEnd = a.DeleteAt(i + numOfItems)
			// Replaces the data in the first and last positions with each other, then decrements the total number of items by one
			a.InsertAt(i+numOfItems, saveStart)
			a.InsertAt(i, saveEnd)
			numOfItems--
		}
	}
}

/* Move is a method that modifies a subarray within a data structure by taking in a target position as well as a number of effected items, and moving the effected
items to a second target position. We start by creating a for loop to seek to the first target position, while checking to see if we're in our target range.
Once we're in our target range, we delete the data at the current position while inserting that data at position 2, and we do this for each item in our range.
*/
// Move takes in a data structure []T, a first position, a number of effected items, and a second target position to move the range to
func (a *List[T]) Move(D []T, pos1 int, numOfItems int, pos2 int) {
	// First we create a loop that will seek to the target position in our array, up to our position plus number of effected items after the target position
	for i := 0; i < pos1+numOfItems; i++ {
		// Then we use an if statement to check on each iteration of the loop for being in our target range.
		if i >= pos1 || i <= pos1+numOfItems {
			// If we are in our target range, we insert at one item after the target position, while deleting the item at the current position.
			a.InsertAt(pos2+1, a.DeleteAt(i))
		}
	}
}

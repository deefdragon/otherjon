package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Write tests for each of the 8 methods: add, addAt, clear, get, isEmpty, removeAt, set, size

// Collection methods
// Testing Add, Adds a *thing* to our metaphorical list of things, plonks it on the end I think? Need to make sure it adds the thing
func TestAdd(t *testing.T) {
	al := NewArrayList[int]()
	assert.NotNil(t, al)
	assert.Equal(t, 0, len(al.backer))
	al.Add(15)
	assert.Equal(t, 1, len(al.backer))
}

// Testing Clear, Should clear our entire list of things, wiping it to 0/nil/empty
func TestClear(t *testing.T) {
	a := NewArrayList[int]()
	assert.Equal(t, 0, len(a.backer))
	a.Add(1)
	a.Add(2)
	assert.Equal(t, 2, len(a.backer))
	a.Clear()
	assert.Equal(t, 0, len(a.backer))
	a.Clear()
	a.Clear()
	a.Clear()
	assert.Equal(t, 0, len(a.backer))
}

// Testing IsEmpty, is it empty? TRUE, does it have stuff? FALSE, simple enough
func TestIsEmpty(t *testing.T) {
	a := NewArrayList[int]()
	assert.True(t, a.IsEmpty())
	a.Add(27)
	assert.False(t, a.IsEmpty())
	a.Clear()
	assert.True(t, a.IsEmpty())
}

//Testing Size, a method that gets us the size of our list by how many valid things are inside, not including empty buckets, should return a whole number worth of things
func TestSize(t *testing.T) {
	a := NewArrayList[int]()
	assert.Equal(t, 0, a.Size())
	a.Add(1)
	a.Add(2)
	a.Add(3)
	assert.Equal(t, 3, a.Size())
	a.Clear()
	assert.Equal(t, 0, a.Size())
}

// List Methods

// Testing AddAt, Like Add which puts a thing in our collection of things, but specifically adds it at a specific point in our list of stuff. Make sure it adds to the right place.
// AddAt will add an item to the specified index, all items at/after will have their index incremented by 1
// Should the provided index be invalid, returns outofboundserror
// in an index with 3 items, Valid indexes are 0-2, additionally index 3 will put the item on the end(same as add)
// Valid indexes are any indexes which there are currently an item, or the index directly following the last item
func TestAddAt(t *testing.T) {
	a := NewArrayList[int]()

	// empty list adding
	err := a.AddAt(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, []int{1}, a.backer)
	//end of list adding
	err = a.AddAt(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2}, a.backer)
	// add on index of last item
	err = a.AddAt(1, 3)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 3, 2}, a.backer)
	// add to middle of list
	err = a.AddAt(1, 4)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 4, 3, 2}, a.backer)
	// add to beginning of list
	err = a.AddAt(0, 5)
	assert.Nil(t, err)
	assert.Equal(t, []int{5, 1, 4, 3, 2}, a.backer)
	// testing negative oob
	err = a.AddAt(-1, 100)
	assert.NotNil(t, err)
	assert.Equal(t, []int{5, 1, 4, 3, 2}, a.backer)
	assert.Equal(t, OutOfBoundsError, err)
	// testing positive OOB
	err = a.AddAt(6, -100)
	assert.NotNil(t, err)
	assert.Equal(t, []int{5, 1, 4, 3, 2}, a.backer)
	assert.Equal(t, OutOfBoundsError, err)
}

// Testing Get, Retrieves a specific/specified item from our list of things. If it gives us a thing we don't want, test fails
// should not return incorrect data, should return OOB error if index is invalid. Index is valid where any item currently exists
func TestGet(t *testing.T) {
	a := NewArrayList[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	a.Add(5)
	// happy path :D
	obs, err := a.Get(1)
	assert.Equal(t, 2, obs)
	assert.Nil(t, err)
	// happy-low path
	obs, err = a.Get(0)
	assert.Equal(t, 1, obs)
	assert.Nil(t, err)
	// happy high path
	obs, err = a.Get(4)
	assert.Equal(t, 5, obs)
	assert.Nil(t, err)
	// unhappy high OOB
	obs, err = a.Get(5)
	assert.Equal(t, 0, obs)
	assert.NotNil(t, err)
	assert.Equal(t, OutOfBoundsError, err)
	// unhappy low OOB
	obs, err = a.Get(-1)
	assert.Equal(t, 0, obs)
	assert.NotNil(t, err)
	assert.Equal(t, OutOfBoundsError, err)
}

// Testing RemoveAt, Removes a specified object at a specified location in the collection of stuff™, make sure it removes the right thing
// Returns an error if used at an invalid index (OOB), valid indexes are any index in which there is an existing item to remove
func TestRemoveAt(t *testing.T) {
	a := NewArrayList[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	a.Add(5)
	// happy paths/multiple times
	assert.Nil(t, a.RemoveAt(0))
	assert.Equal(t, []int{2, 3, 4, 5}, a.backer)
	assert.Nil(t, a.RemoveAt(1))
	assert.Equal(t, []int{2, 4, 5}, a.backer)
	assert.Nil(t, a.RemoveAt(2))
	assert.Equal(t, []int{2, 4}, a.backer)
	// unhappy low path
	err := a.RemoveAt(-1)
	assert.NotNil(t, err)
	assert.Equal(t, OutOfBoundsError, err)
	assert.Equal(t, []int{2, 4}, a.backer)
	// unhappy high path
	err = a.RemoveAt(2)
	assert.NotNil(t, err)
	assert.Equal(t, []int{2, 4}, a.backer)
	assert.Equal(t, OutOfBoundsError, err)
}

// Testing Set, Has the power to remake/reset the entire list of stuff to whatever we specify. Make sure it doesn't keep old stuff and new stuff gets added in its place
func TestSet(t *testing.T) {
	al := NewArrayList[int]()
	// Normal conditions
	assert.NotNil(t, al)
	al.Add(1)
	al.Add(2)
	al.Add(3)
	e := al.Set(0, 0)
	assert.Nil(t, e)
	assert.Equal(t, []int{0, 2, 3}, al.backer)
	e = al.Set(1, 0)
	assert.Nil(t, e)
	e = al.Set(2, 0)
	assert.Nil(t, e)
	assert.Equal(t, []int{0, 0, 0}, al.backer)
	//start checking boundary conditions
	e = al.Set(-1, 10)
	assert.Equal(t, []int{0, 0, 0}, al.backer)
	assert.NotNil(t, e)
	assert.Equal(t, OutOfBoundsError, e)
	e = al.Set(3, 420)
	assert.Equal(t, []int{0, 0, 0}, al.backer)
	assert.Equal(t, OutOfBoundsError, e)
	// potentially outlier conditions?
}

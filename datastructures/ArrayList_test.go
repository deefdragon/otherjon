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
}

// Testing Clear, Should clear our entire list of things, wiping it to 0/nil/empty
func TestClear(t *testing.T) {

}

// Testing IsEmpty, is it empty? TRUE, does it have stuff? FALSE, simple enough
func TestIsEmpty(t *testing.T) {

}

//Testing Size, a method that gets us the size of our list by how many valid things are inside, not including empty buckets, should return a whole number worth of things
func TestSize(t *testing.T) {

}

// List Methods

// Testing AddAt, Like Add which puts a thing in our collection of things, but specifically adds it at a specific point in our list of stuff. Make sure it adds to the right place.
func TestAddAt(t *testing.T) {

}

// Testing Get, Retrieves a specific/specified item from our list of things. If it gives us a thing we don't want, test fails
func TestGet(t *testing.T) {

}

// Testing RemoveAt, Removes a specified object at a specified location in the collection of stuff™, make sure it removes the right thing
func TestRemoveAt(t *testing.T) {

}

// Testing Set, Has the power to remake/reset the entire list of stuff to whatever we specify. Make sure it doesn't keep old stuff and new stuff gets added in its place
func TestSet(t *testing.T) {

}

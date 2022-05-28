package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAdd(t *testing.T) {
	//TODO Currently adds backwards, make it sane
	a := NewLinkedList[int]()
	arr := a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
	a.Add(1)
	arr = a.GetAsArray()
	assert.Equal(t, 1, a.size)
	assert.Equal(t, []int{1}, arr)
	a.Add(2)
	a.Add(3)
	arr = a.GetAsArray()
	assert.Equal(t, 3, a.size)
	assert.Equal(t, []int{3, 2, 1}, arr)
}

func TestLinkedListClear(t *testing.T) {
	a := NewLinkedList[int]()
	arr := a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
	a.Add(1)
	a.Add(2)
	a.Add(3)
	arr = a.GetAsArray()
	assert.Equal(t, 3, a.size)
	assert.Equal(t, []int{3, 2, 1}, arr)
	a.Clear()
	arr = a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
	a.Clear()
	a.Clear()
	a.Clear()
	arr = a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
}

func TestLinkedListIsEmpty(t *testing.T) {
	a := NewLinkedList[int]()
	assert.True(t, a.IsEmpty())
	a.Add(1)
	assert.False(t, a.IsEmpty())
}

func TestLinkedListSize(t *testing.T) {
	a := NewLinkedList[int]()
	assert.Equal(t, 0, a.size)
	a.Add(1)
	a.Add(2)
	a.Add(3)
	assert.Equal(t, 3, a.size)
}

func TestLinkedListGet(t *testing.T) {
	a := NewLinkedList[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	arr := a.GetAsArray()
	assert.Equal(t, 3, a.size)
	assert.Equal(t, []int{3, 2, 1}, arr)
	got, err := a.Get(1)
	assert.Equal(t, 2, got)
	assert.Nil(t, err)
	// TODO figure out why this is failing
	got, err = a.Get(3)
	assert.Equal(t, 0, got)
	assert.NotNil(t, err)

}

//TODO finish Set testing
func TestLinkedListSet(t *testing.T) {
	a := NewLinkedList[int]()
	arr := a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
	a.Add(1)
	a.Set(0, 5)
	arr = a.GetAsArray()
	assert.Equal(t, 1, a.size)
	assert.Equal(t, []int{5}, arr)
	a.Add(2)
	a.Add(3)
	arr = a.GetAsArray()
	assert.Equal(t, 3, a.size)
	assert.Equal(t, []int{3, 2, 5}, arr)
	a.Set(2, 3)
	arr = a.GetAsArray()
	assert.Equal(t, []int{3, 2, 3}, arr)
}

func TestLinkedListAddAt(t *testing.T) {
	a := NewLinkedList[int]()
	arr := a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
	a.AddAt(0, 1)
	a.AddAt(1, 2)
	a.AddAt(2, 3)
	arr = a.GetAsArray()
	assert.Equal(t, 3, a.size)
	assert.Equal(t, []int{1, 2, 3}, arr)

	a.AddAt(0, 4)
	arr = a.GetAsArray()
	assert.Equal(t, 4, a.size)
	assert.Equal(t, []int{4, 1, 2, 3}, arr)

	a.AddAt(1, 5)
	arr = a.GetAsArray()
	assert.Equal(t, 5, a.size)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, arr)

	a.AddAt(4, 6)
	arr = a.GetAsArray()
	assert.Equal(t, 6, a.size)
	assert.Equal(t, []int{4, 5, 1, 2, 6, 3}, arr)

}
func TestLinkedListRemoveAt(t *testing.T) {
	a := NewLinkedList[int]()
	// making sure its empty to start
	arr := a.GetAsArray()
	assert.Equal(t, 0, a.size)
	assert.Equal(t, []int{}, arr)
	//adding some items to our array
	a.Add(1)
	a.Add(2)
	a.Add(3)
	// making sure they add properly
	arr = a.GetAsArray()
	assert.Equal(t, 3, a.size)
	assert.Equal(t, []int{3, 2, 1}, arr)
	// testing removeAt at the beginning and verifying it works
	a.RemoveAt(0)
	arr = a.GetAsArray()
	assert.Equal(t, 2, a.size)
	assert.Equal(t, []int{2, 1}, arr)
	// Testing removeAt at the end of the list
	a.RemoveAt(1)
	arr = a.GetAsArray()
	assert.Equal(t, 1, a.size)
	assert.Equal(t, []int{2}, arr)
	// Removing from a negative OOB, assuring an error & nothing changes
	a.RemoveAt(-1)
	arr = a.GetAsArray()
	assert.Equal(t, 1, a.size)
	assert.Equal(t, []int{2}, arr)
	assert.NotNil(t, a.RemoveAt(-1))
	assert.Equal(t, OutOfBoundsError, a.RemoveAt(-1))
	// Removing from a positive OOB, assuring an error & nothing changes
	a.RemoveAt(1)
	arr = a.GetAsArray()
	assert.Equal(t, 1, a.size)
	assert.Equal(t, []int{2}, arr)
	assert.NotNil(t, a.RemoveAt(1))
	assert.Equal(t, OutOfBoundsError, a.RemoveAt(1))

}

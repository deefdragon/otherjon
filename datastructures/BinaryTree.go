package datastructures

var _ BinaryTree[int] = (*BinaryTreeImpl[int])(nil)

type BinaryTreeImpl[T comparable] struct {
	left  *BinaryTreeImpl[T]
	right *BinaryTreeImpl[T]
	data  T
}

func NewBinaryTree[T comparable]() *BinaryTreeImpl[T] {
	return &BinaryTreeImpl[T]{}
}

func (a *BinaryTreeImpl[T]) Add(T) {

}
func (a *BinaryTreeImpl[T]) Clear() {
	a.data = *new(T)
	a.left = nil
	a.right = nil
}
func (a *BinaryTreeImpl[T]) IsEmpty() bool {
	if a.left == nil && a.right == nil && a.data == *new(T) {
		return true
	}
	return false
}
func (a *BinaryTreeImpl[T]) Size() int {
	total := 1
	if a.left != nil {
		total += a.Size()
	}
	if a.right != nil {
		total += a.Size()
	}
	return total
}
func (a *BinaryTreeImpl[T]) GetLeftNode() BinaryTree[T] {
	return a.left
}
func (a *BinaryTreeImpl[T]) GetRightNode() BinaryTree[T] {
	return a.right
}
func (a *BinaryTreeImpl[T]) SetLeft(stuff T) {
	if a.left == nil {
		a.left = NewBinaryTree[T]()
	}

	a.left.data = stuff
}
func (a *BinaryTreeImpl[T]) SetRight(stuff T) {
	if a.right == nil {
		a.right = NewBinaryTree[T]()
	}
	a.right.data = stuff
}
func (a *BinaryTreeImpl[T]) RemoveLeft() {
	a.left = nil
}
func (a *BinaryTreeImpl[T]) RemoveRight() {
	a.right = nil
}
func (a *BinaryTreeImpl[T]) Get() T {
	return a.data
}

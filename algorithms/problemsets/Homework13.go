package problemsets

type Bookmark string

const (
	BookmarkA Bookmark = "A"
	BookmarkB Bookmark = "B"
)

type C[T any] interface {
	//build(x)
	PlaceMark(pos int, mark Bookmark)
	ReadPage(pos int)
	ShiftMark(mark Bookmark)
	MovePage(mark Bookmark)
}
type Cimpl struct {
}

// each page is a string, each mark is an int

// placeMark places a bookmark between current position and current position plus 1 in linear time
/*
As a static array or as a linked list, placeMark works in O(n) time, either because of the complications of seeking to the target position from a linked list,
or shifting every piece of data over in a static array. In a dynamic array, this remains O(n) also, because placing the bookmark requires a shift in the array and
the creation of new subarrays within the dynamic array (Due to their double sided nature), however, creating a dynamic array sets the data up for much faster running times
later on in other functions
*/
func (a *Cimpl) PlaceMark(pos int, mark Bookmark) {

}

// readPage returns the page at current position in constant time
/*
In a static array, ReadPage is O(1) constant time, because you would just input the index you want to return. In a linked list however, seeking to the target position
takes O(n) time. In a dynamic array we get the benefit of being able to return the page at a target index immediately in O(1) time, like that of a static array.
*/
func (a *Cimpl) ReadPage(pos int) {

}

// shiftMark takes the bookmark in front of the current position and moves the bookmark one space forward or one space back in constant time
/*
As a linked list, shiftMark is done in O(n) time because of needing to seek to the bookmark, then modify its pointer and other pointers. As a static array and a dynamic array,
shiftMark can be done in constant O(1) time, because the only action taking place is a swap with one piece of data to the left or right of the bookmark on a number line.
*/
func (a *Cimpl) ShiftMark(mark Bookmark) {

}

// movePage takes the page currently in front of a bookmark and moves it in front of the second bookmark
/*
As a static array, movePage is done in O(n) time because of the need to shift all data in between bookmark A&B to accomodate for the piece of moved data. As a linked list,
the action takes O(n) time as well, because of the need to seek to each position one at a time. In a dynamic array however, because the bookmarks act as effective start and end
points, we can add or remove data next to them in O(1) constant time, due to being able to identify their index immediately and also being able to modify data around them without
shifting any data unnecessarily
*/
func (a *Cimpl) MovePage(mark Bookmark) {

}

// why is 3 double sided dynamic arrays the correct solution for all of these & how they work & also runtime
/*
Overall, the correct solution for the whole of the list of functions is a dynamic array. While all mentioned data structures take the same time to place marks, the dynamic array
holds the benefits of constant time indexing from static arrays and constant time for adding data when working with the start of a list, end of a list, or around marked points.
The overall time complexity comparisons are as follows:
				 |PlaceMark; ReadPage; ShiftMark; MovePage|
				 |=========================================
	Linked list  |	 O(n)	|	O(n) |	  O(n)	|	 O(n) |
	Static array |	 O(n)	|	O(1) |	  O(1)	|	 O(n) |
	Dynamic array|	 O(n)	|	O(1) |	  O(1)	|	 O(1) |
				 |=========================================

	A dynamic array excels at every function except the initial placing of a mark, while the other data structures fall short for at least one other than placeMark.
*/

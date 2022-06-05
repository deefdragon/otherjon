package problemsets

// takes in data of type int array
func countLongSubarrays(data []int) int {
	//special cases for empty arrays or arrays of length one
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return 1
	}
	total := 0
	subTotal := 0
	maxValue := 0
	// cycles through every position in our array
	for i := 1; i < len(data); i++ {
		// checks if current position is greater than the previous
		if data[i] > data[i-1] {
			// increments our subtotal if so
			subTotal++
		} else {
			//otherwise  checks if our subtotal has reached a higher number than our current maxValue, if so sets our record subtotal to maxValue and resets total
			if subTotal > maxValue {
				maxValue = subTotal
				total = 0
			}
			// if our currently checked subtotal equals our maxValue, total increases, subtotal gets reset for a new subarray
			if subTotal == maxValue {
				total++
			}
			subTotal = 0
		}
	}
	// performs the check one more time outside of the loop to account for the last value in the array
	if subTotal > maxValue {
		maxValue = subTotal
		total = 0
	}
	if subTotal == maxValue {
		total++
	}
	// returns our total
	return total
}

// type d[T any] interface {
// 	insertFirst(T)
// 	deleteFirst() T
// 	insertLast(T)
// 	deleteLast() T
// }

// func swapEnds[T any](arr d[T]) {
// 	savedLast := arr.deleteLast()
// 	savedFirst := arr.deleteFirst()
// 	arr.insertFirst(savedLast)
// 	arr.insertLast(savedFirst)
// }
// func shiftLeft[T any](arr d[T], k int) {
// 	for i := 0; i < k; i++ {
// 		x := arr.deleteFirst()
// 		arr.insertLast(x)
// 	}
// }

// func shiftLeftR[T any](arr d[T], k int) {
// 	if k <= 0 {
// 		return
// 	}
// 	x := arr.deleteFirst()
// 	arr.insertLast(x)
// 	shiftLeftR(arr, k-1)
// }

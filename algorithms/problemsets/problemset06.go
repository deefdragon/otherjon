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

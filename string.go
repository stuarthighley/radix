package radix

// SortStrings sorts a slice of strings in ascending order.
func SortStrings[S ~string](input []S) {

	// Iterate backwards over string chars
	work1 := input
	work2 := make([]S, len(work1))
	for column := maxLen(input) - 1; column >= 0; column-- {

		// Accumulate each element's bucket
		clear(buckets[:])
		for _, e := range work1 {
			buckets[charOrZero(e, column)]++
		}

		// Convert bucket totals to culmulative totals
		for i := 1; i < numBuckets; i++ {
			buckets[i] += buckets[i-1]
		}

		// Use buckets to fill semi-sorted slice
		for i := len(work1) - 1; i >= 0; i-- {
			b := charOrZero(work1[i], column)
			buckets[b]--
			work2[buckets[b]] = work1[i]
		}

		// Swap slices over for next iteration
		work1, work2 = work2, work1
	}

	// One final copy to ensure sorted slice is returned.
	// Go will skip the copy if source and destination match.
	copy(input, work1)
}

// maxLen returns the length of the longest string in the given slice.
func maxLen[S ~string](slice []S) int {
	maxLen := 0
	for _, e := range slice {
		maxLen = max(maxLen, len(e))
	}
	return maxLen
}

// charOrZero returns the char at index i, or 0 if the string is shorter than i.
func charOrZero[S ~string](s S, i int) byte {
	if i >= len(s) {
		return 0
	}
	return s[i]
}

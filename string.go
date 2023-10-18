package radix

// SortStrings sorts a slice of strings in ascending order.
func SortStrings[T ~string](input []T) {

	// Iterate backwards over string chars
	work1 := input
	work2 := make([]T, len(work1))
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

	// One final copy if needed
	if !shallowEqual(input, work1) {
		copy(input, work1)
	}
}

// maxLen returns the length of the longest string in the given slice.
func maxLen[T ~string](s []T) int {
	maxLen := 0
	for _, e := range s {
		maxLen = max(maxLen, len(e))
	}
	return maxLen
}

// charOrZero returns the char at index i, or 0 if the string is shorter than i.
func charOrZero[T ~string](s T, i int) byte {
	if i >= len(s) {
		return 0
	}
	return s[i]
}

// shallowEqual returns true if the slice lengths and pointers are equal
func shallowEqual[T any](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	return len(s1) == 0 || &s1[0] == &s2[0]
}

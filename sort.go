package radix

import (
	"math/bits"
	"reflect"
	"slices"

	"golang.org/x/exp/constraints"
)

const base = 8
const numBuckets = 1 << base // 256
const mask = numBuckets - 1  // 255

var buckets [numBuckets]int

// SortInts sorts a slice of any integer type in ascending order.
func SortInts[T constraints.Integer](input []T) {
	if len(input) < 2 {
		return
	}

	// Prepare XOR mask for signed ints
	typeBitLen := reflect.TypeOf(*new(T)).Bits()
	signMask := T(0)
	testUnsigned := -1
	if T(testUnsigned) < 0 {
		signMask = 1 << (typeBitLen - 1)
	}

	// Get columns to iterate over
	maxBitLen := getMaxBitLen(input)

	// Iterate over all significant columns
	work1 := input
	work2 := make([]T, len(work1))
	for column := 0; column < typeBitLen; column += base {

		// We can skip iterations higher than the absolute values,
		// and don't contain the sign bit
		if column > maxBitLen && column < typeBitLen-base {
			continue
		}

		// Accumulate each element's bucket based on masked bits
		clear(buckets[:])
		for _, e := range work1 {
			b := ((int(e^signMask) >> column) & mask)
			buckets[b]++
		}

		// Convert bucket totals to culmulative totals
		for i := 1; i < numBuckets; i++ {
			buckets[i] += buckets[i-1]
		}

		// Use buckets to fill semi-sorted slice
		for i := len(work1) - 1; i >= 0; i-- {
			e := work1[i]
			b := ((int(e^signMask) >> column) & mask)
			buckets[b]--
			work2[buckets[b]] = e
		}

		// Swap slices over for next iteration
		work1, work2 = work2, work1
	}

	// One final copy if needed
	if !shallowEqual(input, work1) {
		copy(input, work1)
	}
}

// getMaxBitLen returns the highest number of bits used in the ints from the passed slice.
// The sign bit for negative ints is ignored.
func getMaxBitLen[T constraints.Integer](s []T) int {
	highest := slices.Max(s)
	if highest < 0 {
		highest = (^highest) + 1
	}
	lowest := slices.Min(s)
	if lowest < 0 {
		lowest = (^lowest) + 1
	}
	return max(bits.Len64(uint64(highest)), bits.Len64(uint64(lowest)))
}

// shallowEqual returns true if the slice lengths and pointers are equal
func shallowEqual[T any](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	return len(s1) == 0 || &s1[0] == &s2[0]
}

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

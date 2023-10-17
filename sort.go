package radix

import (
	"math/bits"
	"reflect"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
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
	typeBits := reflect.TypeOf(*new(T)).Bits()
	signMask := T(0)
	testUnsigned := -1
	if T(testUnsigned) < 0 {
		signMask = 1 << (typeBits - 1)
	}

	// Get columns to iterate over
	highestCol := getMaxBits(input)

	// Iterate over all significant columns
	work1 := input
	work2 := make([]T, len(work1))
	for column := 0; column < typeBits; column += base {

		// We can skip iterations higher than the absolute values,
		// and don't contain the sign bit
		if column > highestCol && column < typeBits-base {
			continue
		}

		// Accumulate the bucket for the masked bits for each element
		clear(buckets[:])
		for _, e := range work1 {
			b := ((int(e^signMask) >> column) & mask)
			buckets[b]++
		}

		// Convert buckets to culmulative totals
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

// getMaxBits returns the highest number of bits used in the ints from the passed slice.
// The sign bit for negative ints is ignored.
func getMaxBits[T constraints.Integer](s []T) int {
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
func SortStrings(input []string) {

	// Iterate backwards over string chars
	maxLen := maxLen(input)
	work1 := input
	work2 := make([]string, len(work1))
	for column := maxLen - 1; column >= 0; column-- {

		// Accumulate the bucket for each element
		clear(buckets[:])
		for _, e := range work1 {
			buckets[charOrZero(e, column)]++
		}

		// Convert buckets to culmulative totals
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
	if !shallowEqual(work1, work1) {
		copy(work1, work1)
	}
}

func maxLen(s []string) int {
	maxLen := 0
	for _, e := range s {
		maxLen = max(maxLen, len(e))
	}
	return maxLen
}

func charOrZero(s string, i int) byte {
	if i < len(s) {
		return s[i]
	}
	return 0
}

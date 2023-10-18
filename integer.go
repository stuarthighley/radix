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
func SortInts[I constraints.Integer](input []I) {
	if len(input) < 2 {
		return
	}

	// Prepare XOR mask for signed ints
	typeBitLen := reflect.TypeOf(*new(I)).Bits()
	signMask := I(0)
	testUnsigned := -1
	if I(testUnsigned) < 0 {
		signMask = 1 << (typeBitLen - 1)
	}

	// Get columns to iterate over
	maxBitLen := getMaxBitLen(input)

	// Iterate over all significant columns
	work1 := input
	work2 := make([]I, len(work1))
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
func getMaxBitLen[I constraints.Integer](s []I) int {
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

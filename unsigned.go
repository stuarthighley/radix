package radix

import (
	"golang.org/x/exp/constraints"
)

const base = 8 // 2^8 -> 256
const numBuckets = 1 << base
const mask = numBuckets - 1

var buckets [numBuckets]int

// Radix Sort
func Sort[T constraints.Unsigned](slice []T) []T {

	// Allocate copy of input slice
	sorted := make([]T, len(slice))

	// Loop until we reach the largest column
	largest := T(1)
	for column := 0; largest>>column > 0; column += base {

		// Accumulate the appropriate bucket for each element
		clear(buckets[:])
		for _, e := range slice {
			largest |= e
			buckets[(int(e)>>column)&mask]++
		}

		// Convert buckets to culmulative totals
		for i := 1; i < numBuckets; i++ {
			buckets[i] += buckets[i-1]
		}

		// Use buckets to fill 'semi-sorted' array
		for i := len(slice) - 1; i >= 0; i-- {
			e := slice[i]
			b := (int(e) >> column) & mask
			buckets[b]--
			sorted[buckets[b]] = e
		}

		// Swap slices over for next iteration
		slice, sorted = sorted, slice
	}

	return slice
}

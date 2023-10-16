package radix

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

const base = 8
const numBuckets = 1 << base // 256
const mask = numBuckets - 1  // 255

var buckets [numBuckets]int

// Radix Sort
func Sort[T constraints.Integer](input []T) []T {

	bits := reflect.TypeOf(*new(T)).Bits()
	testUnsigned := -1
	signMask := T(0)
	if T(testUnsigned) < 0 {
		signMask = 1 << (bits - 1)
	}

	// Allocate space for semi-output slice
	output := make([]T, len(input))

	// Loop until we reach the bitsUsed column
	bitsUsed := uint64(1)
	for column := 0; column < bits; column += base {

		if bitsUsed>>uint64(column)&mask == 0 {
			continue
		}

		// Accumulate the bucket for the masked bits for each element
		// Also record the largest used column
		clear(buckets[:])
		for _, e := range input {
			if e < 0 {
				bitsUsed |= (^uint64(e) ^ uint64(signMask))
			} else {
				bitsUsed |= uint64(e)
			}

			// bitsUnused |= ^uint64(e)
			b := ((int(e^signMask) >> column) & mask)
			buckets[b]++
		}

		// Convert buckets to culmulative totals
		for i := 1; i < numBuckets; i++ {
			buckets[i] += buckets[i-1]
		}

		// Use buckets to fill semi-sorted slice
		for i := len(input) - 1; i >= 0; i-- {
			e := input[i]
			b := ((int(e^signMask) >> column) & mask)
			buckets[b]--
			output[buckets[b]] = e
		}

		// Swap slices over for next iteration
		input, output = output, input
	}

	return input
}

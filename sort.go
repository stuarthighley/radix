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

// Radix Sort
func Sort[T constraints.Integer](input []T) []T {

	// Fast exit
	if len(input) == 0 {
		return input
	}

	// Prepare XOR mask for signed ints
	typeBits := reflect.TypeOf(*new(T)).Bits()
	signMask := T(0)
	testUnsigned := -1
	if T(testUnsigned) < 0 {
		signMask = 1 << (typeBits - 1)
	}

	// Calculate columns to iterate over
	highest := slices.Max(input)
	if highest < 0 {
		highest = ^highest
	}
	lowest := slices.Min(input)
	if lowest < 0 {
		lowest = (^lowest) + 1
	}
	highestCol := max(bits.Len64(uint64(highest)), bits.Len64(uint64(lowest)))

	// Iterate over all significant columns
	output := make([]T, len(input))
	for column := 0; column < typeBits; column += base {

		// We can skip iterations that are higher than the absolute values,
		// and don't contain the sign bit
		if column > highestCol && column < typeBits-base {
			continue
		}

		// Accumulate the bucket for the masked bits for each element
		// Also record the largest used column
		clear(buckets[:])
		for _, e := range input {
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

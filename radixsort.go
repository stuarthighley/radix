package radixsort

import (
	"golang.org/x/exp/constraints"
)

// Radix Sort
func RadixSort[T constraints.Integer](slice []T) []T {

	const base = 7 // 2^7 -> 128 | 2^8 -> 256 | 2^9 -> 512
	const numBuckets = 1 << base
	const mask = numBuckets - 1

	// typeBits := reflect.TypeOf(*new(T)).Bits()
	// base := T(typeBits / 8)
	// numBuckets := 1 << base
	// mask := T(numBuckets - 1)

	// The only memory allocation
	sorted := make([]T, len(slice))

	// Loop until we reach the largest column
	var largest T = 1
	buckets := make([]int, numBuckets)
	for column := 0; largest>>column > 0; column += base {

		// Accumulate the appropriate bucket for each element
		clear(buckets[:])
		for _, element := range slice {
			largest |= element
			buckets[(int(element)>>column)&mask]++
		}

		// Convert buckets to culmulative totals
		for bkt := 1; bkt < numBuckets; bkt++ {
			buckets[bkt] += buckets[bkt-1]
		}

		// Use buckets to fill 'semiSorted' array
		for i := len(slice) - 1; i >= 0; i-- {
			element := slice[i]
			bkt := (int(element) >> column) & mask
			buckets[bkt]--
			sorted[buckets[bkt]] = element
		}

		// Swap the slices over to init for next loop
		slice, sorted = sorted, slice

	}

	return slice
}

// BenchmarkRadixSortInt-8    	     966	   1377776 ns/op	  401413 B/op	       1 allocs/op
// BenchmarkRadixSortInt8-8   	   10000	    127329 ns/op	   57344 B/op	       1 allocs/op
// BenchmarkGoSortInt-8       	     324	   3597132 ns/op	       0 B/op	       0 allocs/op
// BenchmarkGoSortInt8-8      	     957	   1506262 ns/op	       0 B/op	       0 allocs/op

package radixsort

import (
	"golang.org/x/exp/constraints"
)

// Radix Sort
func RadixSort[T constraints.Unsigned](slice []T) []T {

	const base = 7 // 2^7 -> 128 | 2^8 -> 256
	const numBuckets = 1 << base
	const mask = numBuckets - 1

	// Allocate a copy of the slice, and the buckets
	sorted := make([]T, len(slice))
	buckets := make([]int, numBuckets)

	// Loop until we reach the largest column
	var largest T = 1
	for column := 0; largest>>column > 0; column += base {

		// Accumulate the appropriate bucket for each element
		clear(buckets)
		for _, e := range slice {
			largest |= e
			buckets[(int(e)>>column)&mask]++
		}

		// Convert buckets to culmulative totals
		for i := 1; i < numBuckets; i++ {
			buckets[i] += buckets[i-1]
		}

		// Use buckets to fill 'semiSorted' array
		for i := len(slice) - 1; i >= 0; i-- {
			e := slice[i]
			b := (int(e) >> column) & mask
			buckets[b]--
			sorted[buckets[b]] = e
		}

		// Swap the slices over to init for next loop
		slice, sorted = sorted, slice
	}

	// if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
	// 	bits := t.Bits()
	// 	signBit := 1 << (bits - 1)
	// 	for i := range slice {
	// 		slice[i] ^= T(signBit)
	// 	}
	// }

	return slice
}

// BenchmarkRadixSortInt-8    	     966	   1377776 ns/op	  401413 B/op	       1 allocs/op
// BenchmarkRadixSortInt8-8   	   10000	    127329 ns/op	   57344 B/op	       1 allocs/op
// BenchmarkGoSortInt-8       	     324	   3597132 ns/op	       0 B/op	       0 allocs/op
// BenchmarkGoSortInt8-8      	     957	   1506262 ns/op	       0 B/op	       0 allocs/op

package radixsort

// // Radix Sort Strings
// func RadixSortStrings(slice []string) []string {

// 	const base = 9         // 2^8 -> 256
// 	const numBuckets = 256 // One bucket for each byte value (treats unicode bytes separately, rightly or wrongly)
// 	// const mask = numBuckets - 1
// 	sorted := make([]string, len(slice), len(slice))

// 	// Find longest string TODO or could it be second longest string?
// 	longest := 0
// 	for _, v := range slice {
// 		if len(v) > longest {
// 			longest = len(v)
// 		}
// 	}

// 	// Loop until we reach the first char
// 	for c := longest - 1; c >= 0; c-- {

// 		// Accumulate the appropriate bucket for each element
// 		buckets := [numBuckets]int{0}
// 		for _, v := range slice {
// 			if len(v) > c {
// 				buckets[v[c]]++
// 			}
// 		}

// 		// Convert buckets to culmulative totals
// 		for i := 1; i < numBuckets; i++ {
// 			buckets[i] += buckets[i-1]
// 		}

// 		// Use the buckets to fill a 'semiSorted' array
// 		for i := len(slice) - 1; i >= 0; i-- {
// 			v := slice[i]
// 			var j byte
// 			if len(v) > i {
// 				j = v[c]
// 			} else {
// 				j = 0
// 			}
// 			buckets[j]--
// 			sorted[buckets[j]] = v
// 		}

// 		// Swap the slices over to init for next loop
// 		slice, sorted = sorted, slice

// 	}

// 	return slice
// }

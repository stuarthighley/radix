package radixsort

// // Find largest num & Accumulate the appropriate bucket for each element
// func accumulate(buckets []int, slice []int, column, mask int, largest *int) {
// 	for _, element := range slice {
// 		*largest |= element
// 		buckets[(element>>column)&mask]++
// 	}
// }

// func convertCulmulative(buckets []int) {
// 	for bkt := 1; bkt < len(buckets); bkt++ {
// 		buckets[bkt] += buckets[bkt-1]
// 	}
// }

// func fillSemiSorted(buckets []int, slice, sorted []int, column, mask int) {
// 	for i := len(slice) - 1; i >= 0; i-- {
// 		element := slice[i]
// 		bkt := (element >> column) & mask
// 		buckets[bkt]--
// 		sorted[buckets[bkt]] = element
// 	}
// }

// // Radix Sort
// func RadixSort(slice []int) []int {

// 	const base = 8 // 2^8 -> 256 | 2^9 -> 512
// 	const numBuckets = 1 << base
// 	const mask = numBuckets - 1
// 	sorted := make([]int, len(slice)) // The only memory allocation

// 	// Loop until we reach the largest column
// 	largest := 1
// 	for column := 0; largest>>column > 0; column += base {
// 		buckets := [numBuckets]int{0}
// 		accumulate(buckets[:], slice, column, mask, &largest)
// 		convertCulmulative(buckets[:])
// 		fillSemiSorted(buckets[:], slice, sorted, column, mask)
// 		slice, sorted = sorted, slice // Swap the slices over to init for next loop
// 	}

// 	return slice
// }

// Radix Sort
func RadixSort(slice []int) []int {

	const base = 8 // 2^8 -> 256 | 2^9 -> 512
	const numBuckets = 1 << base
	const mask = numBuckets - 1
	sorted := make([]int, len(slice)) // The only memory allocation

	// // Find largest num
	// largest := 0
	// for _, element := range slice {
	// 	if element > largest {
	// 		largest = element
	// 	}
	// }

	// Loop until we reach the largest column
	largest := 1
	for column := 0; largest>>column > 0; column += base {

		// Accumulate the appropriate bucket for each element
		buckets := [numBuckets]int{0}
		for _, element := range slice {
			largest |= element
			buckets[(element>>column)&mask]++
		}

		// Convert buckets to culmulative totals
		for bkt := 1; bkt < numBuckets; bkt++ {
			buckets[bkt] += buckets[bkt-1]
		}

		// Use buckets to fill 'semiSorted' array
		for i := len(slice) - 1; i >= 0; i-- {
			element := slice[i]
			bkt := (element >> column) & mask
			buckets[bkt]--
			sorted[buckets[bkt]] = element
		}

		// Swap the slices over to init for next loop
		slice, sorted = sorted, slice

	}

	return slice
}

// func RadixSort(slice []int) {

// 	// Algorithm from CLRS, chapter 8

// 	const m = 256
// 	lenSlice := len(slice)

// 	counts := make([]int, m)

// 	src := slice
// 	dst := make([]int, lenSlice)

// 	for shift := uint(0); shift < 8*8; shift += 8 {
// 		// shift := i * 8

// 		for i := range counts {
// 			counts[i] = 0
// 		}

// 		for i := 0; i < lenSlice; i++ {
// 			j := byte(src[i] >> shift)
// 			counts[j]++
// 		}

// 		for j := 1; j < m; j++ {
// 			counts[j] += counts[j-1]
// 		}

// 		for i := lenSlice - 1; i >= 0; i-- {
// 			j := byte(src[i] >> shift)
// 			dst[counts[j]-1] = src[i]
// 			counts[j]--
// 		}

// 		src, dst = dst, src
// 	}
// }

package radix

import (
	"math"
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// SortFloats sorts a slice of any float type in ascending order.
func SortFloats[F constraints.Float](input []F) {
	input = sortNaNs(input)
	if isFloat32[F]() {
		unsafeFlipSortFlip[F, uint32](input)
	} else {
		unsafeFlipSortFlip[F, uint64](input)
	}
}

// sortNaNs put NaNs up front, similar to sort.Float64s, returning a slice of x excluding those nans
func sortNaNs[F constraints.Float](slice []F) []F {
	nans := 0
	for i, e := range slice {
		if math.IsNaN(float64(e)) {
			slice[i] = slice[nans]
			slice[nans] = e
			nans++
		}
	}
	return slice[nans:]
}

// isFloat32 returns true if the passed type is float32, else false
func isFloat32[F constraints.Float]() bool {
	return F(math.SmallestNonzeroFloat32)/2 == 0
}

// unsafeFlipSortFlip converts float slices to unsigned, flips some bits to allow sorting, sorts and unflips.
// F and U must be the same bit size, and len(buf) must be >= len(x)
// This will not work if NaNs are present in x. Remove them first.
func unsafeFlipSortFlip[F constraints.Float, U constraints.Unsigned](slice []F) {
	floatSize := reflect.TypeOf(*new(F)).Size() * 8
	// Change slice type to uint
	p := (*U)(unsafe.Pointer(unsafe.SliceData(slice)))
	uintSlice := unsafe.Slice(p, cap(slice))[:len(slice)]
	// Flip some bits to make the sort work
	floatFlip(uintSlice, 1<<(floatSize-1))
	// Sort the slice as a uint slice
	SortInts(uintSlice)
	// Flip the bits back
	floatUnflip(uintSlice, 1<<(floatSize-1))
}

// floatFlip flips some bits to make the sort work
func floatFlip[U constraints.Unsigned](y []U, topBit U) {
	for i, e := range y {
		if e&topBit == topBit {
			y[i] = e ^ (^U(0))
		} else {
			y[i] = e ^ topBit
		}
	}
}

// floatUnflip undoes floatFlip
func floatUnflip[U constraints.Unsigned](y []U, topBit U) {
	for i, e := range y {
		if e&topBit == topBit {
			y[i] = e ^ topBit
		} else {
			y[i] = e ^ (^U(0))
		}
	}
}

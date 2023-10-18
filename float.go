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
		flipSortUnflip[F, uint32](input)
	} else {
		flipSortUnflip[F, uint64](input)
	}
}

// sortNaNs put NaNs up front, similar to sort.Float64s, returning a slice excluding the NaNs.
func sortNaNs[F constraints.Float](slice []F) []F {
	nans := 0
	for i := range slice {
		if math.IsNaN(float64(slice[i])) {
			slice[i], slice[nans] = slice[nans], slice[i]
			nans++
		}
	}
	return slice[nans:]
}

// isFloat32 returns true if the passed type is float32, else false
func isFloat32[F constraints.Float]() bool {
	return F(math.SmallestNonzeroFloat32)/2 == 0
}

// flipSortUnflip converts float slices to unsigned, flips some bits to allow sorting, sorts and unflips.
// F and U must be the same bit size, and len(buf) must be >= len(x).
// Will not work if NaNs are present in x. Remove them first.
func flipSortUnflip[F constraints.Float, U constraints.Unsigned](slice []F) {
	// Change slice type to uint
	p := (*U)(unsafe.Pointer(unsafe.SliceData(slice)))
	uintSlice := unsafe.Slice(p, cap(slice))[:len(slice)]
	// Flip some bits for the sort to work
	floatFlip(uintSlice)
	// Sort the slice as a uint slice
	SortInts(uintSlice)
	// Flip the bits back
	floatUnflip(uintSlice)
}

// floatFlip flips some bits to make the sort work
// If top bit set, flip every bit. Else, turn on top bit
func floatFlip[U constraints.Unsigned](slice []U) {
	topBit := U(1 << ((reflect.TypeOf(*new(U)).Bits()) - 1))
	for i, e := range slice {
		if e&topBit == topBit {
			slice[i] = ^slice[i]
		} else {
			slice[i] |= topBit
		}
	}
}

// floatUnflip undoes floatFlip
// If top bit set, turn off top bit. Else, flip every bit
func floatUnflip[U constraints.Unsigned](slice []U) {
	topBit := U(1 << ((reflect.TypeOf(*new(U)).Bits()) - 1))
	for i, e := range slice {
		if e&topBit == topBit {
			slice[i] &= ^topBit
		} else {
			slice[i] = ^slice[i]
		}
	}
}

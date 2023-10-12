package radixsort

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

const testLen = 10
const benchmarkLen = 50000

func TestRadixSortUint(t *testing.T) {
	radix := make([]uint, testLen)
	for i := range radix {
		radix[i] = uint(rand.Uint64())
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

// func TestRadixSortIntNeg(t *testing.T) {
// 	radix := make([]int, testLen)
// 	for i := range radix {
// 		radix[i] = rand.Int() - (math.MaxInt >> 1)
// 	}
// 	goSort := slices.Clone(radix)

// 	radix = RadixSort(radix)
// 	slices.Sort(goSort)
// 	assert.Equal(t, goSort, radix)
// }

func TestRadixSortUint8(t *testing.T) {
	radix := make([]uint8, testLen)
	for i := range radix {
		radix[i] = uint8(rand.Intn(math.MaxUint8))
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

// func TestRadixSortInt8Neg(t *testing.T) {
// 	radix := make([]int8, testLen)
// 	for i := range radix {
// 		radix[i] = int8(rand.Intn(math.MaxInt8) - (math.MaxInt8 >> 1))
// 	}
// 	goSort := slices.Clone(radix)

// 	radix = RadixSort(radix)
// 	slices.Sort(goSort)
// 	assert.Equal(t, goSort, radix)
// }

func TestRadixSortUint16(t *testing.T) {
	radix := make([]uint16, testLen)
	for i := range radix {
		radix[i] = uint16(rand.Intn(math.MaxUint16))
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

func TestRadixSortUint32(t *testing.T) {
	radix := make([]uint32, testLen)
	for i := range radix {
		radix[i] = uint32(rand.Uint32())
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

func TestRadixSortUint64(t *testing.T) {
	radix := make([]uint64, testLen)
	for i := range radix {
		radix[i] = uint64(rand.Uint64())
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

func BenchmarkRadixSortUintFullRange(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Uint64())
		}

		b.StartTimer()
		RadixSort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortUintLimitedRange(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Intn(math.MaxInt16))
		}

		b.StartTimer()
		RadixSort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortUint8(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint8, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint8(rand.Intn(math.MaxUint8))
		}

		b.StartTimer()
		RadixSort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkGoSortUintFullRange(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Uint64())
		}

		b.StartTimer()
		slices.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkGoSortSortUintLimitedRange(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Intn(math.MaxInt16))
		}

		b.StartTimer()
		slices.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkGoSortUint8(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint8, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint8(rand.Intn(math.MaxUint8))
		}

		b.StartTimer()
		slices.Sort(unsortedList)
		b.StopTimer()
	}
}

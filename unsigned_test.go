package radix_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuarthighley/radix"
	"golang.org/x/exp/slices"
)

const testLen = 10
const benchmarkLen = 50000

func TestRadixSortUint(t *testing.T) {
	radixSlice := make([]uint, testLen)
	for i := range radixSlice {
		radixSlice[i] = uint(rand.Uint64())
	}
	goSort := slices.Clone(radixSlice)

	radixSlice = radix.Sort(radixSlice)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radixSlice)
}

func TestRadixSortUint8(t *testing.T) {
	radixSlice := make([]uint8, testLen)
	for i := range radixSlice {
		radixSlice[i] = uint8(rand.Intn(math.MaxUint8))
	}
	goSort := slices.Clone(radixSlice)

	radixSlice = radix.Sort(radixSlice)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radixSlice)
}

func TestRadixSortUint16(t *testing.T) {
	radixSlice := make([]uint16, testLen)
	for i := range radixSlice {
		radixSlice[i] = uint16(rand.Intn(math.MaxUint16))
	}
	goSort := slices.Clone(radixSlice)

	radixSlice = radix.Sort(radixSlice)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radixSlice)
}

func TestRadixSortUint32(t *testing.T) {
	radixSlice := make([]uint32, testLen)
	for i := range radixSlice {
		radixSlice[i] = uint32(rand.Uint32())
	}
	goSort := slices.Clone(radixSlice)

	radixSlice = radix.Sort(radixSlice)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radixSlice)
}

func TestRadixSortUint64(t *testing.T) {
	radixSlice := make([]uint64, testLen)
	for i := range radixSlice {
		radixSlice[i] = uint64(rand.Uint64())
	}
	goSort := slices.Clone(radixSlice)

	radixSlice = radix.Sort(radixSlice)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radixSlice)
}

func BenchmarkRadixSortUintFullRange(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]uint, benchmarkLen)
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Uint64())
		}

		b.StartTimer()
		radix.Sort(unsortedList)
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
		radix.Sort(unsortedList)
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
		radix.Sort(unsortedList)
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

package radixsort

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

const testSliceSize = 50000

func TestRadixSortInt(t *testing.T) {
	radix := make([]int, testSliceSize)
	for i := range radix {
		radix[i] = rand.Intn(math.MaxInt)
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

func TestRadixSortInt8(t *testing.T) {
	radix := make([]int8, testSliceSize)
	for i := range radix {
		radix[i] = int8(rand.Intn(math.MaxInt8))
	}
	goSort := slices.Clone(radix)

	radix = RadixSort(radix)
	slices.Sort(goSort)
	assert.Equal(t, goSort, radix)
}

func BenchmarkRadixSortInt(b *testing.B) {

	b.StopTimer()

	for n := 0; n < b.N; n++ {

		unsortedList := make([]int, testSliceSize)
		for i := range unsortedList {
			unsortedList[i] = rand.Intn(math.MaxInt)
		}

		b.StartTimer()
		unsortedList = RadixSort(unsortedList)
		b.StopTimer()

	}

}

func BenchmarkRadixSortInt8(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]int8, testSliceSize)
		for i := range unsortedList {
			unsortedList[i] = int8(rand.Intn(math.MaxInt8))
		}

		b.StartTimer()
		unsortedList = RadixSort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkGoSortInt(b *testing.B) {

	var unsortedList []int
	b.StopTimer()

	for n := 0; n < b.N; n++ {

		unsortedList = make([]int, testSliceSize)
		for i := range unsortedList {
			unsortedList[i] = rand.Intn(math.MaxInt)
		}

		b.StartTimer()
		slices.Sort(unsortedList)
		b.StopTimer()
	}

}

func BenchmarkGoSortInt8(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		unsortedList := make([]int8, testSliceSize)
		for i := range unsortedList {
			unsortedList[i] = int8(rand.Intn(math.MaxInt8))
		}

		b.StartTimer()
		slices.Sort(unsortedList)
		b.StopTimer()
	}
}

package radix_test

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuarthighley/radix"
)

var float32Cases = [][]float32{
	{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586},
	{-74, -59, -238, -784, -9845, -959, -905, -42, -7586, -5467984, -7586},
	{74.3, 59.0, float32(math.Inf(1)), 238.2, -784.0, 2.3, float32(math.Inf(-1)), 9845.768, -959.7485, 905, 7.8, 7.8, 74.3, 59.0, float32(math.Inf(1)), 238.2, -784.0, 2.3},
	{74.3, 59.0, float32(math.Inf(1)), 238.2, -784.0, 2.3, float32(math.NaN()), float32(math.NaN()), float32(math.Inf(-1)), 9845.768, -959.7485, 905, 7.8, 7.8},
}

var float64Cases = [][]float64{
	{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586},
	{-74, -59, -238, -784, -9845, -959, -905, -42, -7586, -5467984, -7586},
	{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8, 74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3},
	{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8},
}

// Tests

func TestSortFloat32s(t *testing.T) {
	for _, test := range float32Cases {
		radix.SortFloats(test)
		assert.True(t, slices.IsSorted(test), "not sorted: %v", test)
	}
}

func TestSortFloat64s(t *testing.T) {
	for _, test := range float64Cases {
		radix.SortFloats(test)
		assert.True(t, slices.IsSorted(test), "not sorted: %v", test)
	}
}

// Benchmarks

func BenchmarkRadixSortFloat32s(b *testing.B) {
	b.StopTimer()
	unsorted := make([]float32, benchmarkLen)
	for i := range unsorted {
		unsorted[i] = rand.Float32()
	}
	slice := make([]float32, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(slice, unsorted)
		b.StartTimer()
		radix.SortFloats(slice)
		b.StopTimer()
	}
}

func BenchmarkSlicesSortFloat32s(b *testing.B) {
	b.StopTimer()
	unsorted := make([]float32, benchmarkLen)
	for i := range unsorted {
		unsorted[i] = rand.Float32()
	}
	slice := make([]float32, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(slice, unsorted)
		b.StartTimer()
		slices.Sort(slice)
		b.StopTimer()
	}
}

func BenchmarkRadixSortFloat64s(b *testing.B) {
	b.StopTimer()
	unsorted := make([]float64, benchmarkLen)
	for i := range unsorted {
		unsorted[i] = rand.Float64()
	}
	slice := make([]float64, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(slice, unsorted)
		b.StartTimer()
		radix.SortFloats(slice)
		b.StopTimer()
	}
}

func BenchmarkSlicesSortFloat64s(b *testing.B) {
	b.StopTimer()
	unsorted := make([]float64, benchmarkLen)
	for i := range unsorted {
		unsorted[i] = rand.Float64()
	}
	slice := make([]float64, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(slice, unsorted)
		b.StartTimer()
		slices.Sort(slice)
		b.StopTimer()
	}
}

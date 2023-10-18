package radix_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuarthighley/radix"
)

// Tests

var stringCases = [][]string{
	{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"},
	{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&z", "***"},
}

func TestSortStrings(t *testing.T) {
	for _, test := range stringCases {
		radix.SortStrings(test)
		assert.True(t, slices.IsSorted(test), "not sorted: %v", test)
	}
}

// Benchmarks

func BenchmarkRadixSortStrings(b *testing.B) {
	b.StopTimer()
	unsorted := make([]string, benchmarkLen)
	for i := range unsorted {
		unsorted[i] = strconv.Itoa(i ^ 0x2cc)
	}
	data := make([]string, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		radix.SortStrings(data)
		b.StopTimer()
	}
}

func BenchmarkSlicesSortStrings(b *testing.B) {
	b.StopTimer()
	unsorted := make([]string, benchmarkLen)
	for i := range unsorted {
		unsorted[i] = strconv.Itoa(i ^ 0x2cc)
	}
	data := make([]string, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		slices.Sort(data)
		b.StopTimer()
	}
}

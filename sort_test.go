package radix_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuarthighley/radix"
	"golang.org/x/exp/slices"
)

const fuzzLen = 100
const benchmarkLen = 100000

var testLens = []int{0, 1, 2, 10, 1000, 100000, 1000000}
var fuzzCases = [][]byte{{}, {1}, {2, 1}, {255, 1, 128}}

func FuzzSortByte(f *testing.F) {
	for _, tc := range fuzzCases {
		f.Add(tc)
	}
	tc := make([]byte, fuzzLen)
	for i := range tc {
		tc[i] = byte(rand.Intn(math.MaxUint8))
	}
	f.Add(tc)

	f.Fuzz(func(t *testing.T, actual []byte) {
		expected := slices.Clone(actual)
		slices.Sort(expected)
		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	})
}

func TestRadixSortUint(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]uint, testLen)
		for i := range actual {
			actual[i] = uint(rand.Uint64())
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	}
}

func TestRadixSortUintLimitedRange(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]uint, testLen)
		for i := range actual {
			actual[i] = uint(rand.Intn(math.MaxUint16))
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	}
}

func TestRadixSortInt(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]int, testLen)
		for i := range actual {
			actual[i] = int(rand.Uint64() - math.MaxInt64)
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
		if t.Failed() {
			break
		}
	}
}

func TestRadixSortIntLimitedRange(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]int, testLen)
		for i := range actual {
			actual[i] = int(rand.Intn(math.MaxUint16) - math.MaxInt16)
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
		if t.Failed() {
			break
		}
	}
}

func TestRadixSortUint8(t *testing.T) {
	for _, testLen := range testLens {

		actual := make([]uint8, testLen)
		for i := range actual {
			actual[i] = uint8(rand.Intn(math.MaxUint8))
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	}
}

func TestRadixSortInt8(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]int8, testLen)
		for i := range actual {
			actual[i] = int8(rand.Intn(math.MaxUint8) - 128)
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
		if t.Failed() {
			break
		}
	}
}

func TestRadixSortUint16(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]uint16, testLen)
		for i := range actual {
			actual[i] = uint16(rand.Intn(math.MaxUint16))
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	}
}

func TestRadixSortInt16(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]int16, testLen)
		for i := range actual {
			actual[i] = int16(rand.Intn(math.MaxUint16) - math.MaxInt16)
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
		if t.Failed() {
			break
		}
	}
}

func TestRadixSortUint32(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]uint32, testLen)
		for i := range actual {
			actual[i] = uint32(rand.Uint32())
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	}
}

func TestRadixSortUint64(t *testing.T) {
	for _, testLen := range testLens {
		actual := make([]uint64, testLen)
		for i := range actual {
			actual[i] = uint64(rand.Uint64())
		}
		expected := slices.Clone(actual)
		slices.Sort(expected)

		actual = radix.Sort(actual)
		assert.Equal(t, expected, actual)
	}
}

func BenchmarkRadixSortUintFullRange(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]uint, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Uint64())
		}

		b.StartTimer()
		radix.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortIntFullRange(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]int, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = int(rand.Uint64() - math.MaxInt)
		}

		b.StartTimer()
		radix.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortUintLimitedRange(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]uint, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = uint(rand.Intn(math.MaxUint16))
		}

		b.StartTimer()
		radix.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortIntLimitedRange(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]int, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = int(rand.Intn(math.MaxUint16) - math.MaxInt16)
		}

		b.StartTimer()
		radix.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortUint8(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]uint8, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = uint8(rand.Intn(math.MaxUint8))
		}

		b.StartTimer()
		radix.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkRadixSortInt8(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]int8, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = int8(rand.Intn(math.MaxUint8) - math.MaxInt8)
		}

		b.StartTimer()
		radix.Sort(unsortedList)
		b.StopTimer()
	}
}

func BenchmarkGoSortUintFullRange(b *testing.B) {
	b.StopTimer()
	unsortedList := make([]uint, benchmarkLen)
	for n := 0; n < b.N; n++ {
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
	unsortedList := make([]uint, benchmarkLen)
	for n := 0; n < b.N; n++ {
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
	unsortedList := make([]uint8, benchmarkLen)
	for n := 0; n < b.N; n++ {
		for i := range unsortedList {
			unsortedList[i] = uint8(rand.Intn(math.MaxUint8))
		}

		b.StartTimer()
		slices.Sort(unsortedList)
		b.StopTimer()
	}
}

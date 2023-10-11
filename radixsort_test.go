package radixsort

import (
	"math/rand"
	"testing"
	"time"
)

const testSliceSize = 50000
const testSliceIntRange = 10000000

func BenchmarkRadixSort(b *testing.B) {

	b.StopTimer()

	//Provide seed
	rand.Seed(time.Now().Unix())

	for n := 0; n < b.N; n++ {

		unsortedList := make([]int, testSliceSize)
		for i := range unsortedList {
			unsortedList[i] = rand.Intn(testSliceIntRange)
		}

		b.StartTimer()
		RadixSort(unsortedList)
		b.StopTimer()

	}

}

func BenchmarkQuicksort(b *testing.B) {

	var unsortedList []int
	b.StopTimer()

	//Provide seed
	rand.Seed(time.Now().Unix())

	for n := 0; n < b.N; n++ {

		unsortedList = make([]int, testSliceSize)
		for i := range unsortedList {
			unsortedList[i] = rand.Intn(testSliceIntRange)
		}

		b.StartTimer()
		QuickSort(unsortedList)
		b.StopTimer()
	}

}

// func TestFirst(t *testing.T) {

// }

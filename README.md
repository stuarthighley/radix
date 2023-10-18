# radix

The `radix` package implements radix sorting, currently for all signed and unsigned integer slices, and string slices.

Usage:

```go
package main

import (
	"fmt"
	"math"

	"github.com/stuarthighley/radix"
)

func main() {
	ints := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	radix.SortInts(ints)
	fmt.Println(ints)

	data := []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
	radix.SortStrings(data)
	fmt.Println(data)

	floats := []float64{74.3, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.Inf(-1), 9845.768, -959.7485, 7.8, 7.8}
	radix.SortFloats(floats)
	fmt.Println(floats)
}
```

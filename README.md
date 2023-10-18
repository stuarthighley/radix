# radix

The `radix` package implements radix sorting, currently for all signed and unsigned integer slices, and string slices.

Usage:

```go
package main

import (
	"fmt"

	"github.com/stuarthighley/radix"
)

func main() {
	ints := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	radix.SortInts(ints)
	fmt.Println(ints)

	data := []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
	radix.SortStrings(data)
	fmt.Println(data)
}
```

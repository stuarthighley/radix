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
	uints := []uint64{4, 2, 1, 3}
	radix.SortInts(uints)
	fmt.Println(uints)

	data := []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
	radix.SortStrings(data)
	fmt.Println(data)
}
```

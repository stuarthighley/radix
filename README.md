# radix

The `radix` package implements radix sorting, currently for all signed and unsigned integer types.

Usage:

```go
import "github.com/stuarthighley/radix"

func Example() {
	uints := []uint64{4, 2, 1, 3}
	radix.SortInts(uints)

	data := []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
	radix.SortStrings(data)
}
```

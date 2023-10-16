# radix

The `radix` package implements radix sorting, currently for all signed and unsigned integer types.

Usage:

```go
import "github.com/stuarthighley/radix"

func Example() {
	data := []uint64{4, 2, 1, 3}
	data = radix.Sort(data)
}
```

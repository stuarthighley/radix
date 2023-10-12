# radix

`radix` package implements radix sorting, currently for all unsigned integer types.

The usage looks like:

```go
import "github.com/stuarthighley/radix"

func Example() {
	data := []uint64{4, 2, 1, 3}
	data = radix.Sort(data)
}
```
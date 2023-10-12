# radixsort

`radixsort` package implements radix sorting, currently for all unsigned integer types.

The usage looks like:

```go
import "github.com/stuarthighley/radixsort"

func Example() {
	data := []int64{3,4,1,2}

	radixsort.Uint64(data, tmpbuf)
}
```
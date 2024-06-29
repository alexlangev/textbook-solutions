# Exercise 3

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("value of b: ", b)
	fmt.Println("value of smallI: ", smallI)
	fmt.Println("value of bigI: ", bigI)
}
```

## Thoughts

I am surprised that arithmetic overflows are a thing in such a young language.

The math package has constants for the max and min values of numerical types.

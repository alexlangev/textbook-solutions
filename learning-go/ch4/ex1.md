# Exercise 1

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randomInts []int

	for i := 0; i < 100; i++ {
		var x int = rand.Intn(100)
		randomInts = append(randomInts, x)
	}

	fmt.Println(randomInts)
}
```

## Thoughts

Assigning `append` to the slice variable is kinda weird. Otherwise its quite straightforward.

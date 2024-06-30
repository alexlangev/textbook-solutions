# Exercise 3

```go
package main

import (
	"fmt"
)

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}

	fmt.Println(total)
}
```

The bug comes from the `total` variable created inside the for-loop which shadows the `total` created in the main function.

## Thoughts

Makes sense!

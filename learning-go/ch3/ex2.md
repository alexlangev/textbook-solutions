# Exercise 2

```go
package main

import (
	"fmt"
)

func main() {
	var message string = "Hi 👩 and 👨"

	face := message[3:7]
	fmt.Println(face)
}
```

## Thoughts

I still don't fully grasp the connection between `strings`, `runes`, `bytes` and `slices`...

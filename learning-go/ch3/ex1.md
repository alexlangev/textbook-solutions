# Exercise 1

```go
package main

import (
	"fmt"
)

func main() {
	var greetings []string = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	s1 := greetings[:2]
	s2 := greetings[1:4]
	s3 := greetings[3:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
```

## Thoughts

It feels like lists in Python.

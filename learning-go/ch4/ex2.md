# Exercise 2

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

	for _, num := range randomInts {
		switch {
		case num%3 == 0 && num%2 == 0:
			fmt.Println("Six!")
		case num%2 == 0:
			fmt.Println("Two!")
		case num%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}

	}
}
```

## Thoughts

I Love the range loops. The switch syntax is a bit weird for now...

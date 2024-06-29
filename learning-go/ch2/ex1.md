# Exercise 1

```go
package main

import (
	"fmt"
)

func main() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println("The value of i is: ", i)
	fmt.Println("The value of f is: ", f)
}
```

## Thoughts

I never know when I should use `var` or `:=`. I guess this will come with experience.

Is `f` pointing to the same value in memory as `i` is?

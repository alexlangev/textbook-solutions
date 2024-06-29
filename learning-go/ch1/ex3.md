# Exercice 3

## My solution

I decided to add a couple of **newlines** using `\n`. I also switched the first `Println` call for a `Printf` with `"world"` as parameter. I then added a `Println` with a `\t` which indents the output by a tab space.

```go
// hello.go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, %s!\n\n", "world")

	fmt.Println("\tWhat will happen?")
}
```

Produces the following output:

```
Hello, world!

	What will happen?

```

## Thoughts

Kinda reminds me of Java. I'm not sure I like the `Printf` way of printing values. It's kinda verbose...

# Exercise 3

```go
package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	Tommy := Employee{
		"Tommy",
		"Johns",
		69,
	}

	Ashley := Employee{
		lastName:  "Madison",
		firstName: "Ashley",
		id:        420,
	}

	var Owen Employee
	Owen.firstName = "Owen"
	Owen.lastName = "Smith"
	Owen.id = 1234

	fmt.Println(Tommy)
	fmt.Println(Ashley)
	fmt.Println(Owen)
}
```

## Thoughts

Simple enough... Their like JS objects but better.

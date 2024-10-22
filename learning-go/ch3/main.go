package main

import (
	"fmt"
)

func main() {
	var message string = "Hi 👩 and 👨"
	var r []rune = []rune(message)
	fmt.Println(string(r[3]))
}

// main package in a go module contains the code that starts a Go program
package main

import "fmt"

func main() {
	// let's mess up the formating of the file.

	fmt.Printf("Hello, \n\n\n\t%s!\n", "world")
	fmt.Println("Is this working?")
}

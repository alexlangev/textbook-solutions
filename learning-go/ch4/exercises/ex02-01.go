package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 0, 100)

	for range 100 {
		s = append(s, rand.Intn(100))
	}

	for _, val := range s {
		switch {
		case val%2 == 0 && val%3 == 0:
			fmt.Println("Six!")
		case val%2 == 0:
			fmt.Println("Two!")
		case val%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind!")
		}
	}
}

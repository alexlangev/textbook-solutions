package main
import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 0, 100)

	for range 100{
		s = append(s, rand.Intn(100))
	}
	
	fmt.Println(s)
}

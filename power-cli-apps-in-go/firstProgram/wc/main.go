package main

import (
	"bufio" // to read text
	"fmt"   // to print the output
	"flag"
	"io"    // provides the io.Reader interface
	"os"    // can use os stuff
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
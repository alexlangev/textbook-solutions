package main

import (
	"bufio" // to read text
	"flag"  // create and manage command-line flags
	"fmt"   // to print formatted output
	"io"    // provides the io.Reader interface
	"os"    // to use os resources
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

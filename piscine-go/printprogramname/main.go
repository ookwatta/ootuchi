package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	start := 0
	for i, r := range path {
		if r == '/' {
			start = i + 1
		}
	}
	for _, r := range path[start:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

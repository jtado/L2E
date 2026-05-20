//go:build ignore

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	printed := make(map[rune]bool)

	for _, r := range s1 {
		if !printed[r] {
			z01.PrintRune(r)
			printed[r] = true
		}
	}
	for _, r := range s2 {
		if !printed[r] {
			z01.PrintRune(r)
			printed[r] = true
		}
	}
	z01.PrintRune('\n')
}

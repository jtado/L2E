//go:build ignore

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	m2 := make(map[rune]bool)
	for _, r := range s2 {
		m2[r] = true
	}

	printed := make(map[rune]bool)

	for _, r := range s1 {
		if m2[r] && !printed[r] {
			z01.PrintRune(r)
			printed[r] = true
		}
	}
	z01.PrintRune('\n')
}

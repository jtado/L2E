//go:build ignore

package main

import (
	"os"
	"strings"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	if s == "" {
		z01.PrintRune('\n')
		return
	}
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		for _, r := range words[i] {
			z01.PrintRune(r)
		}
		if i > 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

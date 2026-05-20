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

	if s1 == "" {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	i1 := 0
	for i2 := 0; i2 < len(r2); i2++ {
		if r1[i1] == r2[i2] {
			i1++
			if i1 == len(r1) {
				z01.PrintRune('1')
				z01.PrintRune('\n')
				return
			}
		}
	}
	z01.PrintRune('0')
	z01.PrintRune('\n')
}

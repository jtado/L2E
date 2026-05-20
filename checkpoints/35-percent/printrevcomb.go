//go:build ignore

package main

import "github.com/01-edu/z01"

func main() {
	first := true
	for i := '9'; i >= '0'; i-- {
		for j := i - 1; j >= '0'; j-- {
			for k := j - 1; k >= '0'; k-- {
				if !first {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				first = false
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
			}
		}
	}
	z01.PrintRune('\n')
}

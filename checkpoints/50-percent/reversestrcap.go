//go:build ignore

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func reverseStrCap(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			isLast := false
			if i == len(runes)-1 {
				isLast = true
			} else if runes[i+1] == ' ' || runes[i+1] == '\t' || runes[i+1] == '\n' {
				isLast = true
			}

			if isLast {
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 'a' + 'A'
				}
			} else {
				if r >= 'A' && r <= 'Z' {
					runes[i] = r - 'A' + 'a'
				}
			}
		}
	}
	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, arg := range os.Args[1:] {
		res := reverseStrCap(arg)
		for _, r := range res {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

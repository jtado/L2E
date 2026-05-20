//go:build ignore

package main

import (
	"fmt"
	"os"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	firstVowelIdx := -1
	for i, r := range s {
		if isVowel(r) {
			firstVowelIdx = i
			break
		}
	}

	if firstVowelIdx == -1 {
		fmt.Println("No vowels")
		return
	}

	if firstVowelIdx == 0 {
		fmt.Println(s + "ay")
		return
	}

	fmt.Println(s[firstVowelIdx:] + s[:firstVowelIdx] + "ay")
}

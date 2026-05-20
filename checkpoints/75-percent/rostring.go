//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	s := os.Args[1]
	words := strings.Fields(s)
	if len(words) == 0 {
		fmt.Println()
		return
	}

	rotated := append(words[1:], words[0])
	fmt.Println(strings.Join(rotated, " "))
}

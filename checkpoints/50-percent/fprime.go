//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	div := 2
	first := true
	for n > 1 {
		if n%div == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(div)
			n /= div
			first = false
		} else {
			div++
		}
	}
	fmt.Println()
}

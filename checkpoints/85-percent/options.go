//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	for _, arg := range os.Args[1:] {
		if len(arg) >= 2 && arg[0] == '-' && arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}

	var optionsInt uint32

	for _, arg := range os.Args[1:] {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		for _, r := range arg[1:] {
			if r < 'a' || r > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			optionsInt |= (1 << uint(r-'a'))
		}
	}

	binStr := fmt.Sprintf("%032b", optionsInt)
	fmt.Printf("%s %s %s %s\n", binStr[0:8], binStr[8:16], binStr[16:24], binStr[24:32])
}

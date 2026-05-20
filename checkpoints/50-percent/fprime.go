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
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	var factors []int
	temp := num
	for d := 2; d*d <= temp; d++ {
		for temp%d == 0 {
			factors = append(factors, d)
			temp /= d
		}
	}
	if temp > 1 {
		factors = append(factors, temp)
	}

	if len(factors) == 0 {
		return
	}

	for i, f := range factors {
		fmt.Print(f)
		if i < len(factors)-1 {
			fmt.Print("*")
		}
	}
	fmt.Println()
}

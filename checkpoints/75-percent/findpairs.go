//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	inner := arrStr[1 : len(arrStr)-1]
	var arr []int
	if len(strings.TrimSpace(inner)) > 0 {
		elements := strings.Split(inner, ",")
		for _, elem := range elements {
			trimmed := strings.TrimSpace(elem)
			num, err := strconv.Atoi(trimmed)
			if err != nil {
				fmt.Printf("Invalid number: %s\n", trimmed)
				return
			}
			arr = append(arr, num)
		}
	}

	var pairs [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	var pairStrings []string
	for _, p := range pairs {
		pairStrings = append(pairStrings, fmt.Sprintf("[%d %d]", p[0], p[1]))
	}
	fmt.Printf("Pairs with sum %d: [%s]\n", target, strings.Join(pairStrings, " "))
}

//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	tokens := strings.Fields(os.Args[1])
	if len(tokens) == 0 {
		fmt.Println("Error")
		return
	}

	var stack []int

	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" || t == "%" {
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res int
			switch t {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a / b
			case "%":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a % b
			}
			stack = append(stack, res)
		} else {
			val, err := strconv.Atoi(t)
			if err != nil {
				fmt.Println("Error")
				return
			}
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}

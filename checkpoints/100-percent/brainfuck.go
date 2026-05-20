//go:build ignore

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	code := os.Args[1]

	// Precompute matching brackets
	matches := make(map[int]int)
	var stack []int
	for i, char := range code {
		if char == '[' {
			stack = append(stack, i)
		} else if char == ']' {
			if len(stack) == 0 {
				continue
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			matches[start] = i
			matches[i] = start
		}
	}

	var arr [2048]byte
	ptr := 0

	for pc := 0; pc < len(code); pc++ {
		char := code[pc]
		switch char {
		case '>':
			ptr++
			if ptr >= 2048 {
				ptr = 0 // Wrap around or handle safely
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = 2047 // Wrap around or handle safely
			}
		case '+':
			arr[ptr]++
		case '-':
			arr[ptr]--
		case '.':
			os.Stdout.Write([]byte{arr[ptr]})
		case '[':
			if arr[ptr] == 0 {
				if dest, ok := matches[pc]; ok {
					pc = dest
				}
			}
		case ']':
			if arr[ptr] != 0 {
				if dest, ok := matches[pc]; ok {
					pc = dest
				}
			}
		}
	}
}

//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

func isValidRegex(expr string) bool {
	if len(expr) < 2 || expr[0] != '(' || expr[len(expr)-1] != ')' {
		return false
	}
	inner := expr[1 : len(expr)-1]
	if len(inner) == 0 {
		return false
	}
	for _, r := range inner {
		if r == '(' || r == ')' {
			return false
		}
	}
	alternatives := strings.Split(inner, "|")
	for _, alt := range alternatives {
		if alt == "" {
			return false
		}
	}
	return true
}

func isWordChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '\'' || r == '’' || r == '-'
}

func cleanWord(w string) string {
	start := 0
	for start < len(w) {
		r := rune(w[start])
		if isWordChar(r) {
			break
		}
		start++
	}
	end := len(w)
	for end > start {
		r := rune(w[end-1])
		if isWordChar(r) {
			break
		}
		end--
	}
	return w[start:end]
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	expr := os.Args[1]
	text := os.Args[2]

	if text == "" {
		return
	}

	if !isValidRegex(expr) {
		return
	}

	inner := expr[1 : len(expr)-1]
	alternatives := strings.Split(inner, "|")

	words := strings.Fields(text)

	matchCount := 0
	for _, w := range words {
		cleaned := cleanWord(w)
		if cleaned == "" {
			continue
		}
		for _, alt := range alternatives {
			if strings.Contains(cleaned, alt) {
				matchCount++
				fmt.Printf("%d: %s\n", matchCount, cleaned)
			}
		}
	}
}

//go:build ignore

package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 2 {
		printString("0")
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		printString("0")
		return
	}
	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	printString(strconv.Itoa(sum))
}

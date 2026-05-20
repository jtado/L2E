//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Roman struct {
	val  int
	sym  string
	calc string
}

var romanMap = []Roman{
	{1000, "M", "M"},
	{900, "CM", "(M-C)"},
	{500, "D", "D"},
	{400, "CD", "(D-C)"},
	{100, "C", "C"},
	{90, "XC", "(C-X)"},
	{50, "L", "L"},
	{40, "XL", "(L-X)"},
	{10, "X", "X"},
	{9, "IX", "(X-I)"},
	{5, "V", "V"},
	{4, "IV", "(V-I)"},
	{1, "I", "I"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	temp := num
	var calcs []string
	var syms []string

	for _, r := range romanMap {
		for temp >= r.val {
			calcs = append(calcs, r.calc)
			syms = append(syms, r.sym)
			temp -= r.val
		}
	}

	fmt.Println(strings.Join(calcs, "+"))
	fmt.Println(strings.Join(syms, ""))
}

package main

import (
	"fmt"
	student "checkpoints/65-percent"
)

func main() {
	// FifthAndSkip
	fmt.Println("=== FifthAndSkip ===")
	fmt.Print(student.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(student.FifthAndSkip("This is a short sentence"))
	fmt.Print(student.FifthAndSkip("1234"))
	fmt.Print(student.FifthAndSkip(""))

	// NotDecimal
	fmt.Println("=== NotDecimal ===")
	fmt.Print(student.NotDecimal("0.1"))
	fmt.Print(student.NotDecimal("174.2"))
	fmt.Print(student.NotDecimal("0.1255"))
	fmt.Print(student.NotDecimal("1.20525856"))
	fmt.Print(student.NotDecimal("-0.0f00d00"))
	fmt.Print(student.NotDecimal(""))
	fmt.Print(student.NotDecimal("-19.525856"))
	fmt.Print(student.NotDecimal("1952"))

	// RevConcatAlternate
	fmt.Println("=== RevConcatAlternate ===")
	fmt.Println(student.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(student.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(student.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(student.RevConcatAlternate([]int{1, 2, 3}, []int{}))

	// Slice
	fmt.Println("=== Slice ===")
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", student.Slice(a, 1))
	fmt.Printf("%#v\n", student.Slice(a, 2, 4))
	fmt.Printf("%#v\n", student.Slice(a, -3))
	fmt.Printf("%#v\n", student.Slice(a, -2, -1))
	fmt.Printf("%#v\n", student.Slice(a, 2, 0))
}

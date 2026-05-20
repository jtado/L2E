package main

import (
	"fmt"
	student "checkpoints/20-percent"
)

func main() {
	// CamelToSnakeCase
	fmt.Println("=== CamelToSnakeCase ===")
	fmt.Println(student.CamelToSnakeCase("HelloWorld"))
	fmt.Println(student.CamelToSnakeCase("helloWorld"))
	fmt.Println(student.CamelToSnakeCase("camelCase"))
	fmt.Println(student.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(student.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(student.CamelToSnakeCase("hey2"))

	// DigitLen
	fmt.Println("=== DigitLen ===")
	fmt.Println(student.DigitLen(100, 10))
	fmt.Println(student.DigitLen(100, 2))
	fmt.Println(student.DigitLen(-100, 16))
	fmt.Println(student.DigitLen(100, -1))

	// FirstWord
	fmt.Println("=== FirstWord ===")
	fmt.Print(student.FirstWord("hello there"))
	fmt.Print(student.FirstWord(""))
	fmt.Print(student.FirstWord("hello   .........  bye"))

	// FishAndChips
	fmt.Println("=== FishAndChips ===")
	fmt.Println(student.FishAndChips(4))
	fmt.Println(student.FishAndChips(9))
	fmt.Println(student.FishAndChips(6))

	// Gcd
	fmt.Println("=== Gcd ===")
	fmt.Println(student.Gcd(42, 10))
	fmt.Println(student.Gcd(42, 12))
	fmt.Println(student.Gcd(14, 77))
	fmt.Println(student.Gcd(17, 3))

	// HashCode
	fmt.Println("=== HashCode ===")
	fmt.Println(student.HashCode("A"))
	fmt.Println(student.HashCode("AB"))
	fmt.Println(student.HashCode("BAC"))
	fmt.Println(student.HashCode("Hello World"))

	// LastWord
	fmt.Println("=== LastWord ===")
	fmt.Print(student.LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(student.LastWord(" lorem,ipsum "))
	fmt.Print(student.LastWord(" "))

	// RepeatAlpha
	fmt.Println("=== RepeatAlpha ===")
	fmt.Println(student.RepeatAlpha("abc"))
	fmt.Println(student.RepeatAlpha("Choumi."))
	fmt.Println(student.RepeatAlpha(""))
	fmt.Println(student.RepeatAlpha("abacadaba 01!"))
}

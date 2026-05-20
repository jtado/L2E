package main

import (
	student "checkpoints/10-percent"
	"fmt"
)

func main() {
	// CheckNumber
	fmt.Println("=== CheckNumber ===")
	fmt.Println(student.CheckNumber("Hello"))
	fmt.Println(student.CheckNumber("Hello1"))

	// CountAlpha
	fmt.Println("=== CountAlpha ===")
	fmt.Println(student.CountAlpha("Hello world"))
	fmt.Println(student.CountAlpha("H e l l o"))
	fmt.Println(student.CountAlpha("H1e2l3l4o"))

	// CountChar
	fmt.Println("=== CountChar ===")
	fmt.Println(student.CountChar("Hello World", 'l'))
	fmt.Println(student.CountChar("5  balloons", 5))
	fmt.Println(student.CountChar("   ", ' '))
	fmt.Println(student.CountChar("The 7 deadly sins", '7'))

	// PrintIf
	fmt.Println("=== PrintIf ===")
	fmt.Print(student.PrintIf("abcdefz"))
	fmt.Print(student.PrintIf("abc"))
	fmt.Print(student.PrintIf(""))
	fmt.Print(student.PrintIf("14"))

	// PrintIfNot
	fmt.Println("=== PrintIfNot ===")
	fmt.Print(student.PrintIfNot("abcdefz"))
	fmt.Print(student.PrintIfNot("abc"))
	fmt.Print(student.PrintIfNot(""))
	fmt.Print(student.PrintIfNot("14"))

	// RectPerimeter
	fmt.Println("=== RectPerimeter ===")
	fmt.Println(student.RectPerimeter(10, 2))
	fmt.Println(student.RectPerimeter(434343, 898989))
	fmt.Println(student.RectPerimeter(10, -2))

	// RetainFirstHalf
	fmt.Println("=== RetainFirstHalf ===")
	fmt.Println(student.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(student.RetainFirstHalf("A"))
	fmt.Println(student.RetainFirstHalf(""))
	fmt.Println(student.RetainFirstHalf("Hello World"))
}

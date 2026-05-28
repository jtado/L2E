package main

import (
	student "checkpoints/tests/35-percent"
	"fmt"
)

func main() {
	// FindPrevPrime
	fmt.Println("=== FindPrevPrime ===")
	fmt.Println(student.FindPrevPrime(5))
	fmt.Println(student.FindPrevPrime(4))
	fmt.Println(student.FindPrevPrime(0))

	// FromTo
	fmt.Println("=== FromTo ===")
	fmt.Print(student.FromTo(1, 10))
	fmt.Print(student.FromTo(10, 1))
	fmt.Print(student.FromTo(10, 10))
	fmt.Print(student.FromTo(100, 10))

	// IsCapitalized
	//fmt.Println("=== IsCapitalized ===")
	//fmt.Println(student.IsCapitalized("Hello! How are you?"))
	//fmt.Println(student.IsCapitalized("Hello How Are You"))
	//fmt.Println(student.IsCapitalized("Whats 4this 100K?"))
	//fmt.Println(student.IsCapitalized("Whatsthis4"))
	//fmt.Println(student.IsCapitalized("!!!!Whatsthis4"))
	//fmt.Println(student.IsCapitalized(""))

	// Itoa
	//fmt.Println("=== Itoa ===")
	//fmt.Println(student.Itoa(12345))
	//fmt.Println(student.Itoa(0))
	//fmt.Println(student.Itoa(-1234))
	//fmt.Println(student.Itoa(987654321))

	// PrintMemory
	//fmt.Println("=== PrintMemory ===")
	//student.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

	// PrintRevComb (Runnable main program, run via subprocess)
	//fmt.Println("=== PrintRevComb ===")
	//cmd := exec.Command("go", "run", "../printrevcomb.go")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//if err := cmd.Run(); err != nil {
	//	fmt.Printf("Error running printrevcomb: %v\n", err)

	// ThirdTimeIsACharm
	//fmt.Println("=== ThirdTimeIsACharm ===")
	//fmt.Print(student.ThirdTimeIsACharm("123456789"))
	//fmt.Print(student.ThirdTimeIsACharm(""))
	//fmt.Print(student.ThirdTimeIsACharm("a b c d e f"))
	//fmt.Print(student.ThirdTimeIsACharm("12"))

	// WeAreUnique
	//fmt.Println("=== WeAreUnique ===")
	//fmt.Println(student.WeAreUnique("foo", "boo"))
	//fmt.Println(student.WeAreUnique("", ""))
	//fmt.Println(student.WeAreUnique("abc", "def"))

	// ZipString
	//fmt.Println("=== ZipString ===")
	//fmt.Println(student.ZipString("YouuungFellllas"))
	//fmt.Println(student.ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	//fmt.Println(student.ZipString("Helloo Therre!"))
}

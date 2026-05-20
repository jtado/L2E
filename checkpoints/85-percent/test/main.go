package main

import (
	"fmt"
	"os"
	"os/exec"
	student "checkpoints/85-percent"
)

func runProgram(progPath string, args ...string) {
	fmt.Printf("go run %s %v -> ", progPath, args)
	cmd := exec.Command("go", append([]string{"run", progPath}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	// ItoaBase
	fmt.Println("=== ItoaBase ===")
	fmt.Println(student.ItoaBase(15, 16))
	fmt.Println(student.ItoaBase(-10, 2))
	fmt.Println(student.ItoaBase(0, 10))
	fmt.Println(student.ItoaBase(12345, 10))

	// Programs
	fmt.Println("=== Programs ===")
	// Options
	runProgram("../options.go")
	runProgram("../options.go", "-abc", "-ijk")
	runProgram("../options.go", "-z")
	runProgram("../options.go", "-abc", "-hijk")
	runProgram("../options.go", "-h")
	runProgram("../options.go", "-zh")
	runProgram("../options.go", "-z", "-h")
	runProgram("../options.go", "-hhhhhh")
	runProgram("../options.go", "-eeeeee")
	runProgram("../options.go", "-%")
	runProgram("../options.go", "-")

	// PigLatin
	runProgram("../piglatin.go")
	runProgram("../piglatin.go", "pig")
	runProgram("../piglatin.go", "Is")
	runProgram("../piglatin.go", "crunch")
	runProgram("../piglatin.go", "crnch")
	runProgram("../piglatin.go", "something", "else")

	// RomanNumbers
	runProgram("../romannumbers.go", "hello")
	runProgram("../romannumbers.go", "123")
	runProgram("../romannumbers.go", "999")
	runProgram("../romannumbers.go", "3999")
	runProgram("../romannumbers.go", "4000")
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	student "checkpoints/75-percent"
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
	// WordFlip
	fmt.Println("=== WordFlip ===")
	fmt.Print(student.WordFlip("First second last"))
	fmt.Print(student.WordFlip(""))
	fmt.Print(student.WordFlip("     "))
	fmt.Print(student.WordFlip(" hello  all  of  you! "))

	// Programs
	fmt.Println("=== Programs ===")
	// FindPairs
	runProgram("../findpairs.go", "[1, 2, 3, 4, 5]", "6")
	runProgram("../findpairs.go", "[-1, 2, -3, 4, -5]", "1")
	runProgram("../findpairs.go", "[1, 2, 3, 4, 5]", "10")
	runProgram("../findpairs.go", "[-1, -2, -3, -4, -5]", "-5")
	runProgram("../findpairs.go", "[1, 2, 3, 4, 20, -4, 5]", "2 5")
	runProgram("../findpairs.go", "[1, 2, 3, 4, 20, p, 5]", "5")
	runProgram("../findpairs.go", "[1, 2, 3, 4", "5")
	runProgram("../findpairs.go", "1, 2, 3, 4", "5")

	// RevWStr
	runProgram("../revwstr.go", "the time of contempt precedes that of indifference")
	runProgram("../revwstr.go", "abcdefghijklm")
	runProgram("../revwstr.go", "he stared at the mountain")
	runProgram("../revwstr.go", "")

	// RoString
	runProgram("../rostring.go", "abc   ")
	runProgram("../rostring.go", "Let there     be light")
	runProgram("../rostring.go", "     AkjhZ zLKIJz , 23y")
}

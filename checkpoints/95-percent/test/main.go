package main

import (
	"fmt"
	"os"
	"os/exec"
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
	// Programs
	fmt.Println("=== Programs ===")
	// Brackets
	runProgram("../brackets.go", "(johndoe)")
	runProgram("../brackets.go", "([)]")
	runProgram("../brackets.go", "", "{[(0 + 0)(1 + 1)](3*(-1)){()}}")
	runProgram("../brackets.go")

	// RpnCalc
	runProgram("../rpncalc.go", "1 2 * 3 * 4 +")
	runProgram("../rpncalc.go", "1 2 3 4 +")
	runProgram("../rpncalc.go")
	runProgram("../rpncalc.go", "     1      3 * 2 -")
	runProgram("../rpncalc.go", "     1      3 * ksd 2 -")
}

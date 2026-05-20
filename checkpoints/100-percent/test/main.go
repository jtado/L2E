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
	// Brainfuck
	runProgram("../brainfuck.go", "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
	fmt.Println()
	runProgram("../brainfuck.go", "+++++[>++++[>++++H>+++++i<<-]>>>++\\n<<<<-]>>--------.>+++++.>.")
	fmt.Println()
	runProgram("../brainfuck.go", "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++.")
	fmt.Println()
	runProgram("../brainfuck.go")

	// Grouping
	runProgram("../grouping.go", "(a)", "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy")
	runProgram("../grouping.go", "(e|n)", "I currently have 4 windows opened up… and I don’t know why.")
	runProgram("../grouping.go", "(hi)", "He swore he just saw his sushi move.")
	runProgram("../grouping.go", "(s)", "")
	runProgram("../grouping.go", "i", "Something in the air")
}

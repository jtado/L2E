package main

import (
	"fmt"
	"os"
	"os/exec"
	student "checkpoints/50-percent"
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
	// CanJump
	fmt.Println("=== CanJump ===")
	fmt.Println(student.CanJump([]uint{2, 3, 1, 1, 4}))
	fmt.Println(student.CanJump([]uint{3, 2, 1, 0, 4}))
	fmt.Println(student.CanJump([]uint{0}))

	// Chunk
	fmt.Println("=== Chunk ===")
	student.Chunk([]int{}, 10)
	student.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	student.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	student.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	student.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)

	// ConcatAlternate
	fmt.Println("=== ConcatAlternate ===")
	fmt.Println(student.ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(student.ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(student.ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(student.ConcatAlternate([]int{1, 2, 3}, []int{}))

	// ConcatSlice
	fmt.Println("=== ConcatSlice ===")
	fmt.Println(student.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(student.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(student.ConcatSlice([]int{1, 2, 3}, []int{}))

	// SaveAndMiss
	fmt.Println("=== SaveAndMiss ===")
	fmt.Println(student.SaveAndMiss("123456789", 3))
	fmt.Println(student.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(student.SaveAndMiss("", 3))
	fmt.Println(student.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(student.SaveAndMiss("what is your name?", 0))
	fmt.Println(student.SaveAndMiss("go Exercise Save and Miss", -5))

	// Programs
	fmt.Println("=== Programs ===")
	// AddPrimeSum
	runProgram("../addprimesum.go", "5")
	runProgram("../addprimesum.go", "7")
	runProgram("../addprimesum.go", "-2")
	runProgram("../addprimesum.go", "0")
	runProgram("../addprimesum.go")
	runProgram("../addprimesum.go", "5", "7")

	// FPrime
	runProgram("../fprime.go", "225225")
	runProgram("../fprime.go", "8333325")
	runProgram("../fprime.go", "9539")
	runProgram("../fprime.go", "804577")
	runProgram("../fprime.go", "42")
	runProgram("../fprime.go", "a")
	runProgram("../fprime.go", "0")
	runProgram("../fprime.go", "1")

	// HiddenP
	runProgram("../hiddenp.go", "fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6")
	runProgram("../hiddenp.go", "abc", "2altrb53c.sse")
	runProgram("../hiddenp.go", "abc", "btarc")
	runProgram("../hiddenp.go", "DD", "DABC")

	// Inter
	runProgram("../inter.go", "padinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
	runProgram("../inter.go", "ddf6vewg64f", "twthgdwthdwfteewhrtag6h4ffdhsd")

	// ReverseStrCap
	runProgram("../reversestrcap.go", "First SMALL TesT")
	runProgram("../reversestrcap.go", "SEconD Test IS a LItTLE EasIEr", "bEwaRe IT'S NoT HARd WhEN ", " Go a dernier 0123456789 for the road e")

	// Union
	runProgram("../union.go", "zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
	runProgram("../union.go", "ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd")
	runProgram("../union.go", "rien", "cette phrase ne cache rien")
	runProgram("../union.go")
	runProgram("../union.go", "rien")

	// WdMatch
	runProgram("../wdmatch.go", "123", "123")
	runProgram("../wdmatch.go", "faya", "fgvvfdxcacpolhyghbreda")
	runProgram("../wdmatch.go", "faya", "fgvvfdxcacpolhyghbred")
	runProgram("../wdmatch.go", "error", "rrerrrfiiljdfxjyuifrrvcoojh")
	runProgram("../wdmatch.go", "quarante deux", "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq ")
}

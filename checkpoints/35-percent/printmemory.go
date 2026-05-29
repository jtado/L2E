package student

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i := 0; i < 10; i++ {
		h1 := arr[i] / 16
		h2 := arr[i] % 16

		if h1 < 10 {
			z01.PrintRune(rune(h1 + '0'))
		} else {
			z01.PrintRune(rune(h1 - 10 + 'a'))
		}

		if h2 < 10 {
			z01.PrintRune(rune(h2 + '0'))
		} else {
			z01.PrintRune(rune(h2 - 10 + 'a'))
		}

		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

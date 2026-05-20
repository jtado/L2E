package student

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexChars := "0123456789abcdef"
	for i := 0; i < len(arr); i++ {
		b := arr[i]
		z01.PrintRune(rune(hexChars[b>>4]))
		z01.PrintRune(rune(hexChars[b&0x0f]))

		switch i {
		case 3, 7, 9:
			z01.PrintRune('\n')
		default:
			z01.PrintRune(' ')
		}
	}

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

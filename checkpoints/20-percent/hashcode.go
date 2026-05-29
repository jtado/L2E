package student

func HashCode(dec string) string {
	l := 0
	for range dec {
		l++
	}

	res := ""
	for _, r := range dec {
		val := (int(r) + l) % 127
		if val < 32 || val > 126 {
			val += 33
		}
		res += string(rune(val))
	}
	return res
}

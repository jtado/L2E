package student

func HashCode(dec string) string {
	var result []rune
	l := len(dec)
	for _, r := range dec {
		val := (int(r) + l) % 127
		if val < 32 || val > 126 {
			val += 33
		}
		result = append(result, rune(val))
	}
	return string(result)
}

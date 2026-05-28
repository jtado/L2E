package student

func HashCode(dec string) string {
	size := 0
	for range dec {
		size++
	}

	var res []rune
	for _, char := range dec {
		val := (int(char) + size) % 127
		if val < 32 {
			val += 33
		}
		res = append(res, rune(val))
	}
	return string(res)
}

package student

func RetainFirstHalf(str string) string {
	runes := []rune(str)
	if len(runes) == 1 {
		return str
	}
	return string(runes[:len(runes)/2])
}

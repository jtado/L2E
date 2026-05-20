package student

func ThirdTimeIsACharm(str string) string {
	runes := []rune(str)
	if len(runes) < 3 {
		return "\n"
	}
	var result []rune
	for i := 2; i < len(runes); i += 3 {
		result = append(result, runes[i])
	}
	return string(result) + "\n"
}

package student

func LastWord(s string) string {
	res := ""
	end := -1
	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] != ' ' {
			end = i
			break
		}
	}

	if end == -1 {
		return "\n"
	}

	start := 0
	for i := end; i >= 0; i-- {
		if runes[i] == ' ' {
			start = i + 1
			break
		}
	}

	for i := start; i <= end; i++ {
		res += string(runes[i])
	}

	return res + "\n"
}

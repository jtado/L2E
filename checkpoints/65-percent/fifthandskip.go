package student

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	var filtered []rune
	for _, r := range str {
		if r != ' ' {
			filtered = append(filtered, r)
		}
	}
	if len(filtered) < 5 {
		return "Invalid Input\n"
	}

	var chunks []string
	for i := 0; i < len(filtered); {
		end := i + 5
		if end > len(filtered) {
			end = len(filtered)
		}
		chunks = append(chunks, string(filtered[i:end]))
		i = end + 1
	}

	var res string
	for idx, chunk := range chunks {
		res += chunk
		if idx < len(chunks)-1 {
			res += " "
		}
	}
	return res + "\n"
}

package student

func LastWord(s string) string {
	runes := []rune(s)
	end := len(runes) - 1
	for end >= 0 && (runes[end] == ' ' || runes[end] == '\t' || runes[end] == '\n') {
		end--
	}
	if end < 0 {
		return "\n"
	}
	start := end
	for start >= 0 && runes[start] != ' ' && runes[start] != '\t' && runes[start] != '\n' {
		start--
	}
	return string(runes[start+1:end+1]) + "\n"
}

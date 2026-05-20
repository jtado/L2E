package student

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	runes := []rune(s)
	hasWord := false
	for i := 0; i < len(runes); i++ {
		isWordStart := false
		if i == 0 {
			if runes[i] != ' ' && runes[i] != '\t' && runes[i] != '\n' {
				isWordStart = true
			}
		} else {
			if (runes[i-1] == ' ' || runes[i-1] == '\t' || runes[i-1] == '\n') &&
				(runes[i] != ' ' && runes[i] != '\t' && runes[i] != '\n') {
				isWordStart = true
			}
		}
		if isWordStart {
			hasWord = true
			firstChar := runes[i]
			if firstChar >= 'a' && firstChar <= 'z' {
				return false
			}
		}
	}
	return hasWord
}

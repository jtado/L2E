package student

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	isStartOfWord := true
	for _, char := range s {
		if isStartOfWord {
			if char >= 'a' && char <= 'z' {
				return false
			}
			isStartOfWord = false
		}
		if char == ' ' {
			isStartOfWord = true
		}
	}
	return true
}

package student

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	unique := ""

	// Check characters in str1 not in str2
	for _, char := range str1 {
		if !contains(str2, char) && !contains(unique, char) {
			unique += string(char)
		}
	}

	// Check characters in str2 not in str1
	for _, char := range str2 {
		if !contains(str1, char) && !contains(unique, char) {
			unique += string(char)
		}
	}

	count := 0
	for range unique {
		count++
	}
	return count
}

func contains(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

package student

func isCamelCase(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	last := s[len(s)-1]
	if last >= 'A' && last <= 'Z' {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return false
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if !isCamelCase(s) {
		return s
	}
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(result)
}

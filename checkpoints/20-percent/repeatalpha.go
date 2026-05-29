package student

func RepeatAlpha(s string) string {
	res := ""
	for _, char := range s {
		count := 1
		if char >= 'a' && char <= 'z' {
			count = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			count = int(char - 'A' + 1)
		}
		for i := 0; i < count; i++ {
			res += string(char)
		}
	}
	return res
}

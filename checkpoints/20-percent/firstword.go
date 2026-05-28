package student

func FirstWord(s string) string {
	res := ""
	started := false
	for _, char := range s {
		if char != ' ' {
			res += string(char)
			started = true
		} else if started {
			break
		}
	}
	return res + "\n"
}

package student

func FirstWord(s string) string {
	var word []rune
	started := false
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if started {
				break
			}
		} else {
			started = true
			word = append(word, r)
		}
	}
	return string(word) + "\n"
}

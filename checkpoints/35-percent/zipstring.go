package student

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	res := ""
	count := 0
	var lastRune rune

	for i, r := range s {
		if i == 0 {
			lastRune = r
			count = 1
			continue
		}
		if r == lastRune {
			count++
		} else {
			res += itoa(count) + string(lastRune)
			lastRune = r
			count = 1
		}
	}
	res += itoa(count) + string(lastRune)
	return res
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune(n%10+'0')) + s
		n /= 10
	}
	return s
}

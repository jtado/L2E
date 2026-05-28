package student

func FromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	res := ""
	i := from
	for {
		res += string(rune(i/10+'0')) + string(rune(i%10+'0'))
		if i == to {
			break
		}
		res += ", "
		if from < to {
			i++
		} else {
			i--
		}
	}
	return res + "\n"
}

package student

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	res := ""
	if from <= to {
		for i := from; i <= to; i++ {
			res += string(byte(i/10+'0')) + string(byte(i%10+'0'))
			if i != to {
				res += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			res += string(byte(i/10+'0')) + string(byte(i%10+'0'))
			if i != to {
				res += ", "
			}
		}
	}
	return res + "\n"
}

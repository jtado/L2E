package student

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	var res string
	if from <= to {
		for i := from; i <= to; i++ {
			res += fmt.Sprintf("%02d", i)
			if i < to {
				res += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			res += fmt.Sprintf("%02d", i)
			if i > to {
				res += ", "
			}
		}
	}
	res += "\n"
	return res
}

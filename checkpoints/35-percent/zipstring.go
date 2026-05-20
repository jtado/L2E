package student

import "strconv"

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	var res string
	current := runes[0]
	count := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == current {
			count++
		} else {
			res += strconv.Itoa(count) + string(current)
			current = runes[i]
			count = 1
		}
	}
	res += strconv.Itoa(count) + string(current)
	return res
}

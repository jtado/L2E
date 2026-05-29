package student

func ThirdTimeIsACharm(str string) string {
	res := ""
	count := 0
	for _, char := range str {
		count++
		if count%3 == 0 {
			res += string(char)
		}
	}
	return res + "\n"
}

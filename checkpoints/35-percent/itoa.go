package student

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	isNeg := n < 0
	for n != 0 {
		digit := n % 10
		if digit < 0 {
			digit = -digit
		}
		res = string(rune(digit+'0')) + res
		n /= 10
	}
	if isNeg {
		res = "-" + res
	}
	return res
}

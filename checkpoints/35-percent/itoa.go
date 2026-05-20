package student

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	var val uint64
	if n < 0 {
		sign = "-"
		val = uint64(^n) + 1
	} else {
		val = uint64(n)
	}
	var res []byte
	for val > 0 {
		digit := val % 10
		res = append([]byte{byte('0' + digit)}, res...)
		val /= 10
	}
	return sign + string(res)
}

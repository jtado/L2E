package student

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	val := int64(n)
	if val < 0 {
		val = -val
	}
	count := 0
	for val > 0 {
		val = val / int64(base)
		count++
	}
	return count
}

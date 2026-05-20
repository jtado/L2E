package student

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}
	if value == 0 {
		return "0"
	}

	neg := value < 0
	var uval uint64
	if neg {
		uval = uint64(-value)
	} else {
		uval = uint64(value)
	}

	baseChars := "0123456789ABCDEF"
	var res []byte
	for uval > 0 {
		rem := uval % uint64(base)
		res = append(res, baseChars[rem])
		uval /= uint64(base)
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	if neg {
		return "-" + string(res)
	}
	return string(res)
}

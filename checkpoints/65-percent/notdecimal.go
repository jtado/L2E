package student

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	isAllDigits := func(s string) bool {
		if s == "" {
			return false
		}
		for _, r := range s {
			if r < '0' || r > '9' {
				return false
			}
		}
		return true
	}

	isAllZeros := func(s string) bool {
		if s == "" {
			return false
		}
		for _, r := range s {
			if r != '0' {
				return false
			}
		}
		return true
	}

	dotCount := 0
	dotIdx := -1
	for i, r := range dec {
		if r == '.' {
			dotCount++
			dotIdx = i
		}
	}

	if dotCount > 1 {
		return dec + "\n"
	}

	if dotCount == 0 {
		hasSign := dec[0] == '+' || dec[0] == '-'
		var body string
		if hasSign {
			body = dec[1:]
		} else {
			body = dec
		}
		if isAllDigits(body) {
			return dec + "\n"
		}
		return dec + "\n"
	}

	part1 := dec[:dotIdx]
	part2 := dec[dotIdx+1:]

	if part1 == "" {
		return dec + "\n"
	}
	hasSign := part1[0] == '+' || part1[0] == '-'
	var part1Body string
	if hasSign {
		part1Body = part1[1:]
	} else {
		part1Body = part1
	}

	if !isAllDigits(part1Body) || !isAllDigits(part2) {
		return dec + "\n"
	}

	if isAllZeros(part2) {
		return dec + "\n"
	}

	combined := part1Body + part2
	for len(combined) > 1 && combined[0] == '0' {
		combined = combined[1:]
	}

	sign := ""
	if part1[0] == '-' && combined != "0" {
		sign = "-"
	}

	return sign + combined + "\n"
}

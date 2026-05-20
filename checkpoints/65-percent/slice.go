package student

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	l := len(a)
	start := nbrs[0]
	if start < 0 {
		start = l + start
		if start < 0 {
			start = 0
		}
	}
	if start > l {
		start = l
	}

	end := l
	if len(nbrs) >= 2 {
		end = nbrs[1]
		if end < 0 {
			end = l + end
			if end < 0 {
				end = 0
			}
		}
		if end > l {
			end = l
		}
	}

	if start >= end {
		return nil
	}

	return a[start:end]
}

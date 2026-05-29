package student

func ConcatAlternate(slice1, slice2 []int) []int {
	first := slice1
	second := slice2

	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	res := make([]int, 0, len(slice1)+len(slice2))
	for i := 0; i < len(first); i++ {
		res = append(res, first[i])
		if i < len(second) {
			res = append(res, second[i])
		}
	}
	return res
}

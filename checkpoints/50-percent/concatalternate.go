package student

func ConcatAlternate(slice1, slice2 []int) []int {
	var first, second []int
	if len(slice1) > len(slice2) {
		first, second = slice1, slice2
	} else if len(slice2) > len(slice1) {
		first, second = slice2, slice1
	} else {
		first, second = slice1, slice2
	}

	var res []int
	i := 0
	for i < len(first) && i < len(second) {
		res = append(res, first[i])
		res = append(res, second[i])
		i++
	}
	for i < len(first) {
		res = append(res, first[i])
		i++
	}
	return res
}

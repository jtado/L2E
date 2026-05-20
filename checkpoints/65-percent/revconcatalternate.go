package student

func reverse(s []int) []int {
	res := make([]int, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	r1 := reverse(slice1)
	r2 := reverse(slice2)

	var res []int

	if len(r1) > len(r2) {
		diff := len(r1) - len(r2)
		res = append(res, r1[:diff]...)
		for i := 0; i < len(r2); i++ {
			res = append(res, r1[diff+i], r2[i])
		}
	} else if len(r2) > len(r1) {
		diff := len(r2) - len(r1)
		res = append(res, r2[:diff]...)
		for i := 0; i < len(r1); i++ {
			res = append(res, r1[i], r2[diff+i])
		}
	} else {
		for i := 0; i < len(r1); i++ {
			res = append(res, r1[i], r2[i])
		}
	}
	return res
}

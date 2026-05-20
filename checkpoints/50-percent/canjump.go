package student

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}
	i := 0
	for i < len(steps)-1 {
		if steps[i] == 0 {
			return false
		}
		i += int(steps[i])
	}
	return i == len(steps)-1
}

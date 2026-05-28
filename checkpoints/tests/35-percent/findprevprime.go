package student

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	for i := nb; i >= 2; i-- {
		isPrime := true
		if i <= 1 {
			isPrime = false
		} else {
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
		}
		if isPrime {
			return i
		}
	}
	return 0
}

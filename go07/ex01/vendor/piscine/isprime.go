package piscine

func IsPrime(nb int) bool {

	if nb <= 1 {
		return false
	}

	seed := []int{2, 3, 5}

	for _, p := range seed {
		if nb == p {
			return true
		}
		if nb%p == 0 {
			return false
		}
	}

	seedPro := 1
	for _, p := range seed {
		seedPro *= p
	}

	primeLine := []int{}
	for i := 0; i < seedPro; i++ {
		var check bool = true
		for _, p := range seed {
			if i%p == 0 {
				check = false
				break
			}
		}
		if check {
			primeLine = append(primeLine, i)
		}
	}

	for i := 0;; i+= seedPro {
		if i*i < nb && i*i >= 0 && i*i < (i+seedPro)*(i+seedPro) {
			for _, p := range primeLine {
				if (i+p) == 1 {
					continue
				}
				if nb%(i+p) == 0 && nb != i+p {
					return false
				}
			}
		} else {
			break
		}
	}
	return true
}
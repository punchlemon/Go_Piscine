package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	prev := 0
	new := 0
	for i, v := range a {
		prev = new
		new = v
		if i == 0 {
			continue
		}
		if f(prev, new) == 1 {
			return false
		}
	}
	return true
}

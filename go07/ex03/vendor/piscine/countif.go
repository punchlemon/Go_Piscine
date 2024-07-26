package piscine

func CountIf(f func(string) bool, arr []string) int {
	count := 0
	for _, v := range arr {
		if f(v) {
			count++
		}
	}
	return count
}

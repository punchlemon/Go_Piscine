package piscine

func StrLen(s string) int {
	len := 0
	if s != "" {
		for range s {
			len++
		}
	}
	return len
}

func Index(s string, toFind string) int {
	s_len := StrLen(s)
	f_len := StrLen(toFind)
	if s_len == 0 || s_len < f_len {
		return -1
	}
	for i := 0; i < s_len; i++ {
		valid := true
		for j := 0; j < f_len; j++ {
			if s[i + j] != toFind[j] {
				valid = false
				break
			}
			if valid {
				return i
			}
		}
	}
	return -1
}
package number

func Contains(str string, predicate func(rune) bool) bool {
	for _, r := range str {
		if predicate(r) {
			return true
		}
	}
	return false
}

func ContainsDigit(str string) bool {
	return Contains(str, IsDigit)
}

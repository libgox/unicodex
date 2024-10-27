package letter

func Contains(str string, predicate func(rune) bool) bool {
	for _, r := range str {
		if predicate(r) {
			return true
		}
	}
	return false
}

func ContainsEnglishUpper(str string) bool {
	return Contains(str, IsEnglishUpper)
}

func ContainsEnglishLower(str string) bool {
	return Contains(str, IsEnglishLower)
}

func ContainsEnglish(str string) bool {
	return Contains(str, IsEnglish)
}

func ContainsGreekUpper(str string) bool {
	return Contains(str, IsGreekUpper)
}

func ContainsGreekLower(str string) bool {
	return Contains(str, IsGreekLower)
}

func ContainsGreek(str string) bool {
	return Contains(str, IsGreek)
}

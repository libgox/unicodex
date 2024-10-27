package letter

func IsEnglishUpper(r rune) bool {
	return r >= EnglishUpperMin && r <= EnglishUpperMax
}

func IsEnglishLower(r rune) bool {
	return r >= EnglishLowerMin && r <= EnglishLowerMax
}

func IsEnglish(r rune) bool {
	return IsEnglishUpper(r) || IsEnglishLower(r)
}

func IsGreekUpper(r rune) bool {
	return r >= GreekUpperMin && r <= GreekUpperMax
}

func IsGreekLower(r rune) bool {
	return r >= GreekLowerMin && r <= GreekLowerMax
}

func IsGreek(r rune) bool {
	return IsGreekUpper(r) || IsGreekLower(r)
}

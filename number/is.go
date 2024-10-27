package number

func IsDigit(r rune) bool {
	return r >= DigitMin && r <= DigitMax
}
